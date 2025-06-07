package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"strings"
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
	Host        string
	Password    string // Changed from Pass to Password for consistency with YAML and env
	Db          int
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

	// Override with environment variables where applicable
	overrideWithEnvVars()

	// Validate required secrets, especially in production
	if server.RunMode == "release" {
		validateSecrets()
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

	// Enable environment variable overriding
	conf.AutomaticEnv()
	conf.SetEnvKeyReplacer(strings.NewReplacer(".", "_")) // e.g., app.jwtsecret becomes APP_JWTSECRET

	// Attempt to read the main config file
	if err := conf.ReadInConfig(); err != nil {
		// If the primary config file is not found, try to load default (e.g. local)
		// This can be useful if APP_CONF is not explicitly set for a specific environment
		// and we want to fall back to a default like 'local.yaml'
		fmt.Printf("Warning: Could not read config file %s: %v. Trying default 'local'.\n", configName, err)
		conf.SetConfigName("local") // default to local if specific one not found
		if err := conf.ReadInConfig(); err != nil {
			panic(fmt.Errorf("fatal error config file %s or default local.yaml not found: %w", configName, err))
		}
	}

	// Load .env file if present, Viper 1.7.0+ supports this directly
	// For older versions, a library like godotenv might be used before Viper initialization.
	// This example assumes Viper handles it or it's handled externally.
	// For robustness, explicitly load .env if you want to ensure it.
	// Example using godotenv (would require `import "github.com/joho/godotenv"`):
	// if _, err := os.Stat(".env"); err == nil {
	// 	 if err := godotenv.Load(); err != nil {
	// 	 	 fmt.Println("Error loading .env file")
	// 	 }
	// }


	return conf
}

// overrideWithEnvVars fetches values from environment variables if they are set,
// overriding values loaded from YAML (or their defaults if not in YAML).
func overrideWithEnvVars() {
	// App secrets
	if envVal := os.Getenv("APP_JWT_SECRET"); envVal != "" {
		app.JWTSecret = envVal
	}
	if envVal := os.Getenv("APP_APP_SECRET"); envVal != "" {
		app.AppSecret = envVal
	}
	if envVal := os.Getenv("APP_APP_KEY"); envVal != "" {
		app.AppKey = envVal
	}

	// Database configuration
	if envVal := os.Getenv("DB_HOST"); envVal != "" {
		database.Host = envVal
	}
	if envVal := os.Getenv("DB_PORT"); envVal != "" {
		// Consider parsing to int if viper doesn't do it automatically
		// For simplicity, assuming viper handles string if struct field is string,
		// or direct os.Getenv for manual parsing if needed.
		// Here, Port is int, so direct unmarshal or viper.GetInt is better.
		// This manual override is mostly for string fields or complex cases.
		// For viper.GetInt: database.Port = viper.GetInt("database.port")
	}
	if envVal := os.Getenv("DB_NAME"); envVal != "" {
		database.DbName = envVal
	}
	if envVal := os.Getenv("DB_USER"); envVal != "" {
		database.User = envVal
	}
	if envVal := os.Getenv("DB_PASSWORD"); envVal != "" {
		database.Password = envVal
	}
	if envVal := os.Getenv("DB_CHARSET"); envVal != "" {
		database.Charset = envVal
	}

	// Redis configuration
	if envVal := os.Getenv("REDIS_HOST"); envVal != "" {
		redis.Host = envVal
	}
	if envVal := os.Getenv("REDIS_PASSWORD"); envVal != "" {
		redis.Password = envVal
	}
	if envVal := os.Getenv("REDIS_DB"); envVal != "" {
		// Similar to DB_PORT, consider viper.GetInt("redis.db")
		// For simplicity, if viper unmarshalled it and it's 0, this won't override unless env is non-empty.
	}
}

// validateSecrets checks if essential secrets are provided, panicking if not.
func validateSecrets() {
	missingSecrets := []string{}
	if app.JWTSecret == "" {
		missingSecrets = append(missingSecrets, "APP_JWT_SECRET")
	}
	if app.AppSecret == "" {
		missingSecrets = append(missingSecrets, "APP_APP_SECRET")
	}
	if app.AppKey == "" {
		missingSecrets = append(missingSecrets, "APP_APP_KEY")
	}
	// In prod, database password should likely be required.
	// For local, it might come from local.yaml directly if not overridden.
	if database.Password == "" && server.RunMode == "release" { // Check only in release mode
		// This check depends on whether local DB password in local.yaml is acceptable
		// or if it *must* be from env var in prod.
		// Assuming DB_PASSWORD must be set via env for prod.
		appConfig := GetConfig(os.Getenv("APP_CONF")) // Re-get to check actual value post unmarshal
		if appConfig.GetString("database.password") == "" { // if it was empty in yaml too
			missingSecrets = append(missingSecrets, "DB_PASSWORD")
		}
	}
	// Add similar check for REDIS_PASSWORD if it's always required in prod
	// if redis.Password == "" && server.RunMode == "release" {
	// 	  appConfig := GetConfig(os.Getenv("APP_CONF"))
	//    if appConfig.GetString("redis.password") == "" { // if it was empty in yaml too
	// 	      missingSecrets = append(missingSecrets, "REDIS_PASSWORD")
	//    }
	// }


	if len(missingSecrets) > 0 {
		panic(fmt.Sprintf("FATAL ERROR: Required secret(s) not set in environment for production: %s", strings.Join(missingSecrets, ", ")))
	}
}
