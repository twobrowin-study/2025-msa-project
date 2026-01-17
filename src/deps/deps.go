package deps

import (
	"github.com/sirupsen/logrus"
	"github.com/uptrace/bun"

	"otus.ru/tbw/msa-25/src/deps/config"
	"otus.ru/tbw/msa-25/src/deps/db"
	"otus.ru/tbw/msa-25/src/deps/log"
)

type Deps struct {
	Log    *logrus.Logger
	Config *config.Config
	DB     *bun.DB
}

func Prepare() (app *Deps) {
	log := log.New()
	config := config.New(log)
	db := db.New(log, config)

	return &Deps{Log: log, Config: config, DB: db}
}
