package get;

import "fmt";
import _ "github.com/go-sql-driver/mysql";
import "database/sql";

func Students (connection *sql.DB) *sql.Rows {
	results, err := connection.Query("SELECT id,name FROM courses");

	if (err != nil) {
		fmt.Println("Couldn't get data");
		panic(err.Error());
	}

	return results;
}