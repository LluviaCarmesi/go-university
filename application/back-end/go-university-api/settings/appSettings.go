package settings;

// Paths
const API_ROOT = "/api";
const COURSES_PATH = API_ROOT + "/courses";
const USERS_PATH = API_ROOT + "/users";
const STUDENTS_PATH = USERS_PATH + "/students";
const APPOINTMENTS_PATH = API_ROOT + "/appointments";
const SEMESTERS_PATH = API_ROOT + "/semesters";
const TAUGHT_COURSES_PATH = API_ROOT + "/taught_courses";
const REGISTRATIONS_PATH = API_ROOT + "/registrations";
const COURSE_SCHEDULES_PATH = API_ROOT + "/course_schedules";

// Queries
const GET_COURSES_QUERY = "SELECT id, name, description FROM courses";
const INSERT_COURSE_QUERY = "INSERT INTO COURSES (id, name, description) ";
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
const GET_TAUGHT_COURSES_QUERY = 
	"SELECT id,course_id,semester_id,professor_email,max_students,location " +
	"FROM taught_courses";
const GET_REGISTRATIONS_QUERY =
	"SELECT id,student_email,taught_course_id,final_grade,status " +
	"FROM registrations";
const GET_COURSE_SCHEDULES_QUERY =
	"SELECT id,taught_course_id,start_time,end_time " +
	"FROM course_schedules";