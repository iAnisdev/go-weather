package main

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Database struct {
	DB *sql.DB
}

func ConnectDB() Database {
	USER := os.Getenv("DB_USER")
	PASS := os.Getenv("DB_PASSWORD")
	HOST := os.Getenv("DB_HOST")
	PORT := os.Getenv("DB_PORT")
	DBNAME := os.Getenv("DB_NAME")
	connectionStr := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s", USER, PASS, HOST, PORT, DBNAME)
	db, err := sql.Open("mysql", connectionStr)

	if err != nil {
		fmt.Println("Failed to connect to database!")
		fmt.Println(err)
		panic(err)
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	fmt.Println("Database connection established")
	return Database{
		DB: db,
	}
}
