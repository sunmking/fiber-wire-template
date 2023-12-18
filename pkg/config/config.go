package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"time"
)

type App struct {
	// 是否开启调试模式
	DEBUG bool
	// JWT密钥
	JWTSecret string
	JWTExpire time.Duration
	// 应用密钥
	AppKey    string
	AppSecret string
}

var app App

type Server struct {
	// 监听端口
	Port int
	// 运行模式
	// 0: 单进程模式
	// 1: 多进程模式
	// 2: 多进程模式+热更新模式
	RunMode      string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var server Server

// Log 日志配置
type Log struct {
	// 日志编码
	LogEncoding string
	// 日志保存路径

	LogSavePath string
	// 日志文件名
	LogFileName string
	// 日志文件最大大小
	MaxSize    int
	MaxAge     int
	MaxBackups int
	Compress   bool
	// 日志级别
	LogLevel string
}

var log Log

// Database 数据库配置
type Database struct {
	Type        string
	Host        string
	Port        int
	User        string
	Password    string
	DbName      string
	Charset     string
	TablePrefix string
}

var database Database

// Redis Redis配置
type Redis struct {
	Host string
	Pass string
	Db   int
	// 最大空闲连接数
	MaxIdle int
	// 最大活动连接数
	MaxActive int
	// 空闲连接超时时间
	IdleTimeout time.Duration
}

var redis Redis

type Config struct {
	AppCnf      *App
	ServerCnf   *Server
	LogCnf      *Log
	DatabaseCnf *Database
	RedisCnf    *Redis
}

func NewConfig(configName string) *Config {
	envCnf := os.Getenv("APP_CONF")
	if envCnf == "" {
		envCnf = configName
	}
	fmt.Println("envCnf:", envCnf)
	conf := GetConfig(envCnf)

	if err := conf.UnmarshalKey("server", &server); err != nil {
		panic(err)
	}
	if err := conf.UnmarshalKey("app", &app); err != nil {
		panic(err)
	}
	if err := conf.UnmarshalKey("log", &log); err != nil {
		panic(err)
	}
	if err := conf.UnmarshalKey("database", &database); err != nil {
		panic(err)
	}
	if err := conf.UnmarshalKey("redis", &redis); err != nil {
		panic(err)
	}

	return &Config{
		AppCnf:      &app,
		ServerCnf:   &server,
		LogCnf:      &log,
		DatabaseCnf: &database,
		RedisCnf:    &redis,
	}
}

func GetConfig(configName string) *viper.Viper {
	var conf = viper.New()
	conf.SetConfigType("yaml")
	conf.AddConfigPath("./config")
	conf.SetConfigName(configName)
	if err := conf.ReadInConfig(); err != nil {
		panic(err)
	}
	return conf
}
