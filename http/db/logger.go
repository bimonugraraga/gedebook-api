package db

import (
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/extra/bundebug"
)

func InitLogger(db *bun.DB, debug bool, level int) {
	db.AddQueryHook(
		bundebug.NewQueryHook(
			bundebug.WithEnabled(debug),
			bundebug.WithVerbose(level == 2),
		),
	)
}
