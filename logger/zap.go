package logger

/*
	zap logger 的二次封装
*/

import (
	"os"
	"time"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/jessun2017/gold/constant"
)

var serName string

type LoggerSet struct {
	Hook    lumberjack.Logger
	ZapCore zapcore.Core
	Level   zapcore.Level
}

// LoggerPreset 默认设置
var LoggerPreset = LoggerSet{
	Hook: lumberjack.Logger{
		// 默认日志文件位置，当前目录下的 logs 目录
		Filename: "/tmp/" + serName + "/" + time.Now().Format(constant.TimeLogFmt) + ".log",
		// 单个文件最大尺寸，单位 MB
		MaxSize: 256,
		// 日志文件最多保存多少个备份
		MaxBackups: 30,
		// 文件最多保存多少天
		MaxAge: 14,
		// 是否压缩
		Compress: true,
	},
	Level: zapcore.InfoLevel,
}

func NewLogger(serviceName string) *zap.Logger {
	serName = serviceName

	// 设置日志级别
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(LoggerPreset.Level)
	//公用编码器
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "linenum",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,  // 编码器
		EncodeTime:     zapcore.ISO8601TimeEncoder,     // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder, //
		EncodeCaller:   zapcore.FullCallerEncoder,      // 全路径编码器
		EncodeName:     zapcore.FullNameEncoder,
	}
	LoggerPreset.ZapCore = zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),                                                        // 编码器配置
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&LoggerPreset.Hook)), // 打印到控制台和文件
		atomicLevel, // 日志级别
	)

	return zap.New(LoggerPreset.ZapCore, zap.AddCaller(), zap.Development(),
		zap.Fields(zap.String("ServiceName", serviceName)))
}
