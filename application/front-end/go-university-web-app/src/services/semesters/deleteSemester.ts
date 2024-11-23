import { SEMESTERS_API_URI } from "../../appSettings";
import type ISemester from "../../interfaces/ISemester";
import isStatusGood from "../../utilities/isStatusGood";

export default async function deleteSemester(item: ISemester) {
    let doesErrorExist = false;
    let errorMessage = "";
    await fetch(`${SEMESTERS_API_URI}${item.Name}`, {
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