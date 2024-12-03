import isStatusGood from "../../utilities/isStatusGood";
import * as strings from "../../strings/ENUSStrings";
import { APPOINTMENTS_API_URI } from "../../appSettings";
import type IAppointment from "../../interfaces/IAppointment";

export default async function getAppointments() {
    const returnedResponse: {
        appointments: IAppointment[],
        isSuccessful: boolean,
        errorMessage: string
    } = {
        appointments: [],
        isSuccessful: false,
        errorMessage: ""
    }

    await fetch(APPOINTMENTS_API_URI)
        .then((response) => {
            returnedResponse.isSuccessful = isStatusGood(response.status);
            return response.json();
        }).
        then((result) => {
            if (!returnedResponse.isSuccessful) {
                returnedResponse.errorMessage = result.ErrorMessage;
            }
            else {
                for (let i = 0; i < result.length; i++) {
                    const currentAppointment = result[i];
                    returnedResponse.appointments.push({
                        ID: currentAppointment.ID,
                        StudentEmail: currentAppointment.StudentEmail,
                        AdminEmail: currentAppointment.AdminEmail,
                        IsComplete: currentAppointment.IsComplete,
                        StartTime: new Date(currentAppointment.StartTime),
                        EndTime: new Date(currentAppointment.EndTime),
                    })
                }
            }
        })
        .catch((error) => {
            returnedResponse.errorMessage = error;
            console.log(error);
        });
    if (returnedResponse.appointments.length == 0) {
        returnedResponse.errorMessage = strings.NO_APPOINTMENTS_ERROR_MESSAGE;
    }
    return returnedResponse;
}