package log

import (
	"github.com/sirupsen/logrus"
)

func New() (*logrus.Logger) {
    var log = logrus.New()

	log.SetLevel(logrus.DebugLevel)

    log.SetFormatter(&logrus.TextFormatter{
        FullTimestamp: true,
    })

	return log
}