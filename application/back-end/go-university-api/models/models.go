package models;

import "time";

type Course struct {
	ID string `json:id`
	Name string `json:name`
	Description string `json:description`
	Credits int `json:credits`
}

type Department struct {
	ID int `json:id`
	Name string `json:name`
}

type ProfessorInDepartment struct {
	ProfessorEmail string `json:professor_email`
	DepartmentID int `json:department_id`
	DepartmentName string `json:department_name`
	IsLeader bool `json:is_leader`
}

type User struct {
	Email string `json:email`
	EmailAlias string `json:email_alias`
	Password string `json:password`
	FirstName string `json:first_name`
	LastName string `json:last_name`
	PhoneNumber string `json:phone_number`
	HomeAddress string `json:home_address`
	Role string `json:role`
	MustChangePW bool `json:must_change_pw`
	Token string `json:token`
}

type Appointment struct {
	ID int `json:id`
	StudentEmail string `json:student_email`
	AdminEmail string `json:admin_email`
	IsComplete bool `json:is_complete`
	StartTime time.Time `json:start_time`
	EndTime time.Time `json:end_time`
	Description string `json:description`
}

type Semester struct {
	Name string `json:name`
	StartDate time.Time `json:start_date`
	EndDate time.Time `json:end_date`
}

type TaughtCourse struct {
	ID int `json:id`
	CourseID string `json:course_id`
	SemesterName string `json:semester_name`
	ProfessorEmail string `json:professor_email`
	MaxStudents int `json:max_students`
	Location string `json:location`
	Day string `json:day`
	StartTime string `json:start_time`
	EndTime string `json:end_time`
}

type Registration struct {
	StudentEmail string `json:student_email`
	TaughtCourseID int `json:taught_course_id`
	CourseID string `json:course_id`
	SemesterName string `json:semester_name`
	FinalGrade float32 `json:final_grade`
	Status string `json:status`
}

type ServiceResponse struct {
	IsSuccessful bool `json:is_successful`
	ErrorMessage string `json:error_message`
}

type ServiceResponseLogin struct {
	IsSuccessful bool `json:is_successful`
	ErrorMessage string `json:error_message`
	Token string `json:token`
	MustChangePW bool `json:must_change_pw`
	Email string `json:email`
}