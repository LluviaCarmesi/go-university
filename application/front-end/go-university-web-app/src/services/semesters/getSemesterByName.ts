import isStatusGood from "../../utilities/isStatusGood";
import * as strings from "../../strings/ENUSStrings";
import type ISemester from "../../interfaces/ISemester";
import { SEMESTERS_API_URI } from "../../appSettings";

export default async function getSemesterByName(id: string) {
    const returnedResponse: {
        semester: ISemester,
        isSuccessful: boolean,
        errorMessage: string
    } = {
        semester: {
            Name: "",
            StartDate: new Date(),
            EndDate: new Date()
        },
        isSuccessful: false,
        errorMessage: ""
    }

    await fetch(`${SEMESTERS_API_URI}${id}`)
        .then((response) => {
            returnedResponse.isSuccessful = isStatusGood(response.status);
            return response.json();
        }).
        then((result) => {
            if (!returnedResponse.isSuccessful) {
                returnedResponse.errorMessage = result.response;
            }
            else {
                returnedResponse.semester.Name = result.Name;
                returnedResponse.semester.StartDate = new Date(result.StartDate);
                returnedResponse.semester.EndDate = new Date(result.EndDate);
            }
        })
        .catch((error) => {
            returnedResponse.errorMessage = error;
            console.log(error);
        });
    if (!returnedResponse.semester.Name) {
        returnedResponse.errorMessage = strings.SEMESTER_DOESNT_EXIST_ERROR_MESSAGE;
    }
    return returnedResponse;
}