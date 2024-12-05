<script lang="ts">
    import Navigation from "../../../components/Navigation.svelte";
    import type IRole from "../../../interfaces/IRole";
    import "../../../styles/items.css";
    import "../../../styles/common.css";
    import type ITaughtCourse from "../../../interfaces/ITaughtCourse";
    import getTaughtCourses from "../../../services/taught-courses/getTaughtCourses";
    import formatTime from "../../../utilities/formatTime";
    import checkCurrentUser from "../../../services/users/checkCurrentUser";
    import { onMount } from "svelte";
    import getCookie from "../../../utilities/getCookie";

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
            getTaughtCoursesResponse();
        }
        getRole();
    });

    let taughtCourses: ITaughtCourse[] = [];

    async function getTaughtCoursesResponse() {
        const getTaughtCoursesResponse = await getTaughtCourses();
        if (getTaughtCoursesResponse.isSuccessful) {
            taughtCourses = getTaughtCoursesResponse.taughtCourses.filter(
                (course) => course.ProfessorEmail === getCookie("email"),
            );
        } else {
            errorMessage = getTaughtCoursesResponse.errorMessage;
        }
        isLoading = false;
    }
</script>

<Navigation {role} />

<div id="itemsTable">
    {#if taughtCourses.length === 0}
        <div class="descriptionContainer">
            <span>No taught courses as of now!</span>
        </div>
    {:else}
        <table>
            <thead>
                <tr>
                    <th>ID</th>
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
                            <a href={"/registrations/course/" + taughtCourse.ProfessorEmail + "-" + taughtCourse.ID}>
                                <span>{taughtCourse.ID}</span>
                            </a>
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
