import isStatusGood from "../../utilities/isStatusGood";
import * as strings from "../../strings/ENUSStrings";
import { REGISTRATIONS_API_URI } from "../../appSettings";
import type IRegistration from "../../interfaces/IRegistration";

export default async function getRegistrations() {
    const returnedResponse: {
        registrations: IRegistration[],
        isSuccessful: boolean,
        errorMessage: string
    } = {
        registrations: [],
        isSuccessful: false,
        errorMessage: ""
    }

    await fetch(REGISTRATIONS_API_URI)
        .then((response) => {
            returnedResponse.isSuccessful = isStatusGood(response.status);
            return response.json();
        }).
        then((result) => {
            if (!returnedResponse.isSuccessful) {
                returnedResponse.errorMessage = result.ErrorMessage;
            }
            else {
                returnedResponse.registrations = result;
            }
        })
        .catch((error) => {
            returnedResponse.errorMessage = error;
            console.log(error);
        });
    if (returnedResponse.registrations.length == 0) {
        returnedResponse.errorMessage = strings.NO_REGISTRATIONS_ERROR_MESSAGE;
    }
    return returnedResponse;
}