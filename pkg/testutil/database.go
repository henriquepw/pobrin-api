package testutil

import (
	_ "github.com/glebarez/go-sqlite"
	"github.com/henriquepw/pobrin-api/internal/database"
	"github.com/jmoiron/sqlx"
)

func GetDB(migrations ...database.Migration) *sqlx.DB {
	db, err := sqlx.Open("sqlite", ":memory:")
	if err != nil {
		panic(err)
	}

	for _, m := range migrations {
		err := m(db)
		if err != nil {
			panic(err)
		}
	}

	return db
}
