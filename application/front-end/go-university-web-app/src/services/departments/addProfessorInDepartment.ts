import { PROFESSORS_IN_DEPARTMENTS_API_URI } from "../../appSettings";
import type IProfessorInDepartment from "../../interfaces/IProfessorToDepartment";
import isStatusGood from "../../utilities/isStatusGood";

export default async function addProfessorInDepartment(item: IProfessorInDepartment) {
    let doesErrorExist = false;
    let errorMessage = "";
    await fetch(PROFESSORS_IN_DEPARTMENTS_API_URI, {
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