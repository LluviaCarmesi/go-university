package get;

import "back-end/models";
//import "github.com/go-sql-driver/mysql";
import "back-end/services";
import "back-end/settings";

func GetCourses() []models.Course {
	dbConnection := services.ConnectToDB();
	defer dbConnection.Close();
	courses := []models.Course{};
	
	results, err := dbConnection.Query(settings.GET_COURSES_QUERY);
	defer results.Close();

	if (err != nil) {
		panic("Error getting data: " + err.Error());
	}

	for results.Next() {
		var course models.Course;

		err := results.Scan(
			&course.ID,
			&course.Name,
			&course.Description,
			&course.Credits);
		if (err != nil) {
			panic("Error scanning row: " + err.Error());
		}

		courses = append(courses, course);
	}

	return courses;
}

func GetCourseByID(courseID string) models.Course {
	dbConnection := services.ConnectToDB();
	defer dbConnection.Close();
	course := models.Course{};

	query := settings.GET_COURSES_QUERY + " WHERE id = '" + courseID + "'";
	results, err := dbConnection.Query(query);
	defer results.Close();

	if (err != nil) {
		panic("Error getting data: " + err.Error());
	}

	for results.Next() {
		err := results.Scan(
			&course.ID,
			&course.Name,
			&course.Description,
			&course.Credits);
		if (err != nil) {
			panic("Error scanning row: " + err.Error());
		}
	}
	return course;
}

func GetDepartments() []models.Department {
	dbConnection := services.ConnectToDB();
	defer dbConnection.Close();
	departments := []models.Department{};
	
	results, err := dbConnection.Query(settings.GET_DEPARTMENTS_QUERY);
	defer results.Close();

	if (err != nil) {
		panic("Error getting data: " + err.Error());
	}

	for results.Next() {
		var department models.Department;

		err := results.Scan(
			&department.ID,
			&department.Name);
		if (err != nil) {
			panic("Error scanning row: " + err.Error());
		}

		departments = append(departments, department);
	}

	return departments;
}

func GetDepartmentByID(departmentID string) models.Department {
	dbConnection := services.ConnectToDB();
	defer dbConnection.Close();
	department := models.Department{};

	query := settings.GET_DEPARTMENTS_QUERY + " WHERE id = '" + departmentID + "'";
	results, err := dbConnection.Query(query);
	defer results.Close();

	if (err != nil) {
		panic("Error getting data: " + err.Error());
	}

	for results.Next() {
		err := results.Scan(
			&department.ID,
			&department.Name);
		if (err != nil) {
			panic("Error scanning row: " + err.Error());
		}
	}
	return department;
}

func GetProfessorsInDepartments() []models.ProfessorInDepartment {
	dbConnection := services.ConnectToDB();
	defer dbConnection.Close();
	professorsInDepartments := []models.ProfessorInDepartment{};
	departments := GetDepartments();
	results, err := dbConnection.Query(settings.GET_PROFESSORS_IN_DEPARTMENTS_QUERY);
	defer results.Close();

	if (err != nil) {
		panic("Error getting data: " + err.Error());
	}

	for results.Next() {
		var professorInDepartment models.ProfessorInDepartment;

		err := results.Scan(
			&professorInDepartment.ProfessorEmail,
			&professorInDepartment.DepartmentID,
			&professorInDepartment.IsLeader);
		if (err != nil) {
			panic("Error scanning row: " + err.Error());
		}
		for i := 0; i < len(departments); i++ {
			if (departments[i].ID == professorInDepartment.DepartmentID) {
				professorInDepartment.DepartmentName = departments[i].Name
			}
			break;
		}
		professorsInDepartments = append(professorsInDepartments, professorInDepartment);
	}

	return professorsInDepartments;
}

func GetProfessorsInDepartmentsByNameAndID(email string, departmentID string) models.ProfessorInDepartment {
	dbConnection := services.ConnectToDB();
	defer dbConnection.Close();
	professorInDepartment := models.ProfessorInDepartment{};

	query := settings.GET_PROFESSORS_IN_DEPARTMENTS_QUERY + " WHERE professor_email = '" + email + "' AND department_id = '" + departmentID + "'";
	results, err := dbConnection.Query(query);
	defer results.Close();

	if (err != nil) {
		panic("Error getting data: " + err.Error());
	}

	for results.Next() {
		err := results.Scan(
			&professorInDepartment.ProfessorEmail,
			&professorInDepartment.DepartmentID,
			&professorInDepartment.IsLeader);
		if (err != nil) {
			panic("Error scanning row: " + err.Error());
		}
	}
	return professorInDepartment;
}

func GetUsers() []models.User {
	dbConnection := services.ConnectToDB();
	defer dbConnection.Close();
	users := []models.User{};
	
	results, err := dbConnection.Query(settings.GET_USERS_QUERY);
	defer results.Close();

	if (err != nil) {
		panic("Error getting data: " + err.Error());
	}

	for results.Next() {
		var user models.User;

		err := results.Scan(
			&user.Email,
			&user.EmailAlias,
			&user.FirstName,
			&user.LastName,
			&user.PhoneNumber,
			&user.HomeAddress,
			&user.Role);
		if (err != nil) {
			panic("Error scanning row: " + err.Error());
		}

		users = append(users, user);
	}

	return users;
}

func GetUserByToken(token string) models.User{
	dbConnection := services.ConnectToDB();
	defer dbConnection.Close();
	user := models.User{};
	
	query := settings.GET_USERS_QUERY + " WHERE token = '" + token + "'";
	results, err := dbConnection.Query(query);
	defer results.Close();

	if (err != nil) {
		panic("Error getting data: " + err.Error());
	}

	for results.Next() {

		err := results.Scan(
			&user.Email,
			&user.EmailAlias,
			&user.FirstName,
			&user.LastName,
			&user.PhoneNumber,
			&user.HomeAddress,
			&user.Role);
		if (err != nil) {
			panic("Error scanning row: " + err.Error());
		}
	}
	return user;
}

func GetUserByEmail(email string) models.User{
	dbConnection := services.ConnectToDB();
	defer dbConnection.Close();
	user := models.User{};
	
	query := settings.GET_USERS_QUERY + " WHERE email_alias = '" + email + "' OR email = '" + email + "'";
	results, err := dbConnection.Query(query);
	defer results.Close();

	if (err != nil) {
		panic("Error getting data: " + err.Error());
	}

	for results.Next() {

		err := results.Scan(
			&user.Email,
			&user.EmailAlias,
			&user.FirstName,
			&user.LastName,
			&user.PhoneNumber,
			&user.HomeAddress,
			&user.Role);
		if (err != nil) {
			panic("Error scanning row: " + err.Error());
		}
	}
	return user;
}

func GetAppointments () []models.Appointment {
	dbConnection := services.ConnectToDB();
	defer dbConnection.Close();
	appointments := []models.Appointment{};
	
	results, err := dbConnection.Query(settings.GET_APPOINTMENTS_QUERY);
	defer results.Close();

	if (err != nil) {
		panic("Error getting data: " + err.Error());
	}

	for results.Next() {
		var appointment models.Appointment;

		err := results.Scan(
			&appointment.ID,
			&appointment.StudentEmail,
			&appointment.AdminEmail,
			&appointment.IsComplete,
			&appointment.StartTime,
			&appointment.EndTime);
		if (err != nil) {
			panic("Error scanning row: " + err.Error());
		}

		appointments = append(appointments, appointment);
	}

	return appointments;
}

func GetAppointmentByID(appointmentID string) models.Appointment {
	dbConnection := services.ConnectToDB();
	defer dbConnection.Close();
	appointment := models.Appointment{};

	query := settings.GET_APPOINTMENTS_QUERY + " WHERE id = '" + appointmentID + "'";
	results, err := dbConnection.Query(query);
	defer results.Close();

	if (err != nil) {
		panic("Error getting data: " + err.Error());
	}

	for results.Next() {
		err := results.Scan(
			&appointment.ID,
			&appointment.StudentEmail,
			&appointment.AdminEmail,
			&appointment.IsComplete,
			&appointment.StartTime,
			&appointment.EndTime);
		if (err != nil) {
			panic("Error scanning row: " + err.Error());
		}
	}
	return appointment;
}

func GetSemesters () []models.Semester {
	dbConnection := services.ConnectToDB();
	defer dbConnection.Close();
	semesters := []models.Semester{};
	
	results, err := dbConnection.Query(settings.GET_SEMESTERS_QUERY);
	defer results.Close();

	if (err != nil) {
		panic("Error getting data: " + err.Error());
	}

	for results.Next() {
		var semester models.Semester;

		err := results.Scan(
			&semester.Name,
			&semester.StartDate,
			&semester.EndDate);
		if (err != nil) {
			panic("Error scanning row: " + err.Error());
		}

		semesters = append(semesters, semester);
	}

	return semesters;
}

func GetSemesterByName(name string) models.Semester {
	dbConnection := services.ConnectToDB();
	defer dbConnection.Close();
	semester := models.Semester{};

	query := settings.GET_SEMESTERS_QUERY + " WHERE name = '" + name + "'";
	results, err := dbConnection.Query(query);
	defer results.Close();

	if (err != nil) {
		panic("Error getting data: " + err.Error());
	}

	for results.Next() {
		err := results.Scan(
			&semester.Name,
			&semester.StartDate,
			&semester.EndDate);
		if (err != nil) {
			panic("Error scanning row: " + err.Error());
		}
	}
	return semester;
}

func GetTaughtCourses() []models.TaughtCourse {
	dbConnection := services.ConnectToDB();
	defer dbConnection.Close();
	taughtCourses := []models.TaughtCourse{};

	results, err := dbConnection.Query(settings.GET_TAUGHT_COURSES_QUERY);
	defer results.Close();

	if (err != nil) {
		panic("Error getting data: " + err.Error());
	}

	for results.Next() {
		var taughtCourse models.TaughtCourse;

		err := results.Scan(
			&taughtCourse.CourseID,
			&taughtCourse.SemesterName,
			&taughtCourse.ProfessorEmail,
			&taughtCourse.MaxStudents,
			&taughtCourse.Location,
			&taughtCourse.StartTime,
			&taughtCourse.EndTime);
		if (err != nil) {
			panic("Error scanning row: " + err.Error());
		}
		
		taughtCourses = append(taughtCourses, taughtCourse);
	}

	return taughtCourses;
}

func GetRegistrations() []models.Registration {
	dbConnection := services.ConnectToDB();
	defer dbConnection.Close();
	registrations := []models.Registration{};

	results, err := dbConnection.Query(settings.GET_REGISTRATIONS_QUERY);
	defer results.Close();

	if (err != nil) {
		panic("Error getting data: " + err.Error());
	}

	for results.Next() {
		var registration models.Registration;

		err := results.Scan(
			&registration.ID,
			&registration.StudentEmail,
			&registration.TaughtCourseID,
			&registration.FinalGrade,
			&registration.Status);
		if (err != nil) {
			panic("Error scanning row: " + err.Error());
		}

		registrations = append(registrations, registration);
	}

	return registrations;
}