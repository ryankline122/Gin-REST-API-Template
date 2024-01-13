package data

import (
    "fmt"
	"database/sql"
	_ "github.com/lib/pq"
)

const (
	dbHost     = "localhost"
	dbPort     = 5432
	dbUser     = "dev"
	dbPassword = "pgdev"
	dbName     = "albums"
)

func ConnectToDatabase() (*sql.DB, error) {
    dbInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
        dbHost, dbPort, dbUser, dbPassword, dbName)
    db, err := sql.Open("postgres", dbInfo)
    if err != nil {
        return nil, err
    }
    return db, nil
}
