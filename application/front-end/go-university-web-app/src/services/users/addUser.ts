import { APPOINTMENTS_API_URI, USERS_API_URI } from "../../appSettings";
import type IUser from "../../interfaces/IUser";
import isStatusGood from "../../utilities/isStatusGood";

export default async function addAppointment(item: IUser) {
    let doesErrorExist = false;
    let errorMessage = "";
    await fetch(USERS_API_URI, {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify(item)
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