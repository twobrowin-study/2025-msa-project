package deps

import (
	"otus.ru/tbw/msa-25/src/deps/config"
	"otus.ru/tbw/msa-25/src/deps/db"
	"otus.ru/tbw/msa-25/src/deps/log"
	"otus.ru/tbw/msa-25/src/deps/prometheus"
)

// Структура со всеми зависимостями приложения
type Deps struct {
	Log        *log.Logger
	Config     *config.Config
	DB         *db.DB
	Prometheus *prometheus.Prometheus
}

// Подготовка всех зависимостей для использования в приложении
func Prepare() *Deps {
	log := log.New()
	config := config.New(log)
	db := db.New(log, config)
	prometheus := prometheus.New()

	return &Deps{Log: log, Config: config, DB: db, Prometheus: prometheus}
}
