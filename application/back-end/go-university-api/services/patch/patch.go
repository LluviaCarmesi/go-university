package patch;

import "back-end/models";
//import "github.com/go-sql-driver/mysql";
import "back-end/settings";
import "back-end/services";
import "back-end/services/get";
import "math/rand";
import "log";
import "context";
import "time";

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
		course.Credits,
		course.ID);
	if err != nil {
		serviceResponse.IsSuccessful = false;
		serviceResponse.ErrorMessage = "Unable to update course: " + err.Error();
		return serviceResponse;
	}
	log.Printf("updated courses for the following id: %d", course.ID);

	return serviceResponse;
}

func EditDepartment(department models.Department) models.ServiceResponse {
	serviceResponse := models.ServiceResponse{
		IsSuccessful: true,
		ErrorMessage: "",
	}
	dbConnection := services.ConnectToDB();
	defer dbConnection.Close();

	query := settings.UPDATE_DEPARTMENT_QUERY;
	_, err := dbConnection.ExecContext(
		context.Background(),
		query,
		department.Name);
	if err != nil {
		serviceResponse.IsSuccessful = false;
		serviceResponse.ErrorMessage = "Unable to update department: " + err.Error();
		return serviceResponse;
	}
	log.Printf("updated departments for the following id: %d", department.ID);

	return serviceResponse;
}

func EditAppointment(appointment models.Appointment) models.ServiceResponse {
	serviceResponse := models.ServiceResponse{
		IsSuccessful: true,
		ErrorMessage: "",
	}
	dbConnection := services.ConnectToDB();
	defer dbConnection.Close();

	query := settings.UPDATE_APPOINTMENT_QUERY;
	_, err := dbConnection.ExecContext(
		context.Background(),
		query,
		appointment.StudentEmail,
		appointment.AdminEmail,
		appointment.IsComplete,
		appointment.StartTime,
		appointment.EndTime,
		appointment.ID);
	if err != nil {
		serviceResponse.IsSuccessful = false;
		serviceResponse.ErrorMessage = "Unable to update appointment: " + err.Error();
		return serviceResponse;
	}
	log.Printf("updated appointments for the following id: %d", appointment.ID);

	return serviceResponse;
}

func LoginUser(user models.User) models.ServiceResponseLogin {
	rand.Seed(time.Now().UnixNano());
	serviceResponseLogin := models.ServiceResponseLogin{
		IsSuccessful: true,
		ErrorMessage: "",
		Token: "",
	}
	dbConnection := services.ConnectToDB();
	defer dbConnection.Close();
	encryptedPassword := user.Password;
	receivedUser := models.User{};
	
	query := settings.GET_USERS_QUERY + " WHERE email = '" + user.Email + "' AND password = '" + encryptedPassword + "'";
	results, err := dbConnection.Query(query);
	defer results.Close();

	if (err != nil) {
		panic("Error getting data: " + err.Error());
	}

	for results.Next() {

		err := results.Scan(
			&receivedUser.Email,
			&receivedUser.EmailAlias,
			&receivedUser.FirstName,
			&receivedUser.LastName,
			&receivedUser.PhoneNumber,
			&receivedUser.HomeAddress,
			&receivedUser.Role);
		if (err != nil) {
			panic("Error scanning row: " + err.Error());
		}
	}
	if (receivedUser.Email == "") {
		panic("No user found");
	}

	newToken := make([]byte, 255);
	for i := 0; i < 255; i++ {
		newToken[i] = settings.ALPHABET[rand.Intn(len(settings.ALPHABET))];
	}

	query = settings.UPDATE_USER_TOKEN_QUERY;
	_, err = dbConnection.ExecContext(
		context.Background(),
		query,
		string(newToken),
		user.Email);
	if err != nil {
		serviceResponseLogin.IsSuccessful = false;
		serviceResponseLogin.ErrorMessage = "Unable to login: " + err.Error();
		return serviceResponseLogin;
	}
	serviceResponseLogin.Token = string(newToken);
	log.Printf("updated users for the following id: %d", user.Email);

	return serviceResponseLogin;
}

func EditUser(user models.User) models.ServiceResponse {
	serviceResponse := models.ServiceResponse{
		IsSuccessful: true,
		ErrorMessage: "",
	}
	recievedUserWithEmailAlias := get.GetUserByEmail(user.EmailAlias);
	if (recievedUserWithEmailAlias.Email == "") {
		recievedUserWithEmail := get.GetUserByEmail(user.Email);
		if (recievedUserWithEmail.Email != user.Email) {
			panic("Someone else has same email");
		}
	} else if (recievedUserWithEmailAlias.Email != user.Email) {
		panic("Someone else has same email alias");
	}

	dbConnection := services.ConnectToDB();
	defer dbConnection.Close();

	query := settings.UPDATE_USER_INFORMATION_QUERY;
	_, err := dbConnection.ExecContext(
		context.Background(),
		query,
		user.EmailAlias,
		user.FirstName,
		user.LastName,
		user.PhoneNumber,
		user.HomeAddress,
		user.MustChangePW,
		user.Email);
	if err != nil {
		serviceResponse.IsSuccessful = false;
		serviceResponse.ErrorMessage = "Unable to update user: " + err.Error();
		return serviceResponse;
	}
	log.Printf("updated users for the following email: %d", user.Email);

	return serviceResponse;
}