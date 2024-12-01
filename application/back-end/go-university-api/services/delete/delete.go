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

func DeleteDepartment(id string) models.ServiceResponse {
	serviceResponse := models.ServiceResponse{
		IsSuccessful: true,
		ErrorMessage: "",
	}
	dbConnection := services.ConnectToDB();
	defer dbConnection.Close();

	query := settings.DELETE_DEPARTMENT_QUERY + "";
	_, err := dbConnection.ExecContext(
		context.Background(),
		query,
		id);
	if err != nil {
		serviceResponse.IsSuccessful = false;
		serviceResponse.ErrorMessage = "Unable to delete department: " + err.Error();
		return serviceResponse;
	}
	log.Printf("deleted from departments for the following id: %d", id);

	return serviceResponse;
}

func DeleteProfessorInDepartment(email string, departmentID string) models.ServiceResponse {
	serviceResponse := models.ServiceResponse{
		IsSuccessful: true,
		ErrorMessage: "",
	}
	dbConnection := services.ConnectToDB();
	defer dbConnection.Close();

	query := settings.DELETE_PROFESSOR_IN_DEPARTMENT_QUERY + "";
	_, err := dbConnection.ExecContext(
		context.Background(),
		query,
		email,
		departmentID);
	if err != nil {
		serviceResponse.IsSuccessful = false;
		serviceResponse.ErrorMessage = "Unable to delete professor in department: " + err.Error();
		return serviceResponse;
	}
	log.Printf(
		"deleted from professors_in_departments for the following email and department id: %d %d",
		email,
		departmentID);

	return serviceResponse;
}

func DeleteAppointment(id string) models.ServiceResponse {
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

func DeleteSemester(name string) models.ServiceResponse {
	serviceResponse := models.ServiceResponse{
		IsSuccessful: true,
		ErrorMessage: "",
	}
	dbConnection := services.ConnectToDB();
	defer dbConnection.Close();

	query := settings.DELETE_SEMESTER_QUERY;
	_, err := dbConnection.ExecContext(
		context.Background(),
		query,
		name);
	if err != nil {
		serviceResponse.IsSuccessful = false;
		serviceResponse.ErrorMessage = "Unable to delete semester: " + err.Error();
		return serviceResponse;
	}
	log.Printf("deleted from semesters for the following name: %d", name);

	return serviceResponse;
}

func DeleteTaughtCourse(id string) models.ServiceResponse {
	serviceResponse := models.ServiceResponse{
		IsSuccessful: true,
		ErrorMessage: "",
	}
	dbConnection := services.ConnectToDB();
	defer dbConnection.Close();

	query := settings.DELETE_TAUGHT_COURSE_QUERY;
	_, err := dbConnection.ExecContext(
		context.Background(),
		query,
		id);
	if err != nil {
		serviceResponse.IsSuccessful = false;
		serviceResponse.ErrorMessage = "Unable to delete taught course: " + err.Error();
		return serviceResponse;
	}
	log.Printf("deleted from taught_courses for the following id: %d", id);

	return serviceResponse;
}

func DeleteRegistration(email string, taughtCourseID string) models.ServiceResponse {
	serviceResponse := models.ServiceResponse{
		IsSuccessful: true,
		ErrorMessage: "",
	}
	dbConnection := services.ConnectToDB();
	defer dbConnection.Close();

	query := settings.DELETE_REGISTRATION_QUERY;
	_, err := dbConnection.ExecContext(
		context.Background(),
		query,
		email,
		taughtCourseID);
	if err != nil {
		serviceResponse.IsSuccessful = false;
		serviceResponse.ErrorMessage = "Unable to delete registration: " + err.Error();
		return serviceResponse;
	}
	log.Printf(
		"deleted from registrations for the following email and taught course id: %d %d",
		email,
		taughtCourseID);

	return serviceResponse;
}