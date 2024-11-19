import * as SETTINGS from "../../appSettings";
import isStatusGood from "../../utilities/isStatusGood";
import * as strings from "../../strings/ENUSStrings";
import type ICourse from "../../interfaces/ICourse";

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

    await fetch(SETTINGS.COURSES_API_URI)
        .then((response) => {
            returnedResponse.isSuccessful = isStatusGood(response.status);
            return response.json();
        }).
        then((result) => {
            if (!returnedResponse.isSuccessful) {
                returnedResponse.errorMessage = result.response;
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