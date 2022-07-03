package initialize

import (
	"api-platform/global"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Logger() {
	encoder := Encoder()
	syncer := WriteSyncer()
	core := zapcore.NewCore(encoder, syncer, zapcore.DebugLevel)
	logger := zap.New(core, zap.AddCaller())
	global.Logger = logger
}

func Encoder() zapcore.Encoder {
	config := zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	encoder := zapcore.NewJSONEncoder(config)

	return encoder
}

func WriteSyncer() zapcore.WriteSyncer {
	//file, _ := os.OpenFile("./test.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0744)
	//writeSyncer := zapcore.AddSync(file)
	zapConfig := global.Config.Zap
	lumberJackLogger := &lumberjack.Logger{
		Filename:   zapConfig.LoggerFile,
		MaxSize:    zapConfig.MaxSize,
		MaxBackups: zapConfig.MaxBackups,
		MaxAge:     zapConfig.MaxAge,
		Compress:   zapConfig.Compress,
	}
	return zapcore.AddSync(lumberJackLogger)
}
