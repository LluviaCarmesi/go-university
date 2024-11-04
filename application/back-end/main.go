package main;
import "back-end/services/get";
import "net/http";
import "log";
import "back-end/settings";
import "encoding/json";

func testGET(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(get.Courses());
}

func main() {
	http.HandleFunc(settings.STUDENTS_PATH, testGET);
	log.Fatal(http.ListenAndServe(":8081", nil));
}