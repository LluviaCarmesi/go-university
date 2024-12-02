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
            StartTime: new Date(),
            EndTime: new Date()
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
                returnedResponse.taughtCourse.ID = result.ID;
                returnedResponse.taughtCourse.CourseID = result.CourseID;
                returnedResponse.taughtCourse.SemesterName = result.SemesterName;
                returnedResponse.taughtCourse.ProfessorEmail = result.ProfessorEmail;
                returnedResponse.taughtCourse.Location = result.Location;
                returnedResponse.taughtCourse.MaxStudents = result.MaxStudents;
                returnedResponse.taughtCourse.StartTime = new Date(result.StartTime);
                returnedResponse.taughtCourse.EndTime = new Date(result.EndTime);
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