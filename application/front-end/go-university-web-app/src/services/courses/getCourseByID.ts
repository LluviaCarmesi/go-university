import isStatusGood from "../../utilities/isStatusGood";
import * as strings from "../../strings/ENUSStrings";
import type ICourse from "../../interfaces/ICourse";
import { COURSES_API_URI } from "../../appSettings";

export default async function getCourseByID(id: string) {
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

    await fetch(`${COURSES_API_URI}${id}`)
        .then((response) => {
            returnedResponse.isSuccessful = isStatusGood(response.status);
            return response.json();
        }).
        then((result) => {
            if (!returnedResponse.isSuccessful) {
                returnedResponse.errorMessage = result.ErrorMessage;
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