package storezap

import (
	"github.com/gubtos/gosubmoduleexample"
	"go.uber.org/zap"
)

var zapLogger *zap.Logger

func init() {
	zapLogger, _ = zap.NewProduction()
}

func Log(data gosubmoduleexample.Data) {
	zapLogger.Info(
		"current data",
		zap.String("val A", data.ValA),
		zap.String("val B", data.ValB))
}
