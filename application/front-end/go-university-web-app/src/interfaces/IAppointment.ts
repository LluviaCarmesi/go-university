export default interface IAppointment {
    ID: number;
    StudentEmail: string;
    AdminEmail: string;
    IsComplete: boolean;
    StartTime: Date;
    EndTime: Date;
    Description: string;
}