import { USERS_LOGIN_API_URI } from "../../appSettings";
import type ILogin from "../../interfaces/ILogin";
import isStatusGood from "../../utilities/isStatusGood";
import setCookie from "../../utilities/setCookie";

export default async function loginUser(item: ILogin) {
    const returnedResponse: {
        token: string,
        isSuccessful: boolean,
        errorMessage: string
    } = {
        token: "",
        isSuccessful: false,
        errorMessage: ""
    }
    await fetch(`${USERS_LOGIN_API_URI}${item.Email}`, {
        method: "PATCH",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify({
            Email: "",
            EmailAlias: "",
            Password: item.Password,
            FirstName: "",
            LastName: "",
            PhoneNumber: "",
            HomeAddress: "",
            Role: "",
            Token: "",
            MustChangePW: false,
        })
    })
        .then((response) => {
            returnedResponse.isSuccessful = isStatusGood(response.status);
            return response.json();
        })
        .then((result) => {
            if (!returnedResponse.isSuccessful) {
                returnedResponse.errorMessage = result.response;
            }
            else {
                setCookie("token", result.Token, 30);
                setCookie("email", item.Email, 30);
            }
        })
        .catch((error) => {
            returnedResponse.errorMessage = error;
            console.log(error);
        });
    return returnedResponse;
}