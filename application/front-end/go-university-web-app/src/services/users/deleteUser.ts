import { USERS_API_URI } from "../../appSettings";
import type IUser from "../../interfaces/IUser";
import isStatusGood from "../../utilities/isStatusGood";

export default async function deleteAppointment(item: IUser) {
    let doesErrorExist = false;
    let errorMessage = "";
    await fetch(`${USERS_API_URI}${item.Email}`, {
        method: "DELETE",
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
                errorMessage = result.ErrorMessage;
            }
        })
        .catch((error) => {
            doesErrorExist = true;
            errorMessage = error;
            console.log(error);
        });
    return { doesErrorExist, errorMessage };
}