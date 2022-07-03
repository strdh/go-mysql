package main

import (
	"testing"
	
	_"github.com/go-sql-driver/mysql"
)

func TestOpenConnection(t * testing.T) {
	db, err := sql.Open("mysql", "root:@tcp(host:3306)/gom")
	if err != nil {
		panic(err)
	}

	defer db.Close()
}