import { APPOINTMENTS_API_URI } from "../../appSettings";
import type IAppointment from "../../interfaces/IAppointment";
import isStatusGood from "../../utilities/isStatusGood";

export default async function editAppointment(item: IAppointment) {
    let doesErrorExist = false;
    let errorMessage = "";
    await fetch(`${APPOINTMENTS_API_URI}${item.ID}`, {
        method: "PATCH",
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