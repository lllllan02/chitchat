package logger

import (
	"fmt"
	"log"
	"os"
	"time"
)

// 日志级别
const (
	DEBUG   = "DEBUG"
	INFO    = "INFO"
	WARNING = "WARNING"
	ERROR   = "ERROR"
	FATAL   = "FATAL"
)

var (
	defaultLogger *log.Logger
	logLevel      string
	logFile       *os.File
)

// 初始化日志
func Init(level string, filePath string) error {
	logLevel = level

	// 如果指定了日志文件路径，则创建日志文件
	if filePath != "" {
		var err error
		// 确保日志文件所在目录存在
		err = os.MkdirAll(filePath[:len(filePath)-len(fmt.Sprintf("%s.log", time.Now().Format("2006-01-02")))], os.ModePerm)
		if err != nil {
			return err
		}

		// 打开日志文件，如果不存在则创建
		logFile, err = os.OpenFile(
			fmt.Sprintf("%s%s.log", filePath, time.Now().Format("2006-01-02")),
			os.O_CREATE|os.O_WRONLY|os.O_APPEND,
			0666,
		)
		if err != nil {
			return err
		}

		// 创建日志记录器
		defaultLogger = log.New(logFile, "", log.LstdFlags)
	} else {
		// 如果没有指定日志文件，则使用标准输出
		defaultLogger = log.New(os.Stdout, "", log.LstdFlags)
	}

	return nil
}

// 关闭日志文件
func Close() {
	if logFile != nil {
		logFile.Close()
	}
}

// Debug 调试日志
func Debug(format string, v ...interface{}) {
	if logLevel == DEBUG {
		logMessage(DEBUG, format, v...)
	}
}

// Info 信息日志
func Info(format string, v ...interface{}) {
	if logLevel == DEBUG || logLevel == INFO {
		logMessage(INFO, format, v...)
	}
}

// Warning 警告日志
func Warning(format string, v ...interface{}) {
	if logLevel == DEBUG || logLevel == INFO || logLevel == WARNING {
		logMessage(WARNING, format, v...)
	}
}

// Error 错误日志
func Error(format string, v ...interface{}) {
	if logLevel == DEBUG || logLevel == INFO || logLevel == WARNING || logLevel == ERROR {
		logMessage(ERROR, format, v...)
	}
}

// Fatal 致命错误日志，记录后终止程序
func Fatal(format string, v ...interface{}) {
	logMessage(FATAL, format, v...)
	os.Exit(1)
}

// 记录日志
func logMessage(level, format string, v ...interface{}) {
	if defaultLogger == nil {
		defaultLogger = log.New(os.Stdout, "", log.LstdFlags)
	}

	logContent := fmt.Sprintf("[%s] %s", level, fmt.Sprintf(format, v...))
	defaultLogger.Println(logContent)
}
