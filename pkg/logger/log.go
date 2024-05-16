package logger

import (
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"log"
	"time"
)

var Logger *zap.Logger

func init() {
	config := zapcore.EncoderConfig{
		MessageKey:  "msg",
		LevelKey:    "level",
		EncodeLevel: zapcore.CapitalLevelEncoder,
		TimeKey:     "ts",
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006-01-02 15:04:05"))
		},
		CallerKey:    "file",
		EncodeCaller: zapcore.ShortCallerEncoder,
		EncodeDuration: func(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendInt64(int64(d) / 10000000)
		},
	}
	encoder := zapcore.NewConsoleEncoder(config)
	FileFormat, saveType, LogLevel := "%Y%m%d", "one", "info"

	fmt.Println(222222222222222222)
	fmt.Println(viper.AllSettings())
	if viper.IsSet("log.level") {
		fmt.Println(111111111111111111)
		LogLevel = viper.GetString("log.level")
		fmt.Println(LogLevel)
	}
	if viper.IsSet("log.file_format") {
		FileFormat = viper.GetString("log.file_format")
	}
	if viper.IsSet("log.file_type") {
		saveType = viper.GetString("log.file_type")
	}

	logLevel := zap.DebugLevel
	switch LogLevel {
	case "debug":
		logLevel = zap.DebugLevel
	case "info":
		logLevel = zap.InfoLevel
	case "error":
		logLevel = zap.ErrorLevel
	case "warning":
		logLevel = zap.WarnLevel
	default:
		logLevel = zap.InfoLevel
	}

	fmt.Println(logLevel)

	switch saveType {
	case "level":
		Logger = getLevelLogger(encoder, logLevel, FileFormat)
	default:
		Logger = getOnceLogger(encoder, logLevel, FileFormat)
	}
}

func getOnceLogger(encoder zapcore.Encoder, level zapcore.Level, format string) *zap.Logger {
	fmt.Println(3333333333333)
	infoWriter := getLoggerWriter("./var/log/run_log", format)
	infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl == zapcore.DebugLevel && lvl >= level
	})
	core := zapcore.NewTee(
		zapcore.NewCore(encoder, zapcore.AddSync(infoWriter), infoLevel),
	)
	return zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.WarnLevel))
}

func getLevelLogger(encoder zapcore.Encoder, level zapcore.Level, format string) *zap.Logger {
	infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl == zapcore.InfoLevel && lvl >= level
	})
	debugLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl == zapcore.DebugLevel && lvl >= level
	})
	errorLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl == zapcore.ErrorLevel && lvl >= level
	})
	warnLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl == zapcore.WarnLevel && lvl >= level
	})

	infoWriter := getLoggerWriter("./var/log/info", format)
	debugWriter := getLoggerWriter("./var/log/debug", format)
	errorWriter := getLoggerWriter("./var/log/error", format)
	warnWriter := getLoggerWriter("./var/log/warn", format)

	// 创建Logger实例
	core := zapcore.NewTee(
		zapcore.NewCore(encoder, zapcore.AddSync(debugWriter), debugLevel),
		zapcore.NewCore(encoder, zapcore.AddSync(infoWriter), infoLevel),
		zapcore.NewCore(encoder, zapcore.AddSync(errorWriter), errorLevel),
		zapcore.NewCore(encoder, zapcore.AddSync(warnWriter), warnLevel),
	)
	return zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.WarnLevel))
}

func getLoggerWriter(file string, format string) io.Writer {
	// 生成rotate
	hook, err := rotatelogs.New(
		file+format+".log",
		rotatelogs.WithLinkName(file),
		// 保存天数
		rotatelogs.WithMaxAge(time.Hour*24*30),
		// 切割频率 每天
		rotatelogs.WithRotationTime(time.Hour*24),
	)
	if err != nil {
		log.Println("日志文件创建失败")
		panic(err)
	}
	return hook
}

func Debug(format string, v ...interface{}) {
	Logger.Sugar().Debugf(format, v...)
}

func Info(format string, v ...interface{}) {
	Logger.Sugar().Infof(format, v...)
}

func Error(format string, v ...interface{}) {
	Logger.Sugar().Errorf(format, v...)
}
