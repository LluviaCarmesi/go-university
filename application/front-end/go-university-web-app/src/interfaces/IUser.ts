export default interface IUser {
    Email: string;
    EmailAlias: string;
    Password: string;
    FirstName: string;
    LastName: string;
    PhoneNumber: string;
    HomeAddress: string;
    Role: string;
    Token: string;
    MustChangePW: boolean;
}