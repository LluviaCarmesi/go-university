package delete;

import "back-end/models";
//import "github.com/go-sql-driver/mysql";
import "back-end/settings";
import "back-end/services";
import "log";
import "context";

func DeleteCourse(id string) models.ServiceResponse {
	serviceResponse := models.ServiceResponse{
		IsSuccessful: true,
		ErrorMessage: "",
	}
	dbConnection := services.ConnectToDB();
	defer dbConnection.Close();

	query := settings.DELETE_COURSE_QUERY;
	_, err := dbConnection.ExecContext(
		context.Background(),
		query,
		id);
	if err != nil {
		serviceResponse.IsSuccessful = false;
		serviceResponse.ErrorMessage = "Unable to delete course: " + err.Error();
		return serviceResponse;
	}
	log.Printf("deleted from courses for the following id: %d", id);

	return serviceResponse;
}

func DeleteDepartment(id int) models.ServiceResponse {
	serviceResponse := models.ServiceResponse{
		IsSuccessful: true,
		ErrorMessage: "",
	}
	dbConnection := services.ConnectToDB();
	defer dbConnection.Close();

	query := settings.DELETE_DEPARTMENT_QUERY;
	_, err := dbConnection.ExecContext(
		context.Background(),
		query,
		id);
	if err != nil {
		serviceResponse.IsSuccessful = false;
		serviceResponse.ErrorMessage = "Unable to update department: " + err.Error();
		return serviceResponse;
	}
	log.Printf("updated departments for the following id: %d", id);

	return serviceResponse;
}

func DeleteAppointment(id int) models.ServiceResponse {
	serviceResponse := models.ServiceResponse{
		IsSuccessful: true,
		ErrorMessage: "",
	}
	dbConnection := services.ConnectToDB();
	defer dbConnection.Close();

	query := settings.DELETE_APPOINTMENT_QUERY;
	_, err := dbConnection.ExecContext(
		context.Background(),
		query,
		id);
	if err != nil {
		serviceResponse.IsSuccessful = false;
		serviceResponse.ErrorMessage = "Unable to delete appointment: " + err.Error();
		return serviceResponse;
	}
	log.Printf("deleted from appointments for the following id: %d", id);

	return serviceResponse;
}

func DeleteUser(email string) models.ServiceResponse {
	serviceResponse := models.ServiceResponse{
		IsSuccessful: true,
		ErrorMessage: "",
	}
	dbConnection := services.ConnectToDB();
	defer dbConnection.Close();

	query := settings.DELETE_USER_QUERY;
	_, err := dbConnection.ExecContext(
		context.Background(),
		query,
		email);
	if err != nil {
		serviceResponse.IsSuccessful = false;
		serviceResponse.ErrorMessage = "Unable to delete user: " + err.Error();
		return serviceResponse;
	}
	log.Printf("deleted from users for the following email: %d", email);

	return serviceResponse;
}