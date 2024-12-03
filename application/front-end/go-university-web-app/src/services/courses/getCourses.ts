import isStatusGood from "../../utilities/isStatusGood";
import * as strings from "../../strings/ENUSStrings";
import type ICourse from "../../interfaces/ICourse";
import { COURSES_API_URI } from "../../appSettings";

export default async function getCourses() {
    const returnedResponse: {
        courses: ICourse[],
        isSuccessful: boolean,
        errorMessage: string
    } = {
        courses: [],
        isSuccessful: false,
        errorMessage: ""
    }

    await fetch(COURSES_API_URI)
        .then((response) => {
            returnedResponse.isSuccessful = isStatusGood(response.status);
            return response.json();
        }).
        then((result) => {
            if (!returnedResponse.isSuccessful) {
                returnedResponse.errorMessage = result.ErrorMessage;
            }
            else {
                returnedResponse.courses = result;
            }
        })
        .catch((error) => {
            returnedResponse.errorMessage = error;
            console.log(error);
        });
    if (returnedResponse.courses.length == 0) {
        returnedResponse.errorMessage = strings.NO_COURSES_ERROR_MESSAGE;
    }
    return returnedResponse;
}