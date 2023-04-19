package conf

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"path/filepath"
	"time"
)

func InitLogger() *zap.SugaredLogger {
	LevelEnabler := zapcore.DebugLevel
	if !viper.GetBool("mode.develop") {
		LevelEnabler = zapcore.InfoLevel
	}
	core := zapcore.NewCore(getEncoder(), zapcore.NewMultiWriteSyncer(getWriteSyncer(), zapcore.AddSync(os.Stdout)), LevelEnabler)
	return zap.New(core).Sugar()
}

func getEncoder() zapcore.Encoder {
	encodeConfig := zap.NewProductionEncoderConfig()
	encodeConfig.TimeKey = "time"
	// leve 字段名转成大写
	encodeConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	// 格式化时间
	encodeConfig.EncodeTime = func(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(t.Local().Format(time.DateTime))
	}
	return zapcore.NewJSONEncoder(encodeConfig)
}

func getWriteSyncer() zapcore.WriteSyncer {
	setSeparator := string(filepath.Separator)
	setRootDir, _ := os.Getwd()
	setLogFilePath := setRootDir + setSeparator + "log" + setSeparator + time.Now().Format(time.DateOnly) + ".txt"
	//fmt.Println(setLogFilePath)

	lumberjackSyncer := &lumberjack.Logger{
		Filename:   setLogFilePath,
		MaxSize:    viper.GetInt("log.MaxSize"),
		MaxBackups: viper.GetInt("log.MaxBackups"),
		MaxAge:     viper.GetInt("log.MaxAge"),
		Compress:   true,
	}

	return zapcore.AddSync(lumberjackSyncer)
}
