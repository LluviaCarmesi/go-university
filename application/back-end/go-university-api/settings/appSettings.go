package settings;

// Paths
const API_ROOT = "/api";
const COURSES_PATH = API_ROOT + "/courses";
const USERS_PATH = API_ROOT + "/users/";
const STUDENTS_PATH = USERS_PATH + "/students";
const APPOINTMENTS_PATH = API_ROOT + "/appointments";
const SEMESTERS_PATH = API_ROOT + "/semesters";
const TAUGHT_COURSES_PATH = API_ROOT + "/taught_courses";
const REGISTRATIONS_PATH = API_ROOT + "/registrations";
const COURSE_SCHEDULES_PATH = API_ROOT + "/course_schedules";

// Queries
const GET_COURSES_QUERY =
	"SELECT id, name, description FROM courses";
const INSERT_COURSE_QUERY =
	"INSERT INTO courses (id,name,description) ";
const UPDATE_COURSE_QUERY =
	"UPDATE courses SET name = ?, description = ? WHERE id = ?";
const DELETE_COURSE_QUERY =
	"DELETE FROM courses WHERE id = ?";
const GET_USERS_QUERY =
	"SELECT email,email_alias,first_name,last_name,phone_number,home_address,role " +
	"FROM users";
const INSERT_USER_QUERY = 
	"INSERT INTO users (email,email_alias,password,first_name,last_name,phone_number," +
	"home_address,role) ";
const GET_STUDENTS_QUERY =
	"SELECT email,email_alias,first_name,last_name,phone_number,home_address,role " +
	"FROM users " + 
	"WHERE role = 'student'";
const GET_APPOINTMENTS_QUERY =
	"SELECT id,student_email,admin_email,is_complete,start_time,end_time " +
	"FROM appointments";
const INSERT_APPOINTMENT_QUERY =
	"INSERT INTO appointments (student_email,admin_email,start_time,end_time) "
const GET_SEMESTERS_QUERY =
	"SELECT id,name,start_date,end_date " +
	"FROM semesters";
const INSERT_SEMESTER_QUERY = 
	"INSERT INTO semesters (name,start_date,end_date) ";
const GET_TAUGHT_COURSES_QUERY = 
	"SELECT id,course_id,semester_id,professor_email,max_students,location " +
	"FROM taught_courses";
const INSERT_TAUGHT_COURSE_QUERY =
	"INSERT INTO taught_courses (course_id,semester_id,professer_email,max_students,location) ";
const GET_REGISTRATIONS_QUERY =
	"SELECT id,student_email,taught_course_id,final_grade,status " +
	"FROM registrations";
const INSERT_REGISTRATION_QUERY =
	"INSERT INTO registrations (student_email,taught_course_id,final_grade,status) "
const GET_COURSE_SCHEDULES_QUERY =
	"SELECT id,taught_course_id,start_time,end_time " +
	"FROM course_schedules";
const INSERT_COURSE_SCHEDULE_QUERY =
	"INSERT INTO course_schedules (taught_course_id,start_time,end_time) ";
const GET_DEPARTMENTS_QUERY =
	"SELECT id,name FROM departments";
const INSERT_DEPARTMENT_QUERY =
	"INSERT INTO departments (name) ";
const GET_PROFESSORS_IN_DEPARTMENTS_QUERY =
	"SELECT professor_email,department_id,is_leader FROM professors_in_departments";
const INSERT_PROFESSOR_IN_DEPARTMENT_QUERY =
	"INSERT INTO professors_in_departments (professor_email,department_id,is_leader) ";