package main

import "time"

type EnvConfig struct {
	Port      string
	JWTKey    string
	JWTIssuer string
	OSS       OSSConfig
	Path      PathConfig
	Duration  DurationConfig
	MYSQL     MySQLConfig
	SMTP      SMTPConfig
	Redis     RedisConfig
	Cookie    CookieConfig
}

type OSSConfig struct {
	Endpoint  string
	KeyID     string
	KeySecret string
	Bucket    string
	Region    string
}

type PathConfig struct {
	RBAC        string
	Enums       string
	EmailBody   string
	EmailHeader string
}

type DurationConfig struct {
	RefreshToken time.Duration
	AccessToken  time.Duration
	VerifyCode   time.Duration
}

type MySQLConfig struct {
	User string
	Pass string
	Host string
	Port string
	DB   string
}

type RedisConfig struct {
	Host string
	Port string
	Pass string
	DB   int
}

type SMTPConfig struct {
	Server  string
	Port    string
	Mail    string
	Pass    string
	From    string
	Subject string
}

type CookieConfig struct {
	Path     string
	Domain   string
	Secure   bool
	HTTPOnly bool
}
