package logger

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"log"
	"os"
	"time"
)

var Logger *zap.Logger

func InitLogger() *zap.Logger {
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

	if viper.IsSet("log.level") {
		LogLevel = viper.GetString("log.level")
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

	switch saveType {
	case "level":
		Logger = getLevelLogger(encoder, FileFormat)
	default:
		Logger = getOnceLogger(encoder, logLevel, FileFormat)
	}
	defer Logger.Sync()

	return Logger
}

func getOnceLogger(encoder zapcore.Encoder, level zapcore.Level, format string) *zap.Logger {
	logWriter := getLoggerWriter("./var/log/run_log", format)
	setLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return level <= lvl
	})
	core := zapcore.NewTee(
		zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(zapcore.AddSync(logWriter), os.Stdout), setLevel),
	)
	return zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.WarnLevel))
}

func getLevelLogger(encoder zapcore.Encoder, format string) *zap.Logger {
	infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl == zapcore.InfoLevel
	})
	debugLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl == zapcore.DebugLevel
	})
	errorLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl == zapcore.ErrorLevel
	})
	warnLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl == zapcore.WarnLevel
	})

	infoWriter := getLoggerWriter("./var/log/info", format)
	debugWriter := getLoggerWriter("./var/log/debug", format)
	errorWriter := getLoggerWriter("./var/log/error", format)
	warnWriter := getLoggerWriter("./var/log/warn", format)

	// 创建Logger实例
	core := zapcore.NewTee(
		zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(zapcore.AddSync(debugWriter), os.Stdout), debugLevel),
		zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(zapcore.AddSync(infoWriter), os.Stdout), infoLevel),
		zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(zapcore.AddSync(errorWriter), os.Stdout), errorLevel),
		zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(zapcore.AddSync(warnWriter), os.Stdout), warnLevel),
	)
	return zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.WarnLevel))
}

func getLoggerWriter(file string, format string) io.Writer {
	// 生成rotate
	saveDays := viper.GetInt("max_save_days")
	duration := time.Hour * 24 * time.Duration(saveDays)
	hook, err := rotatelogs.New(
		file+format+".log",
		rotatelogs.WithLinkName(file),
		// 保存天数
		rotatelogs.WithMaxAge(duration),
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

func Warn(format string, v ...interface{}) {
	Logger.Sugar().Warnf(format, v...)
}

func Error(format string, v ...interface{}) {
	Logger.Sugar().Errorf(format, v...)
}
