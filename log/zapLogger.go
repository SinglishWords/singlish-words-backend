package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"singlishwords/config"
)

func initLogger() *zap.SugaredLogger {
	encoder := getEncoder()
	writeSyncer := getLogWriter()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)

	// zap.AddCaller()  添加将调用函数信息记录到日志中的功能。
	logger := zap.New(core, zap.AddCaller())
	return logger.Sugar()
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder // 修改时间编码器

	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
	//return zapcore.NewJSONEncoder(encoderConfig)
}

func getLogWriter() zapcore.WriteSyncer {
	if config.Log.Info != "stdout" {
		file, _ := os.Create(config.Log.Info)
		return zapcore.AddSync(file)
	}
	return zapcore.NewMultiWriteSyncer(os.Stdout)
}
