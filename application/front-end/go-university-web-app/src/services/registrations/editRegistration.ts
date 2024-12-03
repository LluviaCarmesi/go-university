import { REGISTRATIONS_API_URI } from "../../appSettings";
import type IRegistration from "../../interfaces/IRegistration";
import isStatusGood from "../../utilities/isStatusGood";

export default async function editRegistration(item: IRegistration) {
    let doesErrorExist = false;
    let errorMessage = "";
    await fetch(`${REGISTRATIONS_API_URI}${item.StudentEmail}-${item.TaughtCourseID}`, {
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