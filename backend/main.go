package main

import (
	"fmt"
	"log"
	"time"

	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss"
	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss/credentials"
	"github.com/kelseyhightower/envconfig"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {

	ec := EnvConfig{
		Port: "8080",
		OSS: OSSConfig{
			Region: "cn-hangzhou",
		},
		Path: PathConfig{
			RBAC:        "/app/rbac.yaml",
			EmailBody:   "/app/email.html",
			EmailHeader: "/app/email.txt",
		},
		MYSQL: MySQLConfig{Host: "127.0.0.1", Port: "3306"},
		Redis: RedisConfig{Host: "127.0.0.1", Port: "6379"},
		Duration: DurationConfig{
			RefreshToken: 7 * 24 * time.Hour,
			AccessToken:  10 * time.Minute,
			VerifyCode:   10 * time.Minute,
		},
	}
	if err := envconfig.Process("", &ec); err != nil {
		log.Fatalf("failed to load config: %v\n", err)
	}

	db, err := gorm.Open(mysql.Open(fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		ec.MYSQL.User,
		ec.MYSQL.Pass,
		ec.MYSQL.Host,
		ec.MYSQL.Port,
		ec.MYSQL.DB,
	)), &gorm.Config{TranslateError: true, Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		log.Fatalf("failed to connect database: %v\n", err)
	}
	if err = db.AutoMigrate(Tables...); err != nil {
		log.Fatalf("failed to migrate tables: %v\n", err)
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     ec.Redis.Host + ":" + ec.Redis.Port,
		Password: ec.Redis.Pass,
		DB:       ec.Redis.DB,
	})

	ossClient := oss.NewClient(oss.LoadDefaultConfig().WithCredentialsProvider(credentials.NewEnvironmentVariableCredentialsProvider()))

	cfg := Config{
		Env: &ec,
		DB:  db,
		RDB: rdb,
		OSS: ossClient,
	}

	if err := GetRouter(&cfg).Run(":" + ec.Port); err != nil {
		log.Fatalf("failed to start server: %v\n", err)
	}

}
