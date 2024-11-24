import isStatusGood from "../../utilities/isStatusGood";
import * as strings from "../../strings/ENUSStrings";
import { DEPARTMENTS_API_URI } from "../../appSettings";
import type IDepartment from "../../interfaces/IDepartment";

export default async function getDepartmentByID(id: string) {
    const returnedResponse: {
        department: IDepartment,
        isSuccessful: boolean,
        errorMessage: string
    } = {
        department: {
            ID: 0,
            Name: "",
        },
        isSuccessful: false,
        errorMessage: ""
    }

    await fetch(`${DEPARTMENTS_API_URI}${id}`)
        .then((response) => {
            returnedResponse.isSuccessful = isStatusGood(response.status);
            return response.json();
        }).
        then((result) => {
            if (!returnedResponse.isSuccessful) {
                returnedResponse.errorMessage = result.response;
            }
            else {
                returnedResponse.department = result;
            }
        })
        .catch((error) => {
            returnedResponse.errorMessage = error;
            console.log(error);
        });
    if (!returnedResponse.department.ID) {
        returnedResponse.errorMessage = strings.DEPARTMENT_DOESNT_EXIST_ERROR_MESSAGE;
    }
    return returnedResponse;
}