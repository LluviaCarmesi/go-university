import isStatusGood from "../../utilities/isStatusGood";
import * as strings from "../../strings/ENUSStrings";
import { USERS_API_URI } from "../../appSettings";
import type IUser from "../../interfaces/IUser";

export default async function getUsers() {
    const returnedResponse: {
        users: IUser[],
        isSuccessful: boolean,
        errorMessage: string
    } = {
        users: [],
        isSuccessful: false,
        errorMessage: ""
    }

    await fetch(USERS_API_URI)
        .then((response) => {
            returnedResponse.isSuccessful = isStatusGood(response.status);
            return response.json();
        }).
        then((result) => {
            if (!returnedResponse.isSuccessful) {
                returnedResponse.errorMessage = result.ErrorMessage;
            }
            else {
                returnedResponse.users = result;
            }
        })
        .catch((error) => {
            returnedResponse.errorMessage = error;
            console.log(error);
        });
    if (returnedResponse.users.length == 0) {
        returnedResponse.errorMessage = strings.NO_APPOINTMENTS_ERROR_MESSAGE;
    }
    return returnedResponse;
}