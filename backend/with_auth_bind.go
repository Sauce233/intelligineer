package main

import (
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"

	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Config struct {
	DB  *gorm.DB
	Env *EnvConfig
	RDB *redis.Client
	OSS *oss.Client
}

type Context struct {
	User *UserClaim
	*Config
	*gin.Context
}

type UserClaim struct {
	ID    uint     `json:"id"`
	Roles []string `json:"roles"`
	jwt.RegisteredClaims
}

type Resp struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func Res(message string, data any) *Resp {
	return &Resp{
		Message: message,
		Data:    data,
	}
}

type WithAuthBindFunc[URI, Req, Resp any] func(
	c *Context, i *URI, r *Req,
) (message string, resp Resp)

type Option struct {
	NotLogin bool
}

func WithNotLogin(o *Option) {
	o.NotLogin = true
}

func WithAuthBind[URI, Req, Resp any](
	cfg *Config,
	ahf WithAuthBindFunc[URI, Req, Resp],
) gin.HandlerFunc {

	return func(c *gin.Context) {

		var claim *UserClaim
		/*

			if !opt.NotLogin {

				var err error
				claim, err = GetUserClaimFromAccessToken(c, cfg.Env.JWTKey)
				if err != nil {
					claim, err = GetUserClaimFromRefreshToken(c, cfg)
					if err != nil {
						c.JSON(401, Res("你的会话已过期，请重新登录", nil))
						return
					}
				}

				rbacKey := "rbac:" + c.Request.Method + c.FullPath()
				roles, err := cfg.RDB.SMembers(c.Request.Context(), rbacKey).Result()
				if len(roles) != 0 {
					var hasPerm bool
					for _, role := range roles {
						if slices.Contains(claim.Roles, role) {
							hasPerm = true
							break
						}
					}
					if !hasPerm {
						c.JSON(403, Res("你没有权限执行", nil))
						return
					}
				} else if err != nil && !errors.Is(err, redis.Nil) {
					c.Error(fmt.Errorf("failed to load rbac roles %s: %w", rbacKey, err))
					c.JSON(500, Res("无法获取你的权限信息", nil))
					return
				}
			}
		*/

		var uri URI
		if err := c.ShouldBindUri(&uri); err != nil {
			c.JSON(400, Res("路径参数解析失败", nil))
			return
		}

		var req Req

		if c.Request.Method == "GET" {
			if err := c.ShouldBindQuery(&req); err != nil {
				c.JSON(400, Res("查询字符串参数解析失败", nil))
				return
			}
		} else if c.Request.Method == "POST" || c.Request.Method == "PATCH" {
			if err := c.ShouldBindBodyWithJSON(&req); err != nil && !errors.Is(err, io.EOF) {
				c.JSON(400, Res("请求体解析失败", nil))
				return
			}
		}

		message, data := ahf(&Context{
			Config:  cfg,
			User:    claim,
			Context: c,
		}, &uri, &req)

		c.JSON(200, Res(message, data))
	}
}

func GetUserClaimFromAccessToken(c *gin.Context, jwtKey string) (*UserClaim, error) {

	token := strings.TrimPrefix(c.GetHeader("Authorization"), "Bearer ")
	if token == "" {
		return nil, errors.New("no access key in request header")
	}

	rawClaim, err := jwt.ParseWithClaims(token, new(UserClaim), func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected method: %v", token.Header["alg"])
		}
		return []byte(jwtKey), nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to parse access token: %w", err)
	}

	claim, ok := rawClaim.Claims.(*UserClaim)
	if !ok {
		return nil, errors.New("wrong type for access key")
	}

	return claim, nil
}

func GetUserClaimFromRefreshToken(c *gin.Context, cfg *Config) (*UserClaim, error) {

	refreshToken, err := c.Cookie("token")
	if err != nil {
		return nil, fmt.Errorf("failed to get refresh token from cookie: %w", err)
	}

	strUserId, err := cfg.RDB.Get(c.Request.Context(), "token:"+refreshToken).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to get user id by refresh token from redis: %w", err)
	}

	userId, err := strconv.Atoi(strUserId)
	if err != nil {
		return nil, fmt.Errorf("failed to convert string user id to int: %w", err)
	}

	var roles []string
	if err := cfg.DB.Model(new(UserRole)).Where("user_id = ?", userId).Pluck("role_id", &roles).Error; err != nil {
		return nil, fmt.Errorf("failed to get user roles from MySQL: %w", err)
	}
	if err := cfg.RDB.Del(c.Request.Context(), "token:"+refreshToken).Err(); err != nil {
		return nil, fmt.Errorf("failed to delete old refresh token in redis: %w", err)
	}

	if err := GetRefreshToken(uint(userId), c, cfg); err != nil {
		return nil, fmt.Errorf("failed to get new refresh token: %w", err)
	}

	userClaim := UserClaim{
		ID:    uint(userId),
		Roles: roles,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(cfg.Env.Duration.AccessToken)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    cfg.Env.JWTIssuer,
		},
	}

	newAccessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, userClaim).SignedString([]byte(cfg.Env.JWTKey))
	if err != nil {
		return nil, fmt.Errorf("failed to generate access token: %w", err)
	}
	c.Header("X-Access-Token", newAccessToken)

	return &userClaim, nil
}
