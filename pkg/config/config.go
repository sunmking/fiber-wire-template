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
	appKey    string
	appSecret string
}

var AppCnf = &App{}

type Server struct {
	// 监听端口
	Port string
	// 运行模式
	// 0: 单进程模式
	// 1: 多进程模式
	// 2: 多进程模式+热更新模式
	RunMode      string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var ServerCnf = &Server{}

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

var LogCnf = &Log{}

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

var DatabaseCnf = &Database{}

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

var RedisCnf = &Redis{}

type Config struct {
	AppCnf      *App
	ServerCnf   *Server
	LogCnf      *Log
	DatabaseCnf *Database
	RedisCnf    *Redis
}

func NewConfig(p string) *Config {
	envCnf := os.Getenv("APP_CONF")
	if envCnf == "" {
		envCnf = p
	}
	fmt.Println("envCnf:", envCnf)
	conf := GetConfig(envCnf)

	if err := conf.Unmarshal(ServerCnf); err != nil {
		panic(err)
	}
	if err := conf.Unmarshal(AppCnf); err != nil {
		panic(err)
	}
	if err := conf.Unmarshal(LogCnf); err != nil {
		panic(err)
	}
	if err := conf.Unmarshal(DatabaseCnf); err != nil {
		panic(err)
	}
	if err := conf.Unmarshal(RedisCnf); err != nil {
		panic(err)
	}

	return &Config{
		AppCnf:      AppCnf,
		ServerCnf:   ServerCnf,
		LogCnf:      LogCnf,
		DatabaseCnf: DatabaseCnf,
		RedisCnf:    RedisCnf,
	}
}

func GetConfig(path string) *viper.Viper {
	var conf = viper.New()
	conf.SetConfigType("yaml")
	conf.AddConfigPath(".")
	conf.SetConfigName(path)
	if err := conf.ReadInConfig(); err != nil {
		panic(err)
	}
	return conf
}
