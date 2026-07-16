package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func InitDB() (*sql.DB, error) {

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	name := os.Getenv("DB_NAME")

	if user == "" {
		user = "root"
	}

	if host == "" {
		host = "localhost"
	}

	if port == "" {
		port = "3306"
	}

	if name == "" {
		name = "industrial_supply_store"
	}

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		user,
		pass,
		host,
		port,
		name,
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("cant connect db:", err)
		return nil, err
	}

	if err = db.Ping(); err != nil {
		fmt.Println("cant ping db:", err)
		return nil, err
	}

	return db, nil
}
