import * as SETTINGS from "../appSettings";
import type ICourse from "../interfaces/ICourse";
import isStatusGood from "../utilities/isStatusGood";

export default async function addCourse(item: ICourse) {
    let doesErrorExist = false;
    let errorMessage = "";
    await fetch(`${SETTINGS.COURSES_API_URI}${SETTINGS.ADD_COURSE_URI}`, {
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