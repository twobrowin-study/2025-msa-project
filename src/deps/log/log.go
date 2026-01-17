package log

import (
	"github.com/sirupsen/logrus"
)

// Подготовить логгер с кастомными настройками
func New() *logrus.Logger {
	var log = logrus.New()

	log.SetLevel(logrus.DebugLevel)

	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	return log
}
