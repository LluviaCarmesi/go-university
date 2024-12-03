import isStatusGood from "../../utilities/isStatusGood";
import * as strings from "../../strings/ENUSStrings";
import { DEPARTMENTS_API_URI } from "../../appSettings";
import type IDepartment from "../../interfaces/IDepartment";

export default async function getDepartments() {
    const returnedResponse: {
        departments: IDepartment[],
        isSuccessful: boolean,
        errorMessage: string
    } = {
        departments: [],
        isSuccessful: false,
        errorMessage: ""
    }

    await fetch(DEPARTMENTS_API_URI)
        .then((response) => {
            returnedResponse.isSuccessful = isStatusGood(response.status);
            return response.json();
        }).
        then((result) => {
            if (!returnedResponse.isSuccessful) {
                returnedResponse.errorMessage = result.ErrorMessage;
            }
            else {
                returnedResponse.departments = result;
            }
        })
        .catch((error) => {
            returnedResponse.errorMessage = error;
            console.log(error);
        });
    if (returnedResponse.departments.length == 0) {
        returnedResponse.errorMessage = strings.NO_DEPARTMENTS_ERROR_MESSAGE;
    }
    return returnedResponse;
}