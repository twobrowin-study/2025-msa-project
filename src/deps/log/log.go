package log

import (
	"net/http"
	"time"

	"github.com/sirupsen/logrus"

	"regexp"
)

type Logger struct {
	*logrus.Logger
}

// Подготовить логгер с кастомными настройками
func New() *Logger {
	log := &Logger{
		Logger: logrus.New(),
	}

	log.SetLevel(logrus.DebugLevel)

	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	return log
}

// Middleware логгирования запросов пользователя и время их выполнения
// Исключает логгирование запросов /health и /metrics
func (log *Logger) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)

		notToLog, err := regexp.MatchString(`^/(healthz?|metrics)/?$`, r.URL.Path)
		if err != nil {
			log.Infof("Something went wrong matching url path with regexp to log: %v", err)
		}

		if notToLog {
			return
		}

		log.Infof("%s %s - %s", r.Method, r.URL.Path, time.Since(start))
	})
}
