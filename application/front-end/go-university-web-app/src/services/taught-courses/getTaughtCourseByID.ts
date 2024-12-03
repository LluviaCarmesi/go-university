import isStatusGood from "../../utilities/isStatusGood";
import * as strings from "../../strings/ENUSStrings";
import { TAUGHT_COURSES_API_URI } from "../../appSettings";
import type ITaughtCourse from "../../interfaces/ITaughtCourse";

export default async function getTaughtCourseByID(id: string) {
    const returnedResponse: {
        taughtCourse: ITaughtCourse,
        isSuccessful: boolean,
        errorMessage: string
    } = {
        taughtCourse: {
            ID: 0,
            CourseID: "",
            SemesterName: "",
            ProfessorEmail: "",
            Location: "",
            MaxStudents: 0,
            Day: "",
            StartTime: "",
            EndTime: ""
        },
        isSuccessful: false,
        errorMessage: ""
    }

    await fetch(`${TAUGHT_COURSES_API_URI}${id}`)
        .then((response) => {
            returnedResponse.isSuccessful = isStatusGood(response.status);
            return response.json();
        }).
        then((result) => {
            if (!returnedResponse.isSuccessful) {
                returnedResponse.errorMessage = result.response;
            }
            else {
                returnedResponse.taughtCourse = result;
            }
        })
        .catch((error) => {
            returnedResponse.errorMessage = error;
            console.log(error);
        });
    if (!returnedResponse.taughtCourse.ID) {
        returnedResponse.errorMessage = strings.TAUGHT_COURSE_DOESNT_EXIST_ERROR_MESSAGE;
    }
    return returnedResponse;
}