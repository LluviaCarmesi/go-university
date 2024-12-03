import isStatusGood from "../../utilities/isStatusGood";
import * as strings from "../../strings/ENUSStrings";
import { PROFESSORS_IN_DEPARTMENTS_API_URI } from "../../appSettings";
import type IProfessorInDepartment from "../../interfaces/IProfessorToDepartment";

export default async function getProfessorsInDepartments() {
    const returnedResponse: {
        professorsInDepartments: IProfessorInDepartment[],
        isSuccessful: boolean,
        errorMessage: string
    } = {
        professorsInDepartments: [],
        isSuccessful: false,
        errorMessage: ""
    }

    await fetch(PROFESSORS_IN_DEPARTMENTS_API_URI)
        .then((response) => {
            returnedResponse.isSuccessful = isStatusGood(response.status);
            return response.json();
        }).
        then((result) => {
            if (!returnedResponse.isSuccessful) {
                returnedResponse.errorMessage = result.ErrorMessage;
            }
            else {
                returnedResponse.professorsInDepartments = result;
            }
        })
        .catch((error) => {
            returnedResponse.errorMessage = error;
            console.log(error);
        });
    if (returnedResponse.professorsInDepartments.length == 0) {
        returnedResponse.errorMessage = strings.NO_PROFESSORS_IN_DEPARTMENTS_ERROR_MESSAGE;
    }
    return returnedResponse;
}