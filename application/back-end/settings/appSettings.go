package settings;

// Paths
const API_ROOT = "/api";
const COURSES_PATH = API_ROOT + "/courses";
const USERS_PATH = API_ROOT + "/users";
const STUDENTS_PATH = USERS_PATH + "/students";
const APPOINTMENTS_PATH = API_ROOT + "/appointments";
const SEMESTERS_PATH = API_ROOT + "/semesters";

// Queries
const GET_USERS_QUERY =
	"SELECT email,email_alias,first_name,last_name,phone_number,home_address,role " +
	"FROM users";
const GET_STUDENTS_QUERY =
	"SELECT email,email_alias,first_name,last_name,phone_number,home_address,role " +
	"FROM users " + 
	"WHERE role = 'student'";
const GET_APPOINTMENTS_QUERY =
	"SELECT id,student_email,admin_email,is_complete,start_time,end_time " +
	"FROM appointments";
const GET_SEMESTERS_QUERY =
	"SELECT id,name,start_date,end_date " +
	"FROM semesters";