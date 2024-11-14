package post;

import "back-end/models";
import "github.com/go-sql-driver/mysql";
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
		log.Fatalf("impossible to retrieve last inserted id: %s", err.Error())
	}
	log.Printf("inserted id: %d", insertedID);

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
		log.Fatalf("impossible to retrieve last inserted id: %s", err.Error())
	}
	log.Printf("inserted id: %d", insertedID);

	return serviceResponse;
}