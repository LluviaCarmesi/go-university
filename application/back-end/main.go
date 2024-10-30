package main;
import "fmt";
import "back-end/services";
import "back-end/services/get";
import _ "github.com/go-sql-driver/mysql";

type Course struct {
	ID string `json:id`
	Name string `json:"name"`
}

func main() {

	dbConnection := services.ConnectToDB();
	defer dbConnection.Close();

	getResults := get.Students(dbConnection);
	defer getResults.Close();
	
	for getResults.Next() {
		var course Course;

		err := getResults.Scan(&course.ID, &course.Name);

		if (err != nil) {
			panic(err.Error())
		}

		fmt.Println(course.ID);
	}

}