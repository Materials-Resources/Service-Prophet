package main

import (
	"database/sql"
	_ "github.com/microsoft/go-mssqldb"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mssqldialect"
	"github.com/uptrace/bun/extra/bundebug"
	"os"
)

func connectBun() *bun.DB {
	sqldb, err := sql.Open("sqlserver", os.Getenv("DSN"))
	if err != nil {
		panic(err)
	}
	db := bun.NewDB(sqldb, mssqldialect.New())
	db.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithVerbose(true),
		bundebug.FromEnv("BUNDEBUG"),
	))
	return db
}
