package models;

import "time";

type Course struct {
	ID string `json:id`
	Name string `json:"name"`
	Description string `json:description`
}

type User struct {
	Email string `json:email`
	EmailAlias string `json:email_alias`
	FirstName string `json:first_name`
	LastName string `json:last_name`
	PhoneNumber string `json:phone_number`
	HomeAddress string `json:home_address`
	Role string `json:role`
}

type Appointment struct {
	ID int `json:id`
	StudentEmail string `json:student_email`
	AdminEmail string `json:admin_email`
	IsComplete bool `json:is_complete`
	StartTime time.Time `json:start_time`
	EndTime time.Time `json:end_time`
}

type Semester struct {
	ID int `json:id`
	Name string `json:name`
	StartDate time.Time `json:start_date`
	EndDate time.Time `json:end_time`
}

type TaughtCourse struct {
	ID int `json:id`
	CourseID string `json:course_id`
	SemesterID int `json:semester_id`
	ProfessorEmail string `json:professor_email`
	MaxStudents int `json:max_students`
	Location string `json:location`
}

type Registration struct {
	ID int `json:id`
	StudentEmail string `json:student_email`
	TaughtCourseID int `json:taught_course_id`
	FinalGrade float32 `json:final_grade`
	Status string `json:status`
}

type CourseSchedule struct {
	ID int `json:id`
	TaughtCourseID int `json:taught_course_id`
	StartTime time.Time `json:start_time`
	EndTime time.Time `json:end_time`
}

type ServiceResponse struct {
	IsSuccessful bool `json:is_successful`
	ErrorMessage string `json:error_message`
}