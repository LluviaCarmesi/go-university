<script lang="ts">
    import { onMount } from "svelte";
    import Navigation from "../../../components/Navigation.svelte";
    import type ICourse from "../../../interfaces/ICourse";
    import type IRole from "../../../interfaces/IRole";
    import getCourses from "../../../services/courses/getCourses";
    import checkCurrentUser from "../../../services/users/checkCurrentUser";
    import "../../../styles/items.css";
    import deleteCourse from "../../../services/courses/deleteCourse";

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

    let courses: ICourse[] = [];

    async function removeCourse(course: ICourse) {
        isLoading = true;
        const deleteCourseResponse = await deleteCourse(course);
        isSuccessful = !deleteCourseResponse.doesErrorExist;
        if (!isSuccessful) {
            errorMessage = deleteCourseResponse.errorMessage;
        } else {
            errorMessage = "";
            window.open("show", "_self");
        }
        isLoading = false;
    }

    async function getCoursesResponse() {
        const coursesResponse = await getCourses();
        if (coursesResponse.isSuccessful) {
            courses = coursesResponse.courses;
        } else {
            errorMessage = coursesResponse.errorMessage;
        }
        isLoading = false;
    }
    getCoursesResponse();
</script>

<Navigation {role} />

<div id="itemsTable">
    {#if courses.length === 0}
        <div class="descriptionContainer">
            <span>No courses as of now!</span>
        </div>
    {:else}
        <table>
            <thead>
                <tr>
                    <th>ID</th>
                    <th>Name</th>
                    <th>Description</th>
                    <th>Credits</th>
                    <th>Remove</th>
                </tr>
            </thead>
            {#each courses as course}
                <tbody>
                    <tr>
                        <td>
                            <a href={"edit/" + course.ID}>
                                <span>{course.ID}</span>
                            </a>
                        </td>
                        <td>
                            <span>{course.Name}</span>
                        </td>
                        <td>
                            <span>{course.Description}</span>
                        </td>
                        <td>
                            <span>{course.Credits}</span>
                        </td>
                        <td>
                            <button on:click={() => removeCourse(course)}
                                >Remove</button
                            >
                        </td>
                    </tr>
                </tbody>
            {/each}
        </table>
    {/if}
</div>
