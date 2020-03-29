package logger

/*
对 zap logger 的二次封装
*****************
import "github.com/jessun2017/gold/logger"
...
var l *logger.Logger

func init() {
	l = logger.NewLogger("servername", nil)
}

*****************
*/

import (
	"os"
	"time"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/jessun2017/gold/constant"
)

// Logger 设置格式别名
type Logger = zap.Logger

// Config 日志设置的二次封装
type Config struct {
	FileConfig     lumberjack.Logger
	LogLevelConfig zapcore.Level
}

// LoggerPreset 默认设置
var LoggerPreset = Config{
	FileConfig: lumberjack.Logger{
		// 默认日志文件位置，当前目录下的 logs 目录
		Filename: "/tmp/" + time.Now().Format(constant.TimeLogFmt) + ".log",
		// 单个文件最大尺寸，单位 MB
		MaxSize: 256,
		// 日志文件最多保存多少个备份
		MaxBackups: 30,
		// 文件最多保存多少天
		MaxAge: 14,
		// 是否压缩
		Compress: true,
	},
	LogLevelConfig: zapcore.InfoLevel,
}

// NewLogger 创建新的 Logger 对象，默认存储位置在 /tmp/[serviceName]/ 下
func NewLogger(serviceName string, set *Config) *zap.Logger {
	if set == nil { // 使用默认设置 LoggerPreset
		set = &LoggerPreset
		set.FileConfig.Filename = "/tmp/" + serviceName + "/" +
			time.Now().Format(constant.TimeLogFmt) + ".log"
	}

	// 设置日志级别
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(set.LogLevelConfig)
	//公用编码器
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "lineNum",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,  // 编码器
		EncodeTime:     zapcore.ISO8601TimeEncoder,     // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder, //
		EncodeCaller:   zapcore.FullCallerEncoder,      // 全路径编码器
		EncodeName:     zapcore.FullNameEncoder,
	}
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig), // 编码器配置
		zapcore.NewMultiWriteSyncer(
			zapcore.AddSync(os.Stdout),
			zapcore.AddSync(&set.FileConfig),
		), // 打印到控制台和文件
		atomicLevel, // 日志级别
	)
	return zap.New(core, zap.AddCaller(), zap.Development(),
		zap.Fields(zap.String("ServiceName", serviceName)))
}
