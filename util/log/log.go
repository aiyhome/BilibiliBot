package log

import (
	"fmt"
	"github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

var DefaultLogger *zap.SugaredLogger

func checkPathExist(path string) bool {
	_, err := os.Stat(path)
	if nil == err {
		return true
	}

	return false
}

func timeEncodeFunc(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
}

func Init(directory string, filePattern string, fileLink string, consoleOnly bool) {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = timeEncodeFunc
	encoderConfig.TimeKey = "time"
	encoder := zapcore.NewConsoleEncoder(encoderConfig)

	var writer zapcore.WriteSyncer
	var logLevel zapcore.Level
	if !consoleOnly {
		// 获取当前工作目录
		pwd, err := os.Getwd()
		fmt.Println(pwd)
		if nil != err {
			panic(err)
		}

		// 创建日志目录
		if !checkPathExist(directory) {
			fmt.Printf("log dir %s does not exist, create\n", directory)
			os.Mkdir(directory, os.ModePerm)
		}

		routateWriter, err := rotatelogs.New(
			pwd+"/"+filePattern,
			rotatelogs.WithLinkName(fileLink),
			rotatelogs.WithMaxAge(24*time.Hour*30),
			rotatelogs.WithRotationTime(24*time.Hour),
		)

		if nil != err {
			panic(err)
		}

		logLevel = zapcore.InfoLevel
		writer = zapcore.AddSync(routateWriter)

	} else {
		logLevel = zapcore.DebugLevel
		writer = zapcore.AddSync(os.Stdout)
	}

	logCore := zapcore.NewCore(encoder, writer, logLevel)
	// logger := zap.New(logCore, zap.AddCaller())
	logger := zap.New(logCore)
	DefaultLogger = logger.Sugar()

	DefaultLogger.Info("log initialized")
}

// Info ...
func Info(args ...interface{}) {
	DefaultLogger.Info(args...)
}

// Debug ...
func Debug(args ...interface{}) {
	DefaultLogger.Debug(args...)
}

// Warn ...
func Warn(args ...interface{}) {
	DefaultLogger.Warn(args...)
}

// Error ...
func Error(args ...interface{}) {
	DefaultLogger.Error(args...)
}

// Panic ...
func Panic(args ...interface{}) {
	DefaultLogger.Panic(args...)
}

// DPanic ...
func DPanic(args ...interface{}) {
	DefaultLogger.DPanic(args...)
}

// Fatal ...
func Fatal(args ...interface{}) {
	DefaultLogger.Fatal(args...)
}
