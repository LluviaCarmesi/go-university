import { TAUGHT_COURSES_API_URI } from "../../appSettings";
import type ITaughtCourse from "../../interfaces/ITaughtCourse";
import isStatusGood from "../../utilities/isStatusGood";

export default async function editTaughtCourse(item: ITaughtCourse) {
    let doesErrorExist = false;
    let errorMessage = "";
    await fetch(`${TAUGHT_COURSES_API_URI}${item.ID}`, {
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