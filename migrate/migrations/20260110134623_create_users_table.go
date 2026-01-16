package migrations

import (
	"context"
	"fmt"

	"github.com/uptrace/bun"

	"otus.ru/tbw/msa-25/src/api/user"
)

func init() {
	Migrations.MustRegister(func(ctx context.Context, db *bun.DB) error {
		fmt.Print(" [up migration] ")
		_, err := db.NewCreateTable().Model((*user.User)(nil)).Exec(ctx)
		return err
	}, func(ctx context.Context, db *bun.DB) error {
		fmt.Print(" [down migration] ")
		_, err := db.NewDropTable().Model((*user.User)(nil)).Exec(ctx)
		return err
	})
}
