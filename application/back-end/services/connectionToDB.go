package services;

import "database/sql";
import _ "github.com/go-sql-driver/mysql";
import "fmt";

func ConnectToDB() *sql.DB {
	connection, err := sql.Open("mysql", "root:tooR@tcp(127.0.0.1:3306)/go_university");
	if (err != nil) {
		fmt.Println("Couldn't connect");
		panic(err);
	}

	connection.SetConnMaxLifetime(60);
	connection.SetMaxOpenConns(10);
	connection.SetMaxIdleConns(10);

	return connection;
}