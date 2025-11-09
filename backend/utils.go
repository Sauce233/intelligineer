package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/goccy/go-yaml"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func PreloadAll(depth uint) func(db *gorm.Statement) {
	return func(db *gorm.Statement) {
		db.Preload(clause.Associations, PreloadAssociations(depth-1))
	}
}

type PaginateConfig struct {
	KeyPrefix       string
	PageKey         string
	PageSizeKey     string
	DefaultPage     int
	DefaultPageSize int
}

func Paginate(page, size, def int) func(db *gorm.Statement) {
	if page == 0 {
		page = 1
	}
	if size == 0 {
		size = def
	}
	return func(db *gorm.Statement) {
		db.Offset((page - 1) * size).Limit(size)
	}
}

func PreloadAssociations(level uint) func(db *gorm.DB) *gorm.DB {
	if level == 0 {
		return func(db *gorm.DB) *gorm.DB {
			return db
		}
	}
	return func(db *gorm.DB) *gorm.DB {
		return db.Preload(clause.Associations, PreloadAssociations(level-1))
	}
}

func GetRefreshToken(userId uint, c *gin.Context, cfg *Config) error {

	tokenUUID, err := uuid.NewRandom()
	if err != nil {
		return fmt.Errorf("failed to generate uuid: %w", err)
	}
	token := tokenUUID.String()

	if err := cfg.RDB.Set(c, "token:"+token, userId, cfg.Env.Duration.RefreshToken).Err(); err != nil {
		return fmt.Errorf("failed to store uuid in redis: %w", err)
	}
	c.SetCookie(
		"token",
		token,
		int(cfg.Env.Duration.RefreshToken.Seconds()),
		cfg.Env.Cookie.Path,
		cfg.Env.Cookie.Domain,
		cfg.Env.Cookie.Secure,
		cfg.Env.Cookie.HTTPOnly,
	)

	return nil
}

func LoadRBAC(path string, rdb *redis.Client) error {

	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("failed to open rbac file: %w", err)
	}

	var perms map[string][]any
	if err := yaml.Unmarshal(data, &perms); err != nil {
		return fmt.Errorf("failed to unmarshal rbac file: %w", err)
	}

	for k, v := range perms {
		if err := rdb.Del(context.TODO(), k).Err(); err != nil {
			return fmt.Errorf("failed to delete redis key: %w", err)
		}
		if err := rdb.SAdd(context.TODO(), "rbac:"+k, v...).Err(); err != nil {
			return fmt.Errorf("failed to set redis rbac key: %w", err)
		}
	}

	return nil
}

/*
func PreloadAll(level uint) func(s *gorm.Statement) {
	if level == 0 {
		return func(s *gorm.Statement) {}
	}
	return func(s *gorm.Statement) {
		s.Preload(clause.Associations, PreloadAssociations(level-1))
	}
}
*/

type List struct {
	Data  any   `json:"data"`
	Total int64 `json:"total"`
}

func NewList(data any, total int64) *List {
	return &List{Data: data, Total: total}
}

func GetSQLUintList(n int) string {
	var parts []string
	for i := 0; i < n; i++ {
		parts = append(parts, fmt.Sprintf("SELECT %d AS n", i))
	}
	return "(" + strings.Join(parts, " UNION ALL ") + ")"
}
