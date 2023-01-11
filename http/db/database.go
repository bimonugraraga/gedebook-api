package db

import (
	"context"
	"database/sql"
	"log"
	"sync"

	_ "github.com/lib/pq"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

var dbInstance *bun.DB
var once sync.Once

func InitDB() {
	once.Do(func() {
		dsn := "postgres://postgres:Scarlet13@localhost:5432/db_gedebook?sslmode=disable"
		sqlDb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
		dbInstance = bun.NewDB(sqlDb, pgdialect.New())
		// InitLogger(dbInstance, c.Debug, c.DebugLevel)
	})
}

func GetConn() *bun.DB {
	return dbInstance
}

func OpenConnection() int {
	return dbInstance.Stats().OpenConnections
}

func QueryContext(ctx context.Context, query string) (ids []int, err error) {
	var qr *sql.Rows
	db := GetConn()
	qr, err = db.QueryContext(ctx, query)
	if err != nil {
		log.Fatal("Error DB")
	}

	for qr.Next() {
		var id int
		err = qr.Scan(&id)
		if err == nil {
			ids = append(ids, id)
		}
	}
	return
}
