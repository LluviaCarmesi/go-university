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
			&course.Description);
		if (err != nil) {
			panic("Error scanning row: " + err.Error());
		}

		courses = append(courses, course);
	}

	return courses;
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

func GetProfessorsInDepartments() []models.ProfessorInDepartment {
	dbConnection := services.ConnectToDB();
	defer dbConnection.Close();
	professorsInDepartments := []models.ProfessorInDepartment{};
	
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

		professorsInDepartments = append(professorsInDepartments, professorInDepartment);
	}

	return professorsInDepartments;
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

func GetStudents() []models.User {
	dbConnection := services.ConnectToDB();
	defer dbConnection.Close();
	users := []models.User{};
	
	results, err := dbConnection.Query(settings.GET_STUDENTS_QUERY);
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
			&semester.ID,
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
			&taughtCourse.ID,
			&taughtCourse.CourseID,
			&taughtCourse.SemesterID,
			&taughtCourse.ProfessorEmail,
			&taughtCourse.MaxStudents,
			&taughtCourse.Location);
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

func GetCourseSchedules() []models.CourseSchedule {
	dbConnection := services.ConnectToDB();
	defer dbConnection.Close();
	courseSchedules := []models.CourseSchedule{};

	results, err := dbConnection.Query(settings.GET_COURSE_SCHEDULES_QUERY);
	defer results.Close();
	if (err != nil) {
		panic("Error getting data: " + err.Error());
	}

	for results.Next() {
		var courseSchedule models.CourseSchedule;

		err := results.Scan(
			&courseSchedule.ID,
			&courseSchedule.TaughtCourseID,
			&courseSchedule.StartTime,
			&courseSchedule.EndTime);
		if (err != nil) {
			panic("Error scanning row: " + err.Error());
		}

		courseSchedules = append(courseSchedules, courseSchedule);
	}

	return courseSchedules;
}