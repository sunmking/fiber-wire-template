package log

import (
	"context"
	"fiber-wire-template/pkg/config"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"time"
)

type contextKey int

const (
	requestIDKey contextKey = iota
	correlationIDKey
)

type Logger struct {
	*zap.Logger
}

func NewLog(conf *config.Config) *Logger {
	return initZap(conf)
}

func initZap(conf *config.Config) *Logger {
	// log address "out.log" User-defined
	lp := conf.LogCnf.LogSavePath + "/" + conf.LogCnf.LogFileName
	lv := conf.LogCnf.LogLevel
	var level zapcore.Level
	//debug<info<warn<error<fatal<panic
	switch lv {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
	default:
		level = zap.InfoLevel
	}
	hook := lumberjack.Logger{
		Filename:   lp,                     // Log file path
		MaxSize:    conf.LogCnf.MaxSize,    // Maximum size unit for each log file: M
		MaxBackups: conf.LogCnf.MaxBackups, // The maximum number of backups that can be saved for log files
		MaxAge:     conf.LogCnf.MaxAge,     // Maximum number of days the file can be saved
		Compress:   conf.LogCnf.Compress,   // Compression or not
	}

	var encoder zapcore.Encoder
	if conf.LogCnf.LogEncoding == "console" {
		encoder = zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
			TimeKey:        "ts",
			LevelKey:       "level",
			NameKey:        "Logger",
			CallerKey:      "caller",
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseColorLevelEncoder,
			EncodeTime:     timeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.FullCallerEncoder,
		})
	} else {
		encoder = zapcore.NewJSONEncoder(zapcore.EncoderConfig{
			TimeKey:        "ts",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			FunctionKey:    zapcore.OmitKey,
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     zapcore.EpochTimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		})
	}
	core := zapcore.NewCore(
		encoder,
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)), // Print to console and file
		level,
	)
	if conf.AppCnf.DEBUG == true {
		return &Logger{zap.New(core, zap.Development(), zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))}
	}
	return &Logger{zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))}

}

func timeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	//enc.AppendString(t.Format("2006-01-02 15:04:05"))
	enc.AppendString(t.Format("2006-01-02 15:04:05.000000000"))
}

// With returns a logger based off the root logger and decorates it with the given context and arguments.
//
// If the context contains request ID and/or correlation ID information (recorded via WithRequestID()
// and WithCorrelationID()), they will be added to every log message generated by the new logger.
//
// The arguments should be specified as a sequence of name, value pairs with names being strings.
// The arguments will also be added to every log message generated by the logger.
func (l *Logger) With(ctx context.Context, args ...zap.Field) *Logger {
	if ctx != nil {
		if id, ok := ctx.Value(requestIDKey).(string); ok {
			args = append(args, zap.String("request_id", id))
		}
		if id, ok := ctx.Value(correlationIDKey).(string); ok {
			args = append(args, zap.String("correlation_id", id))
		}
	}
	if len(args) > 0 {
		return &Logger{Logger: l.WithLazy(args...)}
	}
	return l
}

// WithRequest returns a context which knows the request ID and correlation ID in the given request.
func WithRequest(ctx context.Context, req *fiber.Ctx) context.Context {
	id := getRequestID(req)
	if id == "" {
		id = uuid.New().String()
	}
	ctx = context.WithValue(ctx, requestIDKey, id)
	if id := getCorrelationID(req); id != "" {
		ctx = context.WithValue(ctx, correlationIDKey, id)
	}
	return ctx
}

// getCorrelationID extracts the correlation ID from the HTTP request
func getCorrelationID(c *fiber.Ctx) string {
	return c.Get("X-")
}

// getRequestID extracts the correlation ID from the HTTP request
func getRequestID(c *fiber.Ctx) string {
	return c.Get(fiber.HeaderXRequestID)
}
