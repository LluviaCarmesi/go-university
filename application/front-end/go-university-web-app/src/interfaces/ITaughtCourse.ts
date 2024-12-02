export default interface ITaughtCourse {
    ID: number;
	CourseID: string;
	SemesterName: string;
	ProfessorEmail: string;
	MaxStudents: number;
	Location: string;
	StartTime: Date;
	EndTime: Date;
}