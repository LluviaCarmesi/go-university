import isStatusGood from "../../utilities/isStatusGood";
import * as strings from "../../strings/ENUSStrings";
import { PROFESSORS_IN_DEPARTMENTS_API_URI } from "../../appSettings";
import type IProfessorInDepartment from "../../interfaces/IProfessorToDepartment";

export default async function getProfessorInDepartmentByID(id: string) {
    const returnedResponse: {
        professorInDepartment: IProfessorInDepartment,
        isSuccessful: boolean,
        errorMessage: string
    } = {
        professorInDepartment: {
            ProfessorEmail: "",
            DepartmentID: 0,
            DepartmentName: "",
            IsLeader: false
        },
        isSuccessful: false,
        errorMessage: ""
    }

    await fetch(`${PROFESSORS_IN_DEPARTMENTS_API_URI}${id}`)
        .then((response) => {
            returnedResponse.isSuccessful = isStatusGood(response.status);
            return response.json();
        }).
        then((result) => {
            if (!returnedResponse.isSuccessful) {
                returnedResponse.errorMessage = result.ErrorMessage;
            }
            else {
                returnedResponse.professorInDepartment = result;
            }
        })
        .catch((error) => {
            returnedResponse.errorMessage = error;
            console.log(error);
        });
    if (!returnedResponse.professorInDepartment.ProfessorEmail) {
        returnedResponse.errorMessage = strings.PROFESSOR_IN_DEPARTMENT_DOESNT_EXIST_ERROR_MESSAGE;
    }
    return returnedResponse;
}