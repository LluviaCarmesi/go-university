package post;

import "back-end/models";
//import "github.com/go-sql-driver/mysql";
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
	
	query := settings.INSERT_DEPARTMENT_QUERY + "VALUES (?)";
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
	
	query := settings.INSERT_PROFESSOR_IN_DEPARTMENT_QUERY + "VALUES (?, ?, ?)";
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
	
	query := settings.INSERT_USER_QUERY + "VALUES (?, ?, ?, ?, ?, ?, ?, ?)";
	results, err := dbConnection.ExecContext(
		context.Background(),
		query,
		user.Email,
		user.EmailAlias,
		user.Password,
		user.FirstName,
		user.LastName,
		user.PhoneNumber,
		user.HomeAddress,
		user.Role);
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
	
	query := settings.INSERT_APPOINTMENT_QUERY + "VALUES (?, ?, ?, ?)";
	results, err := dbConnection.ExecContext(
		context.Background(),
		query,
		appointment.StudentEmail,
		appointment.AdminEmail,
		appointment.StartTime,
		appointment.EndTime);
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
	
	query := settings.INSERT_SEMESTER_QUERY + "VALUES (?, ?, ?)";
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
	
	query := settings.INSERT_TAUGHT_COURSE_QUERY + "VALUES (?, ?, ?, ?, ?)";
	results, err := dbConnection.ExecContext(
		context.Background(),
		query,
		taughtCourse.CourseID,
		taughtCourse.SemesterID,
		taughtCourse.ProfessorEmail,
		taughtCourse.MaxStudents,
		taughtCourse.Location);
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
	
	query := settings.INSERT_REGISTRATION_QUERY + "VALUES (?, ?, ?, ?)";
	results, err := dbConnection.ExecContext(
		context.Background(),
		query,
		registration.StudentEmail,
		registration.TaughtCourseID,
		registration.FinalGrade,
		registration.Status);
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

func AddCourseSchedule(courseSchedule models.CourseSchedule) models.ServiceResponse {
	serviceResponse := models.ServiceResponse{
		IsSuccessful: true,
		ErrorMessage: "",
	}
	dbConnection := services.ConnectToDB();
	defer dbConnection.Close();
	
	query := settings.INSERT_COURSE_SCHEDULE_QUERY + "VALUES (?, ?, ?)";
	results, err := dbConnection.ExecContext(
		context.Background(),
		query,
		courseSchedule.TaughtCourseID,
		courseSchedule.StartTime,
		courseSchedule.EndTime);
	if err != nil {
		serviceResponse.IsSuccessful = false;
		serviceResponse.ErrorMessage = "Unable to insert course_schedules: " + err.Error();
		return serviceResponse;
	}

	insertedID, err := results.LastInsertId();
	if err != nil {
		log.Fatalf("impossible to retrieve last inserted id for course_schedules: %s", err.Error())
	}
	log.Printf("inserted id for course_schedules: %d", insertedID);

	return serviceResponse;
}