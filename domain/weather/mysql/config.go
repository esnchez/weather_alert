package mysql

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type Database struct {
	db *sql.DB
}

func NewMySQLDatabase() (*Database, error) {

	mysqlURI := GetConnectionString()
	db, err := sql.Open("mysql", mysqlURI)
	if err != nil {
		return nil, err
	}
	//Check if the connection is established properly with the DB
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &Database{db: db}, nil
}

func (d *Database) Close() {
	d.db.Close()
}

func (d *Database) GetDB() *sql.DB {
	return d.db
}

func GetConnectionString() string {

	host := os.Getenv("DB_HOST")
	if host == "" {
		host = "localhost"
	}

	port := os.Getenv("DB_PORT")
	if port == "" {
		port = "3306"
	}

	user := os.Getenv("DB_USER")
	if user == "" {
		user = "user"
	}

	password := os.Getenv("DB_PASSWORD")
	if password == "" {
		password = "password"
	}

	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = "weather_db"
	}

	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, dbName)
}
