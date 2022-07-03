package main

import (
	"testing"
	"database/sql"
	_"github.com/go-sql-driver/mysql"
	"context"
	"time"
	"fmt"
)

func GetConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(host:3306)/gom")
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}

func TestInsertSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "INSERT INTO people VALUES(nil,'Anselma Hanadya Putri')"
	_, err := db.ExecContext(ctx, query)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success Add")
}