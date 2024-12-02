import isStatusGood from "../../utilities/isStatusGood";
import * as strings from "../../strings/ENUSStrings";
import { REGISTRATIONS_API_URI } from "../../appSettings";
import type IRegistration from "../../interfaces/IRegistration";

export default async function getRegistrationByID(id: string) {
    const returnedResponse: {
        registration: IRegistration,
        isSuccessful: boolean,
        errorMessage: string
    } = {
        registration: {
            StudentEmail: "",
            TaughtCourseID: 0,
            CourseID: "",
            FinalGrade: 0,
            Status: ""
        },
        isSuccessful: false,
        errorMessage: ""
    }

    await fetch(`${REGISTRATIONS_API_URI}${id}`)
        .then((response) => {
            returnedResponse.isSuccessful = isStatusGood(response.status);
            return response.json();
        }).
        then((result) => {
            if (!returnedResponse.isSuccessful) {
                returnedResponse.errorMessage = result.response;
            }
            else {
                returnedResponse.registration = result;
            }
        })
        .catch((error) => {
            returnedResponse.errorMessage = error;
            console.log(error);
        });
    if (!returnedResponse.registration.StudentEmail) {
        returnedResponse.errorMessage = strings.REGISTRATION_DOESNT_EXIST_ERROR_MESSAGE;
    }
    return returnedResponse;
}