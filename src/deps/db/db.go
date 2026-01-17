package db

import (
	"database/sql"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"

	"github.com/sirupsen/logrus"

	"otus.ru/tbw/msa-25/src/deps/config"
)

// Создание подключения к БД
func New(log *logrus.Logger, config *config.Config) *bun.DB {
	log.Info("Creating Postgresql DB connection...")

	pgconn := pgdriver.NewConnector(
		pgdriver.WithNetwork("tcp"),
		pgdriver.WithAddr(config.DB.Host+":"+config.DB.Port),
		pgdriver.WithInsecure(config.DB.SSLInsecure),
		pgdriver.WithUser(config.DB.Username),
		pgdriver.WithPassword(string(config.DB.Password)),
		pgdriver.WithDatabase(config.DB.Database),
		pgdriver.WithApplicationName(config.DB.ApplicationName),
		pgdriver.WithDialTimeout(config.DB.Timeout.Dial),
		pgdriver.WithReadTimeout(config.DB.Timeout.Read),
		pgdriver.WithWriteTimeout(config.DB.Timeout.Write),
	)

	return bun.NewDB(sql.OpenDB(pgconn), pgdialect.New(), bun.WithDiscardUnknownColumns())
}
