package models;

import "time";

type Course struct {
	ID int `json:id`
	Name string `json:"name"`
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