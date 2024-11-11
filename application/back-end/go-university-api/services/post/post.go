package post;

import "back-end/models";
//import _ "github.com/go-sql-driver/mysql";
//import "back-end/services";
//import "back-end/settings";
import "fmt";

func AddCourse(course models.Course) models.ServiceResponse {
	fmt.Println(course.ID);
	fmt.Println(course.Name);
	fmt.Println(course.Description);

	serviceResponse := models.ServiceResponse{
		IsSuccessful: true,
		ErrorMessage: "",
	}

	return serviceResponse;
}