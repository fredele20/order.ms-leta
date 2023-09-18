package postgresql

import (
	"database/sql"
	"log"

	database "github.com/fredele20/order.ms-leta/database/sql/sqlc-gen"
)


func ConnectDB(driverName, dataSourceName string) (*database.Queries, error) {
	connection, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Fatal("cannot connect to the database: ", err)
		return nil, err
	}

	log.Println("connected to postgresql db successfully...")

	db := database.New(connection)

	return db, nil
}