package logger

import (
	"github.com/denizzengin/common"
	"go.uber.org/zap"
)

var Log *zap.Logger = New()

func New() *zap.Logger {

	var cfg zap.Config
	if common.IsProduction() {
		cfg = zap.NewProductionConfig()
		cfg.Encoding ="json"
		cfg.OutputPaths= []string{"./beauty.log"} //TODO read from conf module	
	} else {
		cfg = zap.NewDevelopmentConfig()
		cfg.Encoding = "console"		
	}
	var loggerInstance *zap.Logger
	loggerInstance, _ = cfg.Build()
	return loggerInstance
}
