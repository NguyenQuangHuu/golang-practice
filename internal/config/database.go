package config

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

type DatabaseConnection interface {
	PSQLConnection() *sql.DB
}
type PSQLConnection struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
	SslMode  string
}

func (psql *PSQLConnection) PSQLConnection() *sql.DB {
	config := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", psql.Host, psql.Port, psql.User, psql.Password, psql.Name, psql.SslMode)
	db, err := sql.Open("postgres", config)

	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to Database")
	return db
}

func (psql *PSQLConnection) Close() {
	fmt.Println("Disconnected from database")
	err := psql.PSQLConnection().Close()
	if err != nil {
		log.Fatal(err)
	}
}
