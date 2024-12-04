package delete;

import "back-end/models";
//import "github.com/go-sql-driver/mysql";
import "back-end/settings";
import "back-end/services";
import "back-end/services/get";
import "log";
import "context";
import "strconv";

func DeleteCourse(id string) models.ServiceResponse {
	serviceResponse := models.ServiceResponse{
		IsSuccessful: true,
		ErrorMessage: "",
	}
	dbConnection := services.ConnectToDB();
	defer dbConnection.Close();
	taughtCourses := get.GetTaughtCourses();
	registrations := get.GetRegistrations();

	for i := 0; i < len(taughtCourses); i++ {
		if (taughtCourses[i].CourseID == id) {
			for j := 0; j < len(registrations); j++ {
				if (registrations[j].TaughtCourseID == taughtCourses[i].ID) {
					DeleteRegistration(registrations[j].StudentEmail, strconv.Itoa(registrations[j].TaughtCourseID));
				}
			}
			DeleteTaughtCourse(strconv.Itoa(taughtCourses[i].ID));
		}
	}

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
	professorsInDepartments := get.GetProfessorsInDepartments();

	for i := 0; i < len(professorsInDepartments); i++ {
		if (strconv.Itoa(professorsInDepartments[i].DepartmentID) == id) {
			DeleteProfessorInDepartment(professorsInDepartments[i].ProfessorEmail, strconv.Itoa(professorsInDepartments[i].DepartmentID))
		}
	}

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
	professorsInDepartments := get.GetProfessorsInDepartments();
	registrations := get.GetRegistrations();
	appointments := get.GetAppointments();
	taughtCourses := get.GetTaughtCourses();

	for i:= 0; i < len(registrations); i++ {
		if (registrations[i].StudentEmail == email) {
			DeleteRegistration(registrations[i].StudentEmail, strconv.Itoa(registrations[i].TaughtCourseID));
		}
	}

	for i:= 0; i < len(taughtCourses); i++ {
		if (taughtCourses[i].ProfessorEmail == email) {
			DeleteTaughtCourse(strconv.Itoa(taughtCourses[i].ID));
		}
	}

	for i:= 0; i < len(appointments); i++ {
		if (appointments[i].AdminEmail == email || appointments[i].StudentEmail == email) {
			DeleteAppointment(strconv.Itoa(appointments[i].ID));
		}
	}

	for i:= 0; i < len(professorsInDepartments); i++ {
		if (professorsInDepartments[i].ProfessorEmail == email) {
			DeleteProfessorInDepartment(professorsInDepartments[i].ProfessorEmail, strconv.Itoa(professorsInDepartments[i].DepartmentID));
		}
	}

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
	taughtCourses := get.GetTaughtCourses();
	registrations := get.GetRegistrations();

	for i := 0; i < len(taughtCourses); i++ {
		if (taughtCourses[i].SemesterName == name) {
			for j := 0; j < len(registrations); j++ {
				if (registrations[j].TaughtCourseID == taughtCourses[i].ID) {
					DeleteRegistration(registrations[j].StudentEmail, strconv.Itoa(registrations[j].TaughtCourseID));
				}
			}
			DeleteTaughtCourse(strconv.Itoa(taughtCourses[i].ID));
		}
	}

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
	registrations := get.GetRegistrations();

	for j := 0; j < len(registrations); j++ {
		if (strconv.Itoa(registrations[j].TaughtCourseID) == id) {
			DeleteRegistration(registrations[j].StudentEmail, strconv.Itoa(registrations[j].TaughtCourseID));
		}
	}

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