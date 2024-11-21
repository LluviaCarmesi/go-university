import { APPOINTMENTS_API_URI } from "../../appSettings";
import type IAppointment from "../../interfaces/IAppointment";
import isStatusGood from "../../utilities/isStatusGood";

export default async function addAppointment(item: IAppointment) {
    let doesErrorExist = false;
    let errorMessage = "";
    await fetch(APPOINTMENTS_API_URI, {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify({
            ID: item.ID,
            StudentEmail: item.StudentEmail,
            AdminEmail: item.AdminEmail,
            IsComplete: item.IsComplete,
            StartTime: item.StartTime,
            EndTime: item.EndTime
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