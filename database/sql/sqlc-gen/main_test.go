package database

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	driverName     = "postgres"
	dataSourceName = "postgres://fredel:lyfgoes1@localhost:5432/leta-order.ms?sslmode=disable"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	conn, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	testQueries = New(conn)

	os.Exit(m.Run())
}
