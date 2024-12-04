<script lang="ts">
    import Navigation from "../../../components/Navigation.svelte";
    import type IRole from "../../../interfaces/IRole";
    import "../../../styles/items.css";
    import "../../../styles/common.css";
    import type ITaughtCourse from "../../../interfaces/ITaughtCourse";
    import getTaughtCourses from "../../../services/taught-courses/getTaughtCourses";
    import formatTime from "../../../utilities/formatTime";
    import addRegistration from "../../../services/registrations/addRegistration";
    import getCookie from "../../../utilities/getCookie";
    import checkCurrentUser from "../../../services/users/checkCurrentUser";
    import { onMount } from "svelte";

    let role: IRole = {
        isAdmin: false,
        isProfessor: false,
        isStudent: false,
    };

    let isLoading = false;
    let errorMessage = "";
    let isSuccessful = false;

    onMount(() => {
        async function getRole() {
            const checkCurrentUserResponse = await checkCurrentUser();
            role = checkCurrentUserResponse;
        }
        getRole();
    });

    let taughtCourses: ITaughtCourse[] = [];

    async function submitRegistration(taughtCourseID: number) {
        isLoading = true;
        const addRegistrationResponse = await addRegistration({
            TaughtCourseID: taughtCourseID,
            CourseID: "",
            StudentEmail: getCookie("email"),
            FinalGrade: 0,
            Status: "",
        });
        isSuccessful = !addRegistrationResponse.doesErrorExist;
        if (!isSuccessful) {
            errorMessage = addRegistrationResponse.errorMessage;
        } else {
            errorMessage = "";
            window.open("/registrations/add", "_self");
        }
        isLoading = false;
    }

    async function getTaughtCoursesResponse() {
        const getTaughtCoursesResponse = await getTaughtCourses();
        if (getTaughtCoursesResponse.isSuccessful) {
            taughtCourses = getTaughtCoursesResponse.taughtCourses;
        } else {
            errorMessage = getTaughtCoursesResponse.errorMessage;
        }
        isLoading = false;
    }
    getTaughtCoursesResponse();
</script>

<Navigation {role} />

<div id="itemsTable">
    {#if taughtCourses.length === 0}
        <div class="descriptionContainer">
            <span>No taught courses as of now!</span>
        </div>
    {:else}
        <div class="mainHeadingContainer">
            <h2>Register for a Course</h2>
        </div>
        <div class="descriptionContainer">
            <span>Add a course to your registration!</span>
        </div>
        {#if !!errorMessage}
            <div class="errorContainer">
                <span>{errorMessage}</span>
            </div>
        {/if}
        {#if isSuccessful}
            <div class="successContainer">
                <span>Course was added successfully!</span>
            </div>
        {/if}
        <table>
            <thead>
                <tr>
                    <th>Register</th>
                    <th>Course ID</th>
                    <th>Semester Name</th>
                    <th>Professor Email</th>
                    <th>Location</th>
                    <th>Max Students</th>
                    <th>Day</th>
                    <th>Start Time</th>
                    <th>End Time</th>
                </tr>
            </thead>
            {#each taughtCourses as taughtCourse}
                <tbody>
                    <tr>
                        <td>
                            <button
                                on:click={() =>
                                    submitRegistration(taughtCourse.ID)}
                                >Register</button
                            >
                        </td>
                        <td>
                            <span>{taughtCourse.CourseID}</span>
                        </td>
                        <td>
                            <span>{taughtCourse.SemesterName}</span>
                        </td>
                        <td>
                            <span>{taughtCourse.ProfessorEmail}</span>
                        </td>
                        <td>
                            <span>{taughtCourse.Location}</span>
                        </td>
                        <td>
                            <span>{taughtCourse.MaxStudents}</span>
                        </td>
                        <td>
                            <span>{taughtCourse.Day}</span>
                        </td>
                        <td>
                            <span>{formatTime(taughtCourse.StartTime)}</span>
                        </td>
                        <td>
                            <span>{formatTime(taughtCourse.EndTime)}</span>
                        </td>
                    </tr>
                </tbody>
            {/each}
        </table>
    {/if}
</div>
