package base

import (
	"github.com/SAIKAII/skHappy-IM/infra"
	"github.com/sirupsen/logrus"
	"os"
)

var Logger *logrus.Logger

type LoggerStarter struct {
	infra.BaseStarter
}

func (l *LoggerStarter) Setup(ctx infra.StarterContext) {
	// 日志等级
	logrus.SetLevel(logrus.DebugLevel)
	// 日志输出地方
	logrus.SetOutput(os.Stdout)
	// JSON格式输出
	logrus.SetFormatter(&logrus.JSONFormatter{})
	Logger = logrus.StandardLogger()
}
