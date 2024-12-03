import isStatusGood from "../../utilities/isStatusGood";
import * as strings from "../../strings/ENUSStrings";
import type IAppointment from "../../interfaces/IAppointment";
import { APPOINTMENTS_API_URI } from "../../appSettings";

export default async function getAppointmentByID(id: string) {
    const returnedResponse: {
        appointment: IAppointment,
        isSuccessful: boolean,
        errorMessage: string
    } = {
        appointment: {
            ID: 0,
            StudentEmail: "",
            AdminEmail: "",
            IsComplete: false,
            StartTime: new Date(),
            EndTime: new Date()
        },
        isSuccessful: false,
        errorMessage: ""
    }

    await fetch(`${APPOINTMENTS_API_URI}${id}`)
        .then((response) => {
            returnedResponse.isSuccessful = isStatusGood(response.status);
            return response.json();
        }).
        then((result) => {
            if (!returnedResponse.isSuccessful) {
                returnedResponse.errorMessage = result.ErrorMessage;
            }
            else {
                returnedResponse.appointment = result;
                returnedResponse.appointment.StartTime = new Date(result.StartTime);
                returnedResponse.appointment.EndTime = new Date(result.EndTime);
            }
        })
        .catch((error) => {
            returnedResponse.errorMessage = error;
            console.log(error);
        });
    if (!returnedResponse.appointment.ID) {
        returnedResponse.errorMessage = strings.APPOINTMENT_DOESNT_EXIST_ERROR_MESSAGE;
    }
    return returnedResponse;
}