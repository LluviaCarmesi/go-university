package main;
import "back-end/services/get";
import "back-end/services/post";
import "net/http";
import "log";
import "back-end/settings";
import "encoding/json";
import "back-end/models";

func enableSettings(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Headers", "*");
	(*w).Header().Set("Access-Control-Allow-Origin", "http://localhost:5173");
	(*w).Header().Set("Access-Control-Allow-Methods", "*");
}

func courses(w http.ResponseWriter, r *http.Request) {
	enableSettings(&w);
	switch (r.Method) {
		case http.MethodGet:
			json.NewEncoder(w).Encode(get.Courses());
			break;
		case http.MethodPost:
			var course models.Course;
			err := json.NewDecoder(r.Body).Decode(&course);
			if (err != nil) {
				panic("Course not structured properly: " + err.Error());
			}
			postResponse := post.AddCourse(course);
			json.NewEncoder(w).Encode(postResponse);
			break;
		case http.MethodOptions:
			break;
		default:
			panic("Method not supported");
	}

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

func getTaughtCourses (w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(get.TaughtCourses());
}

func getRegistrations (w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(get.Registrations());
}

func getCourseSchedules (w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(get.CourseSchedules());
}

func main() {
	http.HandleFunc(settings.STUDENTS_PATH, getStudents);
	http.HandleFunc(settings.COURSES_PATH, courses);
	http.HandleFunc(settings.USERS_PATH, getUsers);
	http.HandleFunc(settings.APPOINTMENTS_PATH, getAppointments);
	http.HandleFunc(settings.SEMESTERS_PATH, getSemesters);
	http.HandleFunc(settings.TAUGHT_COURSES_PATH, getTaughtCourses);
	http.HandleFunc(settings.REGISTRATIONS_PATH, getRegistrations);
	http.HandleFunc(settings.COURSE_SCHEDULES_PATH, getCourseSchedules);

	log.Fatal(http.ListenAndServe(":8080", nil));
}