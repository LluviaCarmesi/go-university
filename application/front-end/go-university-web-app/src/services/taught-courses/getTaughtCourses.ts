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
                for (let i = 0; i < result.length; i++) {
                    const currentTaughtCourse = result[i];
                    returnedResponse.taughtCourses.push({
                        ID: currentTaughtCourse.Name,
                        CourseID:currentTaughtCourse.CourseID,
                        SemesterName:currentTaughtCourse.SemesterName,
                        ProfessorEmail:currentTaughtCourse.ProfessorEmail,
                        Location:currentTaughtCourse.Location,
                        MaxStudents: currentTaughtCourse.MaxStudents,
                        StartTime: new Date(currentTaughtCourse.StartTime),
                        EndTime: new Date(currentTaughtCourse.EndTime),
                    })
                }
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