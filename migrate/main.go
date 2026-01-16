package main

import (
	"strings"

	"github.com/uptrace/bun/migrate"

	"github.com/urfave/cli/v2"

	"os"

	"otus.ru/tbw/msa-25/migrate/migrations"
	"otus.ru/tbw/msa-25/src/deps"
)

func main() {
	deps := deps.Prepare()

	templateData := map[string]string{
		"Prefix": "otus-msa-25_",
	}
	app := &cli.App{
		Name: "bun",

		Commands: []*cli.Command{
			newDBCommand(migrate.NewMigrator(deps.DB, migrations.Migrations, migrate.WithTemplateData(templateData)), deps),
		},
	}
	if err := app.Run(os.Args); err != nil {
		deps.Log.Fatal(err)
	}
}

func newDBCommand(migrator *migrate.Migrator, deps *deps.Deps) *cli.Command {
	return &cli.Command{
		Name:  "db",
		Usage: "database migrations",
		Subcommands: []*cli.Command{
			{
				Name:  "init",
				Usage: "create migration tables",
				Action: func(c *cli.Context) error {
					if err := migrator.Init(c.Context); err != nil {
						return err
					}
					deps.Log.Info("Created migration tables")
					return nil
				},
			},
			{
				Name:  "migrate",
				Usage: "migrate database",
				Action: func(c *cli.Context) error {
					if err := migrator.Lock(c.Context); err != nil {
						return err
					}
					defer migrator.Unlock(c.Context) //nolint:errcheck

					deps.Log.Info("Performing migrations...")
					group, err := migrator.Migrate(c.Context)
					if err != nil {
						return err
					}
					if group.IsZero() {
						deps.Log.Info("There are no new migrations to run (database is up to date)")
						return nil
					}
					deps.Log.Infof("Migrated to %s\n", group)
					return nil
				},
			},
			{
				Name:  "rollback",
				Usage: "rollback the last migration group",
				Action: func(c *cli.Context) error {
					if err := migrator.Lock(c.Context); err != nil {
						return err
					}
					defer migrator.Unlock(c.Context) //nolint:errcheck

					deps.Log.Info("Performing rollback...")
					group, err := migrator.Rollback(c.Context)
					if err != nil {
						return err
					}
					if group.IsZero() {
						deps.Log.Info("There are no groups to roll back\n")
						return nil
					}
					deps.Log.Infof("Rolled back %s\n", group)
					return nil
				},
			},
			{
				Name:  "lock",
				Usage: "lock migrations",
				Action: func(c *cli.Context) error {
					if err := migrator.Lock(c.Context); err != nil {
						return err
					}
					deps.Log.Info("Locked migrations")
					return nil
				},
			},
			{
				Name:  "unlock",
				Usage: "unlock migrations",
				Action: func(c *cli.Context) error {
					if err := migrator.Unlock(c.Context); err != nil {
						return err
					}
					deps.Log.Info("Unlocked migrations")
					return nil
				},
			},
			{
				Name:  "create_go",
				Usage: "create Go migration",
				Action: func(c *cli.Context) error {
					name := strings.Join(c.Args().Slice(), "_")
					mf, err := migrator.CreateGoMigration(c.Context, name)
					if err != nil {
						return err
					}
					deps.Log.Infof("Created migration %s (%s)\n", mf.Name, mf.Path)
					return nil
				},
			},
			{
				Name:  "create_sql",
				Usage: "create up and down SQL migrations",
				Action: func(c *cli.Context) error {
					name := strings.Join(c.Args().Slice(), "_")
					files, err := migrator.CreateSQLMigrations(c.Context, name)
					if err != nil {
						return err
					}

					for _, mf := range files {
						deps.Log.Infof("Created migration %s (%s)\n", mf.Name, mf.Path)
					}

					return nil
				},
			},
			{
				Name:  "create_tx_sql",
				Usage: "create up and down transactional SQL migrations",
				Action: func(c *cli.Context) error {
					name := strings.Join(c.Args().Slice(), "_")
					files, err := migrator.CreateTxSQLMigrations(c.Context, name)
					if err != nil {
						return err
					}

					for _, mf := range files {
						deps.Log.Infof("Created transaction migration %s (%s)\n", mf.Name, mf.Path)
					}

					return nil
				},
			},
			{
				Name:  "status",
				Usage: "print migrations status",
				Action: func(c *cli.Context) error {
					ms, err := migrator.MigrationsWithStatus(c.Context)
					if err != nil {
						return err
					}
					deps.Log.Infof("Migrations: %s\n", ms)
					deps.Log.Infof("Unapplied migrations: %s\n", ms.Unapplied())
					deps.Log.Infof("Last migration group: %s\n", ms.LastGroup())
					return nil
				},
			},
			{
				Name:  "mark_applied",
				Usage: "mark migrations as applied without actually running them",
				Action: func(c *cli.Context) error {
					group, err := migrator.Migrate(c.Context, migrate.WithNopMigration())
					if err != nil {
						return err
					}
					if group.IsZero() {
						deps.Log.Infof("There are no new migrations to mark as applied\n")
						return nil
					}
					deps.Log.Infof("Marked as applied %s\n", group)
					return nil
				},
			},
		},
	}
}
