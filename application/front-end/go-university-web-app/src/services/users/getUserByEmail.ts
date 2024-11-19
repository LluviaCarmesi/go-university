import * as SETTINGS from "../../appSettings";
import isStatusGood from "../../utilities/isStatusGood";
import * as strings from "../../strings/ENUSStrings";
import type ICourse from "../../interfaces/ICourse";

export default async function getUserByEmail(email: string) {
    const returnedResponse: {
        course: ICourse,
        isSuccessful: boolean,
        errorMessage: string
    } = {
        course: {
            ID: "",
            Name: "",
            Description: "",
            Credits: ""
        },
        isSuccessful: false,
        errorMessage: ""
    }

    await fetch(`${SETTINGS.COURSES_API_URI}${id}`)
        .then((response) => {
            returnedResponse.isSuccessful = isStatusGood(response.status);
            return response.json();
        }).
        then((result) => {
            if (!returnedResponse.isSuccessful) {
                returnedResponse.errorMessage = result.response;
            }
            else {
                returnedResponse.course = result;
            }
        })
        .catch((error) => {
            returnedResponse.errorMessage = error;
            console.log(error);
        });
    if (!returnedResponse.course.ID) {
        returnedResponse.errorMessage = strings.COURSE_DOESNT_EXIST_ERROR_MESSAGE;
    }
    return returnedResponse;
}