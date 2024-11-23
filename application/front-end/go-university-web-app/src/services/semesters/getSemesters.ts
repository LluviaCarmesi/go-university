import isStatusGood from "../../utilities/isStatusGood";
import * as strings from "../../strings/ENUSStrings";
import { SEMESTERS_API_URI } from "../../appSettings";
import type ISemester from "../../interfaces/ISemester";

export default async function getSemesters() {
    const returnedResponse: {
        semesters: ISemester[],
        isSuccessful: boolean,
        errorMessage: string
    } = {
        semesters: [],
        isSuccessful: false,
        errorMessage: ""
    }

    await fetch(SEMESTERS_API_URI)
        .then((response) => {
            returnedResponse.isSuccessful = isStatusGood(response.status);
            return response.json();
        }).
        then((result) => {
            if (!returnedResponse.isSuccessful) {
                returnedResponse.errorMessage = result.response;
            }
            else {
                for (let i = 0; i < result.length; i++) {
                    const currentSemester = result[i];
                    returnedResponse.semesters.push({
                        Name: currentSemester.Name,
                        StartDate: new Date(currentSemester.StartDate),
                        EndDate: new Date(currentSemester.EndDate),
                    })
                }
            }
        })
        .catch((error) => {
            returnedResponse.errorMessage = error;
            console.log(error);
        });
    if (returnedResponse.semesters.length == 0) {
        returnedResponse.errorMessage = strings.NO_SEMESTERS_ERROR_MESSAGE;
    }
    return returnedResponse;
}