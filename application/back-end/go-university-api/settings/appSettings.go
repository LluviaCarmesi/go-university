package settings;

// Paths
const API_ROOT = "/api/";
const COURSES_PATH = API_ROOT + "courses/";
const USERS_PATH = API_ROOT + "users/";
const STUDENTS_PATH = USERS_PATH + "students/";
const APPOINTMENTS_PATH = API_ROOT + "appointments/";
const SEMESTERS_PATH = API_ROOT + "semesters/";
const TAUGHT_COURSES_PATH = API_ROOT + "taught-courses/";
const REGISTRATIONS_PATH = API_ROOT + "registrations/";
const DEPARTMENTS_PATH = API_ROOT + "departments/";
const PROFESSORS_IN_DEPARTMENTS_PATH = API_ROOT + "professors-in-departments/";
const USERS_LOGIN_PATH = USERS_PATH + "login/";
const USERS_BY_TOKEN_PATH = USERS_PATH + "by_token/";

// Queries
const GET_COURSES_QUERY =
	"SELECT id,name,description,credits FROM courses";
const INSERT_COURSE_QUERY =
	"INSERT INTO courses (id,name,description,credits) " +
	"VALUES (?, ?, ?, ?)";
const UPDATE_COURSE_QUERY =
	"UPDATE courses SET name = ?, description = ?, credits = ? WHERE id = ?";
const DELETE_COURSE_QUERY =
	"DELETE FROM courses WHERE id = ?";
const GET_USERS_QUERY =
	"SELECT email,email_alias,first_name,last_name,phone_number,home_address,role " +
	"FROM users";
const INSERT_USER_QUERY = 
	"INSERT INTO users (email,email_alias,password,first_name,last_name,phone_number," +
	"home_address,role,must_change_pw) " +
	"VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)";
const UPDATE_USER_INFORMATION_QUERY =
	"UPDATE users SET email_alias = ?, first_name = ?, last_name = ?, phone_number = ?, " +
	"home_address = ?, must_change_pw = ? WHERE email = ?";
const UPDATE_USER_PASSWORD_QUERY =
	"UPDATE users SET password = ? WHERE email = ?";
const UPDATE_USER_TOKEN_QUERY =
	"UPDATE users SET token = ? WHERE email = ?";
const DELETE_USER_QUERY =
	"DELETE FROM users WHERE email = ?";
const GET_APPOINTMENTS_QUERY =
	"SELECT id,student_email,admin_email,is_complete,start_time,end_time " +
	"FROM appointments";
const INSERT_APPOINTMENT_QUERY =
	"INSERT INTO appointments (student_email,admin_email,is_complete,start_time,end_time) " +
	"VALUES (?, ?, ?, ?, ?)";
const UPDATE_APPOINTMENT_QUERY =
	"UPDATE appointments SET student_email = ?, admin_email = ?, is_complete = ?, start_time = ?, " +
	"end_time = ? WHERE id = ?";
const DELETE_APPOINTMENT_QUERY =
	"DELETE FROM appointments WHERE id = ?";
const GET_SEMESTERS_QUERY =
	"SELECT name,start_date,end_date " +
	"FROM semesters";
const INSERT_SEMESTER_QUERY = 
	"INSERT INTO semesters (name,start_date,end_date) VALUES (?, ?, ?)";
const UPDATE_SEMESTER_QUERY =
	"UPDATE semesters SET start_date = ?, end_date = ? WHERE name = ?";
const DELETE_SEMESTER_QUERY =
	"DELETE FROM semesters WHERE name = ?";
const GET_TAUGHT_COURSES_QUERY = 
	"SELECT id,course_id,semester_name,professor_email,max_students,location,day,start_time,end_time " +
	"FROM taught_courses";
const INSERT_TAUGHT_COURSE_QUERY =
	"INSERT INTO taught_courses (course_id,semester_name,professor_email,max_students,location," +
	"day,start_time,end_time) VALUES (?, ?, ?, ?, ?, ?, ?, ?)";
const UPDATE_TAUGHT_COURSE_QUERY =
	"UPDATE taught_courses SET course_id = ?, professor_email = ?, max_students = ?, location = ?, " +
	"day = ?, start_time = ?, end_time = ? " +
	"WHERE id = ?";
const DELETE_TAUGHT_COURSE_QUERY = 
	"DELETE FROM taught_courses WHERE id = ?";
const GET_REGISTRATIONS_QUERY =
	"SELECT student_email,taught_course_id,final_grade,status " +
	"FROM registrations";
const INSERT_REGISTRATION_QUERY =
	"INSERT INTO registrations (student_email,taught_course_id,final_grade,status) " +
	"VALUES (?, ?, ?, ?)";
const UPDATE_REGISTRATION_QUERY =
	"UPDATE registrations SET final_grade = ?, status = ?, " +
	"WHERE student_email = ? AND taught_course_id = ?";
const DELETE_REGISTRATION_QUERY = 
	"DELETE FROM registrations WHERE student_email = ? AND taught_course_id = ?";
const GET_DEPARTMENTS_QUERY =
	"SELECT id,name FROM departments";
const INSERT_DEPARTMENT_QUERY =
	"INSERT INTO departments (name) VALUES (?)";
const UPDATE_DEPARTMENT_QUERY =
	"UPDATE departments SET name = ? WHERE id = ?";
const DELETE_DEPARTMENT_QUERY =
	"DELETE FROM departments WHERE id = ?";
const GET_PROFESSORS_IN_DEPARTMENTS_QUERY =
	"SELECT professor_email,department_id,is_leader FROM professors_in_departments";
const INSERT_PROFESSOR_IN_DEPARTMENT_QUERY =
	"INSERT INTO professors_in_departments (professor_email,department_id,is_leader) " +
	"VALUES (?, ?, ?)";
const UPDATE_PROFESSORS_IN_DEPARTMENT_QUERY =
	"UPDATE professors_in_departments SET is_leader = ? " +
	"WHERE professor_email = ? AND department_id = ?";
const DELETE_PROFESSOR_IN_DEPARTMENT_QUERY =
	"DELETE FROM professors_in_departments WHERE professor_email = ? AND department_id = ?";

// Strings
const ALPHABET = "abcdefghijlmnopqrstuvwxyzABCDEFGHIJLMNOPQRSTUVWXYZ";