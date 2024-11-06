package get;

import "fmt";
import "back-end/models";
import _ "github.com/go-sql-driver/mysql";
import "back-end/services";
import "back-end/settings";

func Courses () []models.Course {
	dbConnection := services.ConnectToDB();
	defer dbConnection.Close();
	courses := []models.Course{};
	
	results, err := dbConnection.Query("SELECT id,name FROM courses");
	defer results.Close();

	if (err != nil) {
		fmt.Println("Couldn't get data");
		panic(err.Error());
	}

	for results.Next() {
		var course models.Course;

		err := results.Scan(&course.ID, &course.Name);
		if (err != nil) {
			panic(err.Error())
		}

		courses = append(courses, course);
	}

	return courses;
}

func Users () []models.User {
	dbConnection := services.ConnectToDB();
	defer dbConnection.Close();
	users := []models.User{};
	
	results, err := dbConnection.Query(settings.GET_USERS_QUERY);
	defer results.Close();

	if (err != nil) {
		fmt.Println("Couldn't get data");
		panic(err.Error());
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
			panic(err.Error())
		}

		users = append(users, user);
	}

	return users;
}

func Students () []models.User {
	dbConnection := services.ConnectToDB();
	defer dbConnection.Close();
	users := []models.User{};
	
	results, err := dbConnection.Query(settings.GET_STUDENTS_QUERY);
	defer results.Close();

	if (err != nil) {
		fmt.Println("Couldn't get data");
		panic(err.Error());
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
			panic(err.Error())
		}

		users = append(users, user);
	}

	return users;
}

func Appointments () []models.Appointment {
	dbConnection := services.ConnectToDB();
	defer dbConnection.Close();
	appointments := []models.Appointment{};
	
	results, err := dbConnection.Query(settings.GET_APPOINTMENTS_QUERY);
	defer results.Close();

	if (err != nil) {
		fmt.Println("Couldn't get data");
		panic(err.Error());
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
			panic(err.Error())
		}

		appointments = append(appointments, appointment);
	}

	return appointments;
}

func Semesters () []models.Semester {
	dbConnection := services.ConnectToDB();
	defer dbConnection.Close();
	semesters := []models.Semester{};
	
	results, err := dbConnection.Query(settings.GET_SEMESTERS_QUERY);
	defer results.Close();

	if (err != nil) {
		fmt.Println("Couldn't get data");
		panic(err.Error());
	}

	for results.Next() {
		var semester models.Semester;

		err := results.Scan(
			&semester.ID,
			&semester.Name,
			&semester.StartDate,
			&semester.EndDate);
		if (err != nil) {
			panic(err.Error())
		}

		semesters = append(semesters, semester);
	}

	return semesters;
}