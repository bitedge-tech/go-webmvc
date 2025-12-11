package logger

import (
	config "go-webmvc/config"
	"os"
	"path/filepath"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	Log   *zap.Logger
	Sugar *zap.SugaredLogger
)

func Init(cfg config.LogConfig) error {
	// 创建日志目录
	if err := os.MkdirAll(filepath.Dir(cfg.File.Filename), 0755); err != nil {
		return err
	}

	// 设置日志级别
	level := getZapLevel(cfg.Level)

	// 创建编码器
	encoder := getEncoder(cfg.Format)

	// 创建多个输出核心
	cores := []zapcore.Core{}

	// 文件输出
	if cfg.Output == "file" || cfg.Output == "both" {
		fileCore := zapcore.NewCore(encoder, getFileWriter(cfg.File), level)
		cores = append(cores, fileCore)
	}

	// 控制台输出
	if cfg.Output == "stdout" || cfg.Output == "both" {
		consoleCore := zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), level)
		cores = append(cores, consoleCore)
	}

	// 创建 Log
	core := zapcore.NewTee(cores...)
	Log = zap.New(core,
		zap.AddCaller(),
		zap.AddCallerSkip(1),
		zap.AddStacktrace(zap.ErrorLevel),
	)
	Sugar = Log.Sugar()

	return nil
}

func getEncoder(format string) zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	if format == "console" {
		encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		return zapcore.NewConsoleEncoder(encoderConfig)
	}
	return zapcore.NewJSONEncoder(encoderConfig)
}

func getZapLevel(level string) zapcore.Level {
	switch level {
	case "debug":
		return zapcore.DebugLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	default:
		return zapcore.InfoLevel
	}
}

func getFileWriter(cfg config.LogFileConfig) zapcore.WriteSyncer {
	return zapcore.AddSync(&lumberjack.Logger{
		Filename:   cfg.Filename,
		MaxSize:    cfg.MaxSize,
		MaxBackups: cfg.MaxBackups,
		MaxAge:     cfg.MaxAge,
		Compress:   cfg.Compress,
	})
}

// 便捷方法
func Debug(msg string, fields ...zap.Field) {
	Log.Debug(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	Log.Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	Log.Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	Log.Error(msg, fields...)
}

func Sync() {
	_ = Log.Sync()
}
