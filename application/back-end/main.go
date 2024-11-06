package main;
import "back-end/services/get";
import "net/http";
import "log";
import "back-end/settings";
import "encoding/json";

func getCourses(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(get.Courses());
}

func getStudents(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(get.Students());
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(get.Users());
}

func getAppointments(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(get.Appointments());
}

func getSemesters (w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(get.Semesters());
}

func main() {
	http.HandleFunc(settings.STUDENTS_PATH, getStudents);
	http.HandleFunc(settings.COURSES_PATH, getCourses);
	http.HandleFunc(settings.USERS_PATH, getUsers);
	http.HandleFunc(settings.APPOINTMENTS_PATH, getAppointments);
	http.HandleFunc(settings.SEMESTERS_PATH, getSemesters);
	log.Fatal(http.ListenAndServe(":8080", nil));
}