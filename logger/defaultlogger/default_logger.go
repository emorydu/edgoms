package defaultLogger

import (
	"os"

	"github.com/emorydu/edgoms/pkg/constants"
	"github.com/emorydu/edgoms/pkg/logger"
	"github.com/emorydu/edgoms/pkg/logger/config"
	"github.com/emorydu/edgoms/pkg/logger/logrous"
	"github.com/emorydu/edgoms/pkg/logger/models"
	"github.com/emorydu/edgoms/pkg/logger/zap"
)

var l logger.Logger

func initLogger() {
	logType := os.Getenv("LogConfig_LogType")

	switch logType {
	case "Zap", "":
		l = zap.NewZapLogger(
			&config.LogOptions{LogType: models.Zap, CallerEnabled: false},
			constants.Dev,
		)
		break
	case "Logrus":
		l = logrous.NewLogrusLogger(
			&config.LogOptions{LogType: models.Logrus, CallerEnabled: false},
			constants.Dev,
		)
		break
	default:
	}
}

func GetLogger() logger.Logger {
	if l == nil {
		initLogger()
	}

	return l
}
