import { APPOINTMENTS_API_URI } from "../../appSettings";
import type IUser from "../../interfaces/IUser";
import isStatusGood from "../../utilities/isStatusGood";

export default async function addAppointment(item: IUser) {
    let doesErrorExist = false;
    let errorMessage = "";
    await fetch(APPOINTMENTS_API_URI, {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify({
            Email: item.Email,
            EmailAlias: item.EmailAlias,
            Password: item.Password,
            FirstName: item.FirstName,
            LastName: item.LastName,
            PhoneNumber: item.PhoneNumber,
            HomeAddress: item.HomeAddress,
            Role: item.Role,
            Token: item.Token,
            MustChangePW: item.MustChangePW
        })
    })
        .then((response) => {
            doesErrorExist = !isStatusGood(response.status);
            return response.json();
        })
        .then((result) => {
            if (doesErrorExist) {
                errorMessage = result.response;
            }
        })
        .catch((error) => {
            doesErrorExist = true;
            errorMessage = error;
            console.log(error);
        });
    return { doesErrorExist, errorMessage };
}