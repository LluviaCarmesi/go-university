package main;
import "back-end/services/get";
import "back-end/services/post";
import "back-end/services/patch";
import "back-end/services/delete";
import "net/http";
import "log";
import "back-end/settings";
import "encoding/json";
import "back-end/models";
import "strings";
import "strconv";
import "fmt";

func enableSettings(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Headers", "*");
	(*w).Header().Set("Access-Control-Allow-Origin", "http://localhost:5173");
	(*w).Header().Set("Access-Control-Allow-Methods", "*");
}

func courses(w http.ResponseWriter, r *http.Request) {
	enableSettings(&w);
	switch (r.Method) {
		case http.MethodGet:
			pathParts := strings.Split(r.URL.Path, "/");
			courseID := pathParts[3];
			if (courseID == "") {
				json.NewEncoder(w).Encode(get.GetCourses());
				return;
			}
			json.NewEncoder(w).Encode(get.GetCourseByID(courseID));
			break;
		case http.MethodPost:
			var course models.Course;
			err := json.NewDecoder(r.Body).Decode(&course);
			if (err != nil) {
				panic("Course not structured properly: " + err.Error());
			}
			postResponse := post.AddCourse(course);
			if (!postResponse.IsSuccessful) {
				fmt.Println(postResponse.IsSuccessful);
				w.WriteHeader(http.StatusBadRequest);
			}
			json.NewEncoder(w).Encode(postResponse);
			break;
		case http.MethodPatch:
			pathParts := strings.Split(r.URL.Path, "/");
			if len(pathParts) < 4 {
				panic("Need ID in path");
			}
			courseID := pathParts[3];
			if (courseID == "") {
				panic("Need ID in path");
			}
			var course models.Course;
			err := json.NewDecoder(r.Body).Decode(&course);
			if (err != nil) {
				panic("Course not structured properly: " + err.Error());
			}
			course.ID = courseID;
			patchResponse := patch.EditCourse(course);
			if (!patchResponse.IsSuccessful) {
				fmt.Println(patchResponse.IsSuccessful);
				w.WriteHeader(http.StatusBadRequest);
			}
			json.NewEncoder(w).Encode(patchResponse);
			break;
		case http.MethodDelete:
			pathParts := strings.Split(r.URL.Path, "/");
			if len(pathParts) < 4 {
				panic("Need ID in path");
			}
			courseID := pathParts[3];
			if (courseID == "") {
				panic("Need ID in path");
			}
			deleteResponse := delete.DeleteCourse(courseID);
			if (!deleteResponse.IsSuccessful) {
				fmt.Println(deleteResponse.IsSuccessful);
				w.WriteHeader(http.StatusBadRequest);
			}
			json.NewEncoder(w).Encode(deleteResponse);
			break;
		case http.MethodOptions:
			break;
		default:
			panic("Method not supported");
	}
}

func users(w http.ResponseWriter, r *http.Request) {
	enableSettings(&w);
	switch (r.Method) {
		case http.MethodGet:
			pathParts := strings.Split(r.URL.Path, "/");
			email := pathParts[3];
			if (email == "") {
				json.NewEncoder(w).Encode(get.GetUsers());
				return;
			}
			json.NewEncoder(w).Encode(get.GetUserByEmail(email));
			break;
		case http.MethodPost:
			var user models.User;
			err := json.NewDecoder(r.Body).Decode(&user);
			if (err != nil) {
				panic("User not structured properly: " + err.Error());
			}
			postResponse := post.AddUser(user);
			if (!postResponse.IsSuccessful) {
				fmt.Println(postResponse.IsSuccessful);
				w.WriteHeader(http.StatusBadRequest);
			}
			json.NewEncoder(w).Encode(postResponse);
			break;
		case http.MethodPatch:
			pathParts := strings.Split(r.URL.Path, "/");
			if len(pathParts) < 4 {
				panic("Need Email in path");
			}
			email := pathParts[3];
			if (email == "") {
				panic("Need Email in path");
			}
			var user models.User;
			err := json.NewDecoder(r.Body).Decode(&user);
			if (err != nil) {
				panic("User not structured properly: " + err.Error());
			}
			user.Email = email;
			patchResponse := patch.EditUser(user);
			if (!patchResponse.IsSuccessful) {
				fmt.Println(patchResponse.IsSuccessful);
				w.WriteHeader(http.StatusBadRequest);
			}
			json.NewEncoder(w).Encode(patchResponse);
			break;
		case http.MethodDelete:
			pathParts := strings.Split(r.URL.Path, "/");
			if len(pathParts) < 4 {
				panic("Need Email in path");
			}
			email := pathParts[3];
			if (email == "") {
				panic("Need Email in path");
			}
			deleteResponse := delete.DeleteUser(email);
			if (!deleteResponse.IsSuccessful) {
				fmt.Println(deleteResponse.IsSuccessful);
				w.WriteHeader(http.StatusBadRequest);
			}
			json.NewEncoder(w).Encode(deleteResponse);
			break;
		case http.MethodOptions:
			break;
		default:
			panic("Method not supported");
	}
}

func loginUser(w http.ResponseWriter, r *http.Request) {
	enableSettings(&w);
	switch (r.Method) {
		case http.MethodPatch:
			var user models.User;
	
			err := json.NewDecoder(r.Body).Decode(&user);
			if (err != nil) {
				panic("User not structured properly: " + err.Error());
			}
			pathParts := strings.Split(r.URL.Path, "/");
			if len(pathParts) < 5 {
				panic("Need Email in path");
			}
			userEmail := pathParts[4];
			if (userEmail == "") {
				return;
			}
			user.Email = userEmail;
			patchResponse := patch.LoginUser(user);
			if (!patchResponse.IsSuccessful) {
				fmt.Println(patchResponse.IsSuccessful);
				w.WriteHeader(http.StatusBadRequest);
			}
			json.NewEncoder(w).Encode(patchResponse);
			break;
		case http.MethodOptions:
			break;
		default:
			panic("Method not supported");
	}
}

func getUserByToken(w http.ResponseWriter, r *http.Request) {
	enableSettings(&w);
	switch (r.Method) {
		case http.MethodGet:
			pathParts := strings.Split(r.URL.Path, "/");
			if len(pathParts) < 5 {
				panic("Need Token in path");
			}
			token := pathParts[4];
			if (token == "") {
				return;
			}
			json.NewEncoder(w).Encode(get.GetUserByToken(token));
			break;
		default:
			panic("Method not supported");
	}
}

func appointments(w http.ResponseWriter, r *http.Request) {
	enableSettings(&w);
	switch (r.Method) {
		case http.MethodGet:
			pathParts := strings.Split(r.URL.Path, "/");
			appointmentID := pathParts[3];
			if (appointmentID == "") {
				json.NewEncoder(w).Encode(get.GetAppointments());
				return;
			}
			json.NewEncoder(w).Encode(get.GetAppointmentByID(appointmentID));
			break;
		case http.MethodPost:
			var appointment models.Appointment;
			err := json.NewDecoder(r.Body).Decode(&appointment);
			if (err != nil) {
				panic("Appointment not structured properly: " + err.Error());
			}
			postResponse := post.AddAppointment(appointment);
			if (!postResponse.IsSuccessful) {
				fmt.Println(postResponse.IsSuccessful);
				w.WriteHeader(http.StatusBadRequest);
			}
			json.NewEncoder(w).Encode(postResponse);
			break;
		case http.MethodPatch:
			pathParts := strings.Split(r.URL.Path, "/");
			if len(pathParts) < 4 {
				panic("Need ID in path");
			}
			appointmentID := pathParts[3];
			if (appointmentID == "") {
				panic("Need ID in path");
			}
			var appointment models.Appointment;
			err := json.NewDecoder(r.Body).Decode(&appointment);
			if (err != nil) {
				panic("Appointment not structured properly: " + err.Error());
			}
			appointmentIDInt, err := strconv.Atoi(appointmentID);
			if err != nil {
				panic("appointmentID is not an int");
			}
			appointment.ID = appointmentIDInt;
			patchResponse := patch.EditAppointment(appointment);
			if (!patchResponse.IsSuccessful) {
				fmt.Println(patchResponse.IsSuccessful);
				w.WriteHeader(http.StatusBadRequest);
			}
			json.NewEncoder(w).Encode(patchResponse);
			break;
		case http.MethodDelete:
			pathParts := strings.Split(r.URL.Path, "/")
			if len(pathParts) < 4 {
				panic("Need ID in path");
			}
			appointmentID := pathParts[3];
			if (appointmentID == "") {
				panic("Need ID in path");
			}
			appointmentIDInt, err := strconv.Atoi(appointmentID);
			if err != nil {
				panic("appointmentID is not an int");
			}
			deleteResponse := delete.DeleteAppointment(appointmentIDInt);
			if (!deleteResponse.IsSuccessful) {
				fmt.Println(deleteResponse.IsSuccessful);
				w.WriteHeader(http.StatusBadRequest);
			}
			json.NewEncoder(w).Encode(deleteResponse);
			break;
		case http.MethodOptions:
			break;
		default:
			panic("Method not supported");
	}
}

func semesters(w http.ResponseWriter, r *http.Request) {
	enableSettings(&w);
	switch (r.Method) {
		case http.MethodGet:
			pathParts := strings.Split(r.URL.Path, "/");
			name := pathParts[3];
			if (name == "") {
				json.NewEncoder(w).Encode(get.GetSemesters());
				return;
			}
			json.NewEncoder(w).Encode(get.GetSemesterByName(name));
			break;
		case http.MethodPost:
			var semester models.Semester;
			err := json.NewDecoder(r.Body).Decode(&semester);
			if (err != nil) {
				panic("Semester not structured properly: " + err.Error());
			}
			postResponse := post.AddSemester(semester);
			if (!postResponse.IsSuccessful) {
				fmt.Println(postResponse.IsSuccessful);
				w.WriteHeader(http.StatusBadRequest);
			}
			json.NewEncoder(w).Encode(postResponse);
			break;
		case http.MethodPatch:
			pathParts := strings.Split(r.URL.Path, "/");
			if len(pathParts) < 4 {
				panic("Need Name in path");
			}
			name := pathParts[3];
			if (name == "") {
				panic("Need Name in path");
			}
			var semester models.Semester;
			err := json.NewDecoder(r.Body).Decode(&semester);
			if (err != nil) {
				panic("Semester not structured properly: " + err.Error());
			}
			semester.Name = name;
			patchResponse := patch.EditSemester(semester);
			if (!patchResponse.IsSuccessful) {
				fmt.Println(patchResponse.IsSuccessful);
				w.WriteHeader(http.StatusBadRequest);
			}
			json.NewEncoder(w).Encode(patchResponse);
			break;
		case http.MethodDelete:
			pathParts := strings.Split(r.URL.Path, "/")
			if len(pathParts) < 4 {
				panic("Need Name in path");
			}
			name := pathParts[3];
			if (name == "") {
				panic("Need Name in path");
			}
			deleteResponse := delete.DeleteSemester(name);
			if (!deleteResponse.IsSuccessful) {
				fmt.Println(deleteResponse.IsSuccessful);
				w.WriteHeader(http.StatusBadRequest);
			}
			json.NewEncoder(w).Encode(deleteResponse);
			break;
		case http.MethodOptions:
			break;
		default:
			panic("Method not supported");
	}
}

func taughtCourses (w http.ResponseWriter, r *http.Request) {
	enableSettings(&w);
	switch (r.Method) {
		case http.MethodGet:
			pathParts := strings.Split(r.URL.Path, "/");
			appointmentID := pathParts[3];
			if (appointmentID == "") {
				json.NewEncoder(w).Encode(get.GetAppointments());
				return;
			}
			json.NewEncoder(w).Encode(get.GetAppointmentByID(appointmentID));
			break;
		case http.MethodPost:
			var appointment models.Appointment;
			err := json.NewDecoder(r.Body).Decode(&appointment);
			if (err != nil) {
				panic("Appointment not structured properly: " + err.Error());
			}
			postResponse := post.AddAppointment(appointment);
			if (!postResponse.IsSuccessful) {
				fmt.Println(postResponse.IsSuccessful);
				w.WriteHeader(http.StatusBadRequest);
			}
			json.NewEncoder(w).Encode(postResponse);
			break;
		case http.MethodPatch:
			pathParts := strings.Split(r.URL.Path, "/");
			if len(pathParts) < 4 {
				panic("Need ID in path");
			}
			appointmentID := pathParts[3];
			if (appointmentID == "") {
				panic("Need ID in path");
			}
			var appointment models.Appointment;
			err := json.NewDecoder(r.Body).Decode(&appointment);
			if (err != nil) {
				panic("Appointment not structured properly: " + err.Error());
			}
			appointmentIDInt, err := strconv.Atoi(appointmentID);
			if err != nil {
				panic("appointmentID is not an int");
			}
			appointment.ID = appointmentIDInt;
			patchResponse := patch.EditAppointment(appointment);
			if (!patchResponse.IsSuccessful) {
				fmt.Println(patchResponse.IsSuccessful);
				w.WriteHeader(http.StatusBadRequest);
			}
			json.NewEncoder(w).Encode(patchResponse);
			break;
		case http.MethodDelete:
			pathParts := strings.Split(r.URL.Path, "/")
			if len(pathParts) < 4 {
				panic("Need ID in path");
			}
			appointmentID := pathParts[3];
			if (appointmentID == "") {
				panic("Need ID in path");
			}
			appointmentIDInt, err := strconv.Atoi(appointmentID);
			if err != nil {
				panic("appointmentID is not an int");
			}
			deleteResponse := delete.DeleteAppointment(appointmentIDInt);
			if (!deleteResponse.IsSuccessful) {
				fmt.Println(deleteResponse.IsSuccessful);
				w.WriteHeader(http.StatusBadRequest);
			}
			json.NewEncoder(w).Encode(deleteResponse);
			break;
		case http.MethodOptions:
			break;
		default:
			panic("Method not supported");
	}
}

func registrations (w http.ResponseWriter, r *http.Request) {
	enableSettings(&w);
	switch (r.Method) {
		case http.MethodGet:
			pathParts := strings.Split(r.URL.Path, "/");
			appointmentID := pathParts[3];
			if (appointmentID == "") {
				json.NewEncoder(w).Encode(get.GetAppointments());
				return;
			}
			json.NewEncoder(w).Encode(get.GetAppointmentByID(appointmentID));
			break;
		case http.MethodPost:
			var appointment models.Appointment;
			err := json.NewDecoder(r.Body).Decode(&appointment);
			if (err != nil) {
				panic("Appointment not structured properly: " + err.Error());
			}
			postResponse := post.AddAppointment(appointment);
			if (!postResponse.IsSuccessful) {
				fmt.Println(postResponse.IsSuccessful);
				w.WriteHeader(http.StatusBadRequest);
			}
			json.NewEncoder(w).Encode(postResponse);
			break;
		case http.MethodPatch:
			pathParts := strings.Split(r.URL.Path, "/");
			if len(pathParts) < 4 {
				panic("Need ID in path");
			}
			appointmentID := pathParts[3];
			if (appointmentID == "") {
				panic("Need ID in path");
			}
			var appointment models.Appointment;
			err := json.NewDecoder(r.Body).Decode(&appointment);
			if (err != nil) {
				panic("Appointment not structured properly: " + err.Error());
			}
			appointmentIDInt, err := strconv.Atoi(appointmentID);
			if err != nil {
				panic("appointmentID is not an int");
			}
			appointment.ID = appointmentIDInt;
			patchResponse := patch.EditAppointment(appointment);
			if (!patchResponse.IsSuccessful) {
				fmt.Println(patchResponse.IsSuccessful);
				w.WriteHeader(http.StatusBadRequest);
			}
			json.NewEncoder(w).Encode(patchResponse);
			break;
		case http.MethodDelete:
			pathParts := strings.Split(r.URL.Path, "/")
			if len(pathParts) < 4 {
				panic("Need ID in path");
			}
			appointmentID := pathParts[3];
			if (appointmentID == "") {
				panic("Need ID in path");
			}
			appointmentIDInt, err := strconv.Atoi(appointmentID);
			if err != nil {
				panic("appointmentID is not an int");
			}
			deleteResponse := delete.DeleteAppointment(appointmentIDInt);
			if (!deleteResponse.IsSuccessful) {
				fmt.Println(deleteResponse.IsSuccessful);
				w.WriteHeader(http.StatusBadRequest);
			}
			json.NewEncoder(w).Encode(deleteResponse);
			break;
		case http.MethodOptions:
			break;
		default:
			panic("Method not supported");
	}
}

func main() {
	http.HandleFunc(settings.COURSES_PATH, courses);
	http.HandleFunc(settings.USERS_PATH, users);
	http.HandleFunc(settings.APPOINTMENTS_PATH, appointments);
	http.HandleFunc(settings.SEMESTERS_PATH, semesters);
	http.HandleFunc(settings.TAUGHT_COURSES_PATH, taughtCourses);
	http.HandleFunc(settings.REGISTRATIONS_PATH, registrations);
	http.HandleFunc(settings.USERS_BY_TOKEN_PATH, getUserByToken);
	http.HandleFunc(settings.USERS_LOGIN_PATH, loginUser);

	log.Fatal(http.ListenAndServe(":8080", nil));
}