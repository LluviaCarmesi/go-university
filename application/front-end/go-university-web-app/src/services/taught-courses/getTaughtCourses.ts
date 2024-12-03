import isStatusGood from "../../utilities/isStatusGood";
import * as strings from "../../strings/ENUSStrings";
import { TAUGHT_COURSES_API_URI } from "../../appSettings";
import type ITaughtCourse from "../../interfaces/ITaughtCourse";

export default async function getTaughtCourses() {
    const returnedResponse: {
        taughtCourses: ITaughtCourse[],
        isSuccessful: boolean,
        errorMessage: string
    } = {
        taughtCourses: [],
        isSuccessful: false,
        errorMessage: ""
    }

    await fetch(TAUGHT_COURSES_API_URI)
        .then((response) => {
            returnedResponse.isSuccessful = isStatusGood(response.status);
            return response.json();
        }).
        then((result) => {
            if (!returnedResponse.isSuccessful) {
                returnedResponse.errorMessage = result.response;
            }
            else {
                returnedResponse.taughtCourses = result;
            }
        })
        .catch((error) => {
            returnedResponse.errorMessage = error;
            console.log(error);
        });
    if (returnedResponse.taughtCourses.length == 0) {
        returnedResponse.errorMessage = strings.NO_TAUGHT_COURSES_ERROR_MESSAGE;
    }
    return returnedResponse;
}