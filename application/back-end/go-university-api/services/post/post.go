package post;

import "back-end/models";
//import "github.com/go-sql-driver/mysql";
import "back-end/services";
import "back-end/settings";
import "back-end/services/get";
import "log";
import "context";
import "back-end/utilities";

func AddCourse(course models.Course) models.ServiceResponse {
	serviceResponse := models.ServiceResponse{
		IsSuccessful: true,
		ErrorMessage: "",
	}
	dbConnection := services.ConnectToDB();
	defer dbConnection.Close();
	
	query := settings.INSERT_COURSE_QUERY;
	results, err := dbConnection.ExecContext(
		context.Background(),
		query,
		course.ID,
		course.Name,
		course.Description,
		course.Credits);
	if err != nil {
		serviceResponse.IsSuccessful = false;
		serviceResponse.ErrorMessage = "Unable to insert course: " + err.Error();
		return serviceResponse;
	}

	insertedID, err := results.LastInsertId();
	if err != nil {
		log.Fatalf("impossible to retrieve last inserted id for courses: %s", err.Error())
	}
	log.Printf("inserted id for courses: %d", insertedID);

	return serviceResponse;
}

func AddDepartment(department models.Department) models.ServiceResponse {
	serviceResponse := models.ServiceResponse{
		IsSuccessful: true,
		ErrorMessage: "",
	}
	dbConnection := services.ConnectToDB();
	defer dbConnection.Close();
	
	query := settings.INSERT_DEPARTMENT_QUERY;
	results, err := dbConnection.ExecContext(
		context.Background(),
		query,
		department.Name);
	if err != nil {
		serviceResponse.IsSuccessful = false;
		serviceResponse.ErrorMessage = "Unable to insert department: " + err.Error();
		return serviceResponse;
	}

	insertedID, err := results.LastInsertId();
	if err != nil {
		log.Fatalf("impossible to retrieve last inserted id for departments: %s", err.Error())
	}
	log.Printf("inserted id for departments: %d", insertedID);

	return serviceResponse;
}

func AddProfessorInDepartment(professorInDepartment models.ProfessorInDepartment) models.ServiceResponse {
	serviceResponse := models.ServiceResponse{
		IsSuccessful: true,
		ErrorMessage: "",
	}
	dbConnection := services.ConnectToDB();
	defer dbConnection.Close();
	
	query := settings.INSERT_PROFESSOR_IN_DEPARTMENT_QUERY;
	results, err := dbConnection.ExecContext(
		context.Background(),
		query,
		professorInDepartment.ProfessorEmail,
		professorInDepartment.DepartmentID,
		professorInDepartment.IsLeader);
	if err != nil {
		serviceResponse.IsSuccessful = false;
		serviceResponse.ErrorMessage = "Unable to insert professor in department: " + err.Error();
		return serviceResponse;
	}

	insertedID, err := results.LastInsertId();
	if err != nil {
		log.Fatalf("impossible to retrieve last inserted id for professors_in_departments: %s", err.Error())
	}
	log.Printf("inserted id for professors_in_departments: %d", insertedID);

	return serviceResponse;
}

func AddUser(user models.User) models.ServiceResponse {
	serviceResponse := models.ServiceResponse{
		IsSuccessful: true,
		ErrorMessage: "",
	}
	dbConnection := services.ConnectToDB();
	defer dbConnection.Close();
	encryptedPassword, passwordErr := utilities.EncryptPassword(user.Password);
	if (passwordErr != nil) {
		panic("Error encrypting password: " + passwordErr.Error());
	}
	
	query := settings.INSERT_USER_QUERY;
	results, err := dbConnection.ExecContext(
		context.Background(),
		query,
		user.Email,
		user.EmailAlias,
		encryptedPassword,
		user.FirstName,
		user.LastName,
		user.PhoneNumber,
		user.HomeAddress,
		user.Role,
		user.MustChangePW);
	if err != nil {
		serviceResponse.IsSuccessful = false;
		serviceResponse.ErrorMessage = "Unable to insert user: " + err.Error();
		return serviceResponse;
	}

	insertedID, err := results.LastInsertId();
	if err != nil {
		log.Fatalf("impossible to retrieve last inserted id for users: %s", err.Error())
	}
	log.Printf("inserted id for users: %d", insertedID);

	return serviceResponse;
}

func AddAppointment(appointment models.Appointment) models.ServiceResponse {
	serviceResponse := models.ServiceResponse{
		IsSuccessful: true,
		ErrorMessage: "",
	}
	dbConnection := services.ConnectToDB();
	defer dbConnection.Close();
	appointments := get.GetAppointments();

	for i := 0; i < len(appointments); i++ {
		if (((appointment.StartTime.After(appointments[i].StartTime) || appointment.StartTime.Equal(appointments[i].StartTime)) && (appointment.StartTime.Before(appointments[i].EndTime) || appointment.StartTime.Equal(appointments[i].EndTime))) || (appointment.EndTime.After(appointments[i].StartTime) || appointment.EndTime.Equal(appointments[i].StartTime)) && (appointment.EndTime.Before(appointments[i].EndTime) || appointment.EndTime.Equal(appointments[i].EndTime))) {
			serviceResponse.IsSuccessful = false;
			serviceResponse.ErrorMessage = "Appointment time conflicts with another";
			return serviceResponse;
		}
	}

	query := settings.INSERT_APPOINTMENT_QUERY;
	results, err := dbConnection.ExecContext(
		context.Background(),
		query,
		appointment.StudentEmail,
		appointment.AdminEmail,
		appointment.IsComplete,
		appointment.StartTime,
		appointment.EndTime,
		appointment.Description);
	if err != nil {
		serviceResponse.IsSuccessful = false;
		serviceResponse.ErrorMessage = "Unable to insert appointment: " + err.Error();
		return serviceResponse;
	}

	insertedID, err := results.LastInsertId();
	if err != nil {
		log.Fatalf("impossible to retrieve last inserted id for appointments: %s", err.Error())
	}
	log.Printf("inserted id for appointments: %d", insertedID);

	return serviceResponse;
}

func AddSemester(semester models.Semester) models.ServiceResponse {
	serviceResponse := models.ServiceResponse{
		IsSuccessful: true,
		ErrorMessage: "",
	}
	dbConnection := services.ConnectToDB();
	defer dbConnection.Close();
	
	query := settings.INSERT_SEMESTER_QUERY;
	results, err := dbConnection.ExecContext(
		context.Background(),
		query,
		semester.Name,
		semester.StartDate,
		semester.EndDate);
	if err != nil {
		serviceResponse.IsSuccessful = false;
		serviceResponse.ErrorMessage = "Unable to insert semester: " + err.Error();
		return serviceResponse;
	}

	insertedID, err := results.LastInsertId();
	if err != nil {
		log.Fatalf("impossible to retrieve last inserted id for semesters: %s", err.Error())
	}
	log.Printf("inserted id for semesters: %d", insertedID);

	return serviceResponse;
}

func AddTaughtCourse(taughtCourse models.TaughtCourse) models.ServiceResponse {
	serviceResponse := models.ServiceResponse{
		IsSuccessful: true,
		ErrorMessage: "",
	}
	dbConnection := services.ConnectToDB();
	defer dbConnection.Close();
	
	query := settings.INSERT_TAUGHT_COURSE_QUERY;
	results, err := dbConnection.ExecContext(
		context.Background(),
		query,
		taughtCourse.CourseID,
		taughtCourse.SemesterName,
		taughtCourse.ProfessorEmail,
		taughtCourse.MaxStudents,
		taughtCourse.Location,
		taughtCourse.Day,
		taughtCourse.StartTime,
		taughtCourse.EndTime);
	if err != nil {
		serviceResponse.IsSuccessful = false;
		serviceResponse.ErrorMessage = "Unable to insert taught_course: " + err.Error();
		return serviceResponse;
	}

	insertedID, err := results.LastInsertId();
	if err != nil {
		log.Fatalf("impossible to retrieve last inserted id for taught_courses: %s", err.Error())
	}
	log.Printf("inserted id for taught_courses: %d", insertedID);

	return serviceResponse;
}

func AddRegistration(registration models.Registration) models.ServiceResponse {
	serviceResponse := models.ServiceResponse{
		IsSuccessful: true,
		ErrorMessage: "",
	}
	dbConnection := services.ConnectToDB();
	defer dbConnection.Close();
	taughtCourses := get.GetTaughtCourses();
	registrations := get.GetRegistrations();
	currentSemester := "";
	courses := get.GetCourses();
	creditsThisSemester := 0;
	registrationStatus := "";
	taughtCourseID := 0;
	currentStudentCount := 1;
	maxStudentCount := 0;

	for i:= 0; i < len(taughtCourses); i++ {
		if (taughtCourses[i].ID == registration.TaughtCourseID) {
			currentSemester = taughtCourses[i].SemesterName;
			taughtCourseID = taughtCourses[i].ID;
			maxStudentCount = taughtCourses[i].MaxStudents;
			for j:= 0; j < len(courses); j++ {
				if (courses[j].ID == taughtCourses[i].CourseID) {
					creditsThisSemester += courses[j].Credits;
				}
			}
		}
	}

	for i:= 0; i < len(registrations); i++ {
		if (registrations[i].StudentEmail == registration.StudentEmail) {
			for j:= 0; j < len(taughtCourses); j++ {
				if (taughtCourses[j].ID == registrations[i].TaughtCourseID && currentSemester == taughtCourses[j].SemesterName) {
					for h:= 0; h < len(courses); h++ {
						if (courses[h].ID == taughtCourses[j].CourseID) {
							creditsThisSemester += courses[h].Credits;
						}
					}
				}
			}
		}
	}

	if (creditsThisSemester >= 18) {
		serviceResponse.IsSuccessful = false;
		serviceResponse.ErrorMessage = "Can't add more than 18 credits in courses";
		return serviceResponse;
	}
	
	for i := 0; i < len (registrations); i++ {
		if (registrations[i].TaughtCourseID == taughtCourseID) {
			currentStudentCount++;
		}
	}

	if (currentStudentCount > maxStudentCount) {
		registrationStatus = "Waiting Approval";
	} else {
		registrationStatus = "Registered";
	}

	query := settings.INSERT_REGISTRATION_QUERY;
	results, err := dbConnection.ExecContext(
		context.Background(),
		query,
		registration.StudentEmail,
		registration.TaughtCourseID,
		registration.FinalGrade,
		registrationStatus);
	if err != nil {
		serviceResponse.IsSuccessful = false;
		serviceResponse.ErrorMessage = "Unable to insert registrations: " + err.Error();
		return serviceResponse;
	}

	insertedID, err := results.LastInsertId();
	if err != nil {
		log.Fatalf("impossible to retrieve last inserted id for registrations: %s", err.Error())
	}
	log.Printf("inserted id for registrations: %d", insertedID);

	return serviceResponse;
}