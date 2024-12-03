import { COURSES_API_URI } from "../../appSettings";
import type ICourse from "../../interfaces/ICourse";
import isStatusGood from "../../utilities/isStatusGood";

export default async function editCourse(item: ICourse) {
    let doesErrorExist = false;
    let errorMessage = "";
    await fetch(`${COURSES_API_URI}${item.ID}`, {
        method: "PATCH",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify({
            ID: item.ID,
            Name: item.Name,
            Description: item.Description,
            Credits: parseInt(item.Credits)
        })
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