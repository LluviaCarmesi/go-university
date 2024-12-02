// URIs
export const SERVER = "http://localhost:8080/api/";
export const COURSES_API_URI = SERVER + "courses/";
export const APPOINTMENTS_API_URI = SERVER + "appointments/";
export const DEPARTMENTS_API_URI = SERVER + "departments/";
export const PROFESSORS_IN_DEPARTMENTS_API_URI = SERVER + "professors-in-departments/";
export const USERS_API_URI = SERVER + "users/";
export const USERS_BY_TOKEN_API_URI = USERS_API_URI + "by_token/";
export const USERS_LOGIN_API_URI = USERS_API_URI + "login/";
export const SEMESTERS_API_URI = SERVER + "semesters/";
export const TAUGHT_COURSES_API_URI = SERVER + "taught-courses/";
export const REGISTRATIONS_API_URI = SERVER + "registrations/";

// Options
export const ROLE_OPTIONS = [
    {
        value: "admin",
        label: "Admin"
    },
    {
        value: "professor",
        label: "Professor"
    },
    {
        value: "student",
        label: "Student"
    }
];