package patch;

import "back-end/models";
import "github.com/go-sql-driver/mysql";
import "back-end/settings";
import "back-end/services";
import "log";
import "context";

func EditCourse(course models.Course) models.ServiceResponse {
	serviceResponse := models.ServiceResponse{
		IsSuccessful: true,
		ErrorMessage: "",
	}
	dbConnection := services.ConnectToDB();
	defer dbConnection.Close();

	query := settings.UPDATE_COURSE_QUERY;
	_, err := dbConnection.ExecContext(
		context.Background(),
		query,
		course.Name,
		course.Description,
		course.ID);
	if err != nil {
		serviceResponse.IsSuccessful = false;
		serviceResponse.ErrorMessage = "Unable to update course: " + err.Error();
		return serviceResponse;
	}
	log.Printf("updated courses for the following id: %d", course.ID);

	return serviceResponse;
}