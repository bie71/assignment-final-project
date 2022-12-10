package migrations

import (
	_ "assigment-final-project/helper"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log"
	"net/url"
	"os"
	"testing"
)

func TestCreateMigrations(t *testing.T) {

	dbDriver := os.Getenv("DB_DRIVER")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USERNAME")
	dbHost := os.Getenv("DB_HOST")
	dbPass := os.Getenv("DB_PASSWORD")
	dbPort := os.Getenv("DB_PORT")

	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Jakarta")
	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())

	up(dbDriver, dsn)
	log.Println("UP")
	//down(dbDriver, dsn)
	//log.Println("DOWN")
	log.Println("Success")
}

func up(dbDriver, dsn string) {
	m, err := migrate.New(
		"file://",
		dbDriver+"://"+dsn)

	if err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil { //to up or create table
		log.Fatal(err)
	}
}

func down(dbDriver, dsn string) {
	m, err := migrate.New(
		"file://",
		dbDriver+"://"+dsn)

	if err != nil {
		log.Fatal(err)
	}
	if err := m.Down(); err != nil { //to up or create table
		log.Fatal(err)
	}
}
