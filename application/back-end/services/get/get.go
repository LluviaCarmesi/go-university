package get;

import "fmt";
import "back-end/models";
import _ "github.com/go-sql-driver/mysql";
import "back-end/services";

func Courses () []models.Course {
	dbConnection := services.ConnectToDB();
	defer dbConnection.Close();
	courses := []models.Course{};
	
	results, err := dbConnection.Query("SELECT id,name FROM courses");
	defer results.Close();

	if (err != nil) {
		fmt.Println("Couldn't get data");
		panic(err.Error());
	}

	for results.Next() {
		var course models.Course;

		err := results.Scan(&course.ID, &course.Name);
		if (err != nil) {
			panic(err.Error())
		}

		courses = append(courses, course);
	}

	return courses;
}