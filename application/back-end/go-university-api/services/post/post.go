package post;

import "back-end/models";
import _ "github.com/go-sql-driver/mysql";
import "back-end/services";
import "back-end/settings";
import "log";
import "context";

func AddCourse(course models.Course) models.ServiceResponse {
	serviceResponse := models.ServiceResponse{
		IsSuccessful: true,
		ErrorMessage: "",
	}
	dbConnection := services.ConnectToDB();
	defer dbConnection.Close();
	
	query := settings.INSERT_COURSE_QUERY + "VALUES (?, ?, ?)";
	results, err := dbConnection.ExecContext(
		context.Background(),
		query,
		course.ID,
		course.Name,
		course.Description);
	if err != nil {
		serviceResponse.IsSuccessful = false;
		serviceResponse.ErrorMessage = "Unable to insert course: " + err.Error();
		return serviceResponse;
	}

	insertedID, err := results.LastInsertId();
	if err != nil {
		log.Fatalf("impossible to retrieve last inserted id: %s", err.Error())
	}
	log.Printf("inserted id: %d", insertedID);

	return serviceResponse;
}