package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"sync"

	"gedebook.com/api/env"
	_ "github.com/lib/pq"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

var dbInstance *bun.DB
var once sync.Once

func InitDB(c *env.DatabaseConfig) *bun.DB {
	once.Do(func() {
		// dsn := "postgres://postgres:postgres@localhost:5432/db_gedebook?sslmode=disable"
		dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
			c.User, c.Password, c.Host, c.Port, c.Name, c.SSLMode)
		sqlDb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
		dbInstance = bun.NewDB(sqlDb, pgdialect.New())
		// InitLogger(dbInstance, c.Debug, c.DebugLevel)
	})

	return dbInstance
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
