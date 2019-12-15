package logger

import (
	"encoding/json"
	"go.uber.org/zap"
	"log"
)

var logger *zap.Logger
var sugar *zap.SugaredLogger

var JSONCONFIG string = `	{
		"level": "info",
		"encoding": "json",
		"outputPaths": ["stdout"],
		"errorOutputPaths": ["stderr"],
		"initialFields": {},
		"encoderConfig": {
		"messageKey": "message",
			"levelKey": "level",
			"levelEncoder": "lowercase",
			"timeKey": "timestamp",
			"timeEncoder": "iso8601"
	}
	}`

func InitLogger() *zap.SugaredLogger {
	//поправить путь

	//rawJSON, err := ioutil.ReadFile("./logger/logger-config.json")
	//if err != nil {
	//	log.Println(err)
	//	logger, err = zap.NewProduction()
	//	if err != nil {
	//		log.Panicln(err)
	//	}
	//	sugar := logger.Sugar()
	//	sugar.Warn("using default production logger")
	//	return sugar
	//}

	var cfg zap.Config

	//if err := json.Unmarshal(rawJSON, &cfg); err != nil {
	//	log.Panicln(err)
	//}

	var err error
	jsonbytes := []byte(JSONCONFIG)
	if err := json.Unmarshal(jsonbytes, &cfg); err != nil {
		log.Panicln(err)
	}
	logger, err = cfg.Build()
	if err != nil {
		log.Panicln(err)
	}

	sugar = logger.Sugar()
	sugar.Info("logger construction succeeded")
	return sugar
}

func Debug(args ...interface{}) {
	sugar.Debug(args...)
}

func Debugf(template string, args ...interface{}) {
	sugar.Debugf(template, args...)
}

func Debugw(msg string, keysAndValues ...interface{}) {
	sugar.Debugw(msg, keysAndValues...)
}

func Error(args ...interface{}) {
	sugar.Error(args...)
}

func Errorf(template string, args ...interface{}) {
	sugar.Errorf(template, args...)
}

func Errorw(msg string, keysAndValues ...interface{}) {
	sugar.Errorw(msg, keysAndValues...)
}

func Info(args ...interface{}) {
	sugar.Info(args...)
}

func Infof(template string, args ...interface{}) {
	sugar.Infof(template, args...)
}

func Infow(msg string, keysAndValues ...interface{}) {
	sugar.Infow(msg, keysAndValues...)
}

func Panic(args ...interface{}) {
	sugar.Panic(args...)
}

func Panicf(template string, args ...interface{}) {
	sugar.Panicf(template, args...)
}

func Panicw(msg string, keysAndValues ...interface{}) {
	sugar.Panicw(msg, keysAndValues...)
}

func Warn(args ...interface{}) {
	sugar.Warn(args...)
}

func Warnf(template string, args ...interface{}) {
	sugar.Warnf(template, args...)
}

func Warnw(msg string, keysAndValues ...interface{}) {
	sugar.Warnw(msg, keysAndValues...)
}

func With(args ...interface{}) {
	sugar.With(args...)
}
