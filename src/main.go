package main

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"//mysql driver
	//it allows handling of broken connections, automatic connection pooling etc.
)
type User struct {
	Name string `json:"name"`
}

func main(){
	//We'll connect to a locally running MySQL instance.
	fmt.Println("Go MySQL")

	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/testdb")
	//it return a database object and an error object.
	//first arg is the driver name, second arg is the connection string(data source name)
	//which will be the concatination of the username and password, IP address, and the database name.
	//as well as the port number.
	if err !=nil{
		panic(err.Error())
	}
	defer db.Close()//it'll keep the database connection open until you call db.Close()

	insert, err:=db.Query("INSERT INTO users VALUES('shikharvashistha')")
	//it'll return a Query object and an error object.

	if err !=nil{
		panic(err.Error())
	}
	
	defer insert.Close()//it'll keep the database connection open until you call insert.Close()	
	fmt.Println("Sucessfully inserted into user tables")
	//now we return the result of the query in meaningfull way.
	results, err := db.Query("SELECT * FROM users")
	if err !=nil{
		panic(err.Error())
	}
	for results.Next(){
		var user User
		err=results.Scan(&user.Name)
		if err!=nil{
			panic(err.Error())
		}
		fmt.Println(user.Name)//print the name of the user inserted before
	}
}