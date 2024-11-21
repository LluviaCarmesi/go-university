import isStatusGood from "../../utilities/isStatusGood";
import * as strings from "../../strings/ENUSStrings";
import type IUser from "../../interfaces/IUser";
import { USERS_BY_TOKEN_API_URI } from "../../appSettings";

export default async function getUserByToken(token: string) {
    const returnedResponse: {
        user: IUser,
        isSuccessful: boolean,
        errorMessage: string
    } = {
        user: {
            Email: "",
            EmailAlias: "",
            Password: "",
            FirstName: "",
            LastName: "",
            PhoneNumber: "",
            HomeAddress: "",
            Role: "",
            Token: "",
            MustChangePW: false
        },
        isSuccessful: false,
        errorMessage: ""
    }

    await fetch(`${USERS_BY_TOKEN_API_URI}${token}`)
        .then((response) => {
            returnedResponse.isSuccessful = isStatusGood(response.status);
            return response.json();
        }).
        then((result) => {
            if (!returnedResponse.isSuccessful) {
                returnedResponse.errorMessage = result.response;
            }
            else {
                returnedResponse.user = result;
            }
        })
        .catch((error) => {
            returnedResponse.errorMessage = error;
            console.log(error);
        });
    if (!returnedResponse.user.Email) {
        returnedResponse.errorMessage = strings.USER_DOESNT_EXIST_ERROR_MESSAGE;
    }
    return returnedResponse;
}