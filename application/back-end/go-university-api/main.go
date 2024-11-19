package main;
import "back-end/services/get";
import "back-end/services/post";
import "back-end/services/patch";
import "net/http";
import "log";
import "back-end/settings";
import "encoding/json";
import "back-end/models";
import "strings";

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
			json.NewEncoder(w).Encode(patchResponse);
			break;
		case http.MethodDelete:
			pathParts := strings.Split(r.URL.Path, "/")
			if len(pathParts) < 3 {
				panic("Need ID in path");
			}
			courseID := pathParts[2];
			if (courseID == "") {

			}
			break;
		case http.MethodOptions:
			break;
		default:
			panic("Method not supported");
	}
}

func getStudents(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(get.GetStudents());
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(get.GetUsers());
}

func getAppointments(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(get.GetAppointments());
}

func getSemesters (w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(get.GetSemesters());
}

func getTaughtCourses (w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(get.GetTaughtCourses());
}

func getRegistrations (w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(get.GetRegistrations());
}

func main() {
	http.HandleFunc(settings.STUDENTS_PATH, getStudents);
	http.HandleFunc(settings.COURSES_PATH, courses);
	http.HandleFunc(settings.USERS_PATH, getUsers);
	http.HandleFunc(settings.APPOINTMENTS_PATH, getAppointments);
	http.HandleFunc(settings.SEMESTERS_PATH, getSemesters);
	http.HandleFunc(settings.TAUGHT_COURSES_PATH, getTaughtCourses);
	http.HandleFunc(settings.REGISTRATIONS_PATH, getRegistrations);

	log.Fatal(http.ListenAndServe(":8080", nil));
}