package main

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"html/template"
	"math/rand"
	"net/smtp"
	"time"

	"github.com/redis/go-redis/v9"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func GetMyInfo(c *Context, i *struct{}, q *struct{}, j *struct{}) (message string, resp any) {

	type UserRole struct {
	}

	type Record struct {
	}

	type Profile struct {
	}

	type UserAttendance struct {
	}

	type User struct {
		ID              uint             `json:"id"`
		CreatedAt       time.Time        `json:"createdAt"`
		UpdatedAt       time.Time        `json:"updatedAt"`
		Avatar          string           `json:"avatar"`
		Password        string           `json:"-"`
		Name            string           `json:"name"`
		Email           string           `json:"email"`
		Roles           []Role           `json:"roles"`
		Records         []Record         `json:"records"`
		Profiles        []Profile        `json:"profiles"`
		UserAttendances []UserAttendance `json:"userAttendances"`
	}

	var user User
	if err := c.DB.Scopes(PreloadAssociations(2)).Take(&user, c.User.ID).Error; err != nil {
		c.Error(err)
		return "获取个人信息失败", nil
	}

	return "", &user
}

func Login(c *Context, i *struct{}, q *struct{}, j *struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}) (string, bool) {

	user, err := gorm.G[User](c.DB).Where("email = ?", j.Email).Select("id", "password").Take(context.TODO())
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return "找不到这个用户", false
	} else if err != nil {
		c.Error(err)
		return "查找用户失败", false
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(j.Password)); errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
		return "密码不正确", false
	} else if err != nil {
		c.Error(fmt.Errorf("failed to check password: %w", err))
		return "密码校验失败", false
	}

	if err := GetRefreshToken(user.ID, c.Context, c.Config); err != nil {
		c.Error(fmt.Errorf("failed to generate token: %w", err))
		return "令牌生成失败", false
	}

	return "登录成功", true
}

/*
func ListRole(c *Context, u *struct{}, r *struct{}) (message string, resp any) {

	var roles []Role
	if err := c.DB.Find(&roles).Error; err != nil {
		c.Error(err)
		return "获取角色列表失败", nil
	}

	return "", roles
}
*/

func Retrieve(c *Context, i *struct{}, q *struct{}, j *struct {
	Email      string `json:"email" binding:"required,email"`
	VerifyCode string `json:"verifyCode" binding:"required"`
	Password   string `json:"password"`
}) (message string, resp bool) {

	verifyCode, err := c.RDB.Get(context.TODO(), "verify:email:"+j.Email).Result()
	if errors.Is(err, redis.Nil) {
		return "你还没有申请验证码", false
	} else if err != nil {
		c.Error(fmt.Errorf("failed to get email code: %w", err))
		return "获取验证码失败", false
	}

	if verifyCode != j.VerifyCode {
		return "验证码不正确", false
	}

	password, err := bcrypt.GenerateFromPassword([]byte(j.Password), bcrypt.DefaultCost)
	if err != nil {
		c.Error(fmt.Errorf("failed to generate password: %w", err))
		return "密码生成失败", false
	}

	row, err := gorm.G[User](c.DB).Where("email = ?", j.Email).Update(context.TODO(), "password", string(password))
	if row == 0 {
		return "你尚未注册", false
	} else if err != nil {
		c.Error(fmt.Errorf("failed to update user password: %w", err))
		return "密码更新失败", false
	}

	return "密码重置成功", true
}

func GetEmailCode(c *Context, i *struct {
	Email string `uri:"email" binding:"required,email"`
}, q *struct{}, j *struct{}) (string, bool) {

	code := fmt.Sprintf("%06d", rand.Intn(1000000))

	if err := c.RDB.Set(context.TODO(), "verify:email:"+i.Email, code, c.Env.Duration.VerifyCode).Err(); err != nil {
		c.Error(fmt.Errorf("failed to store code to redis: %w", err))
		return "存储验证码失败", false
	}

	bodyTmpl, err := template.ParseFiles(c.Env.Path.EmailBody)
	if err != nil {
		c.Error(fmt.Errorf("failed to read email body template: %w", err))
		return "读取文件内容模板失败", false
	}

	cfg := c.Env.SMTP
	var buf bytes.Buffer
	if _, err = buf.Write([]byte(fmt.Sprintf(
		"From: \"%s\" <%s>\r\nTo: %s\r\nSubject: %s\r\nMIME-Version: 1.0\r\nContent-Type: text/html; charset=\"UTF-8\"\r\n\r\n",
		cfg.From, cfg.Mail, i.Email, cfg.Subject,
	))); err != nil {
		c.Error(fmt.Errorf("failed to write email header: %w", err))
		return "写入邮件头部失败", false
	}

	if err := bodyTmpl.Execute(&buf, &struct {
		VerifyCode string
		Expiration string
		IP         string
	}{
		code,
		time.Now().Add(c.Env.Duration.VerifyCode).Local().Format("2006-01-02 15:04:05"),
		c.ClientIP(),
	}); err != nil {
		c.Error(fmt.Errorf("failed to execute email body template: %w", err))
		return "解析邮件内容失败", false
	}

	if err := smtp.SendMail(
		cfg.Server+":"+cfg.Port,
		smtp.PlainAuth("", cfg.Mail, cfg.Pass, cfg.Server),
		cfg.Mail,
		[]string{i.Email},
		buf.Bytes(),
	); err != nil {
		c.Error(fmt.Errorf("failed to send email: %w", err))
		return "邮件发送失败", false
	}

	return "邮件发送成功", true
}
