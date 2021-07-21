package main

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func main(){
	fmt.Println("Go MySQL")

	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/testdb")
	if err !=nil{
		panic(err.Error())
	}
	defer db.Close()

	insert, err:=db.Query("INSERT INTO users VALUES('shikharvashistha')")

	if err !=nil{
		panic(err.Error())
	}
	defer insert.Close()
	fmt.Println("Sucessfully inserted into user tables")
}