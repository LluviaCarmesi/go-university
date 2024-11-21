<script lang="ts">
    import Navigation from "../../../components/Navigation.svelte";
    import type ICourse from "../../../interfaces/ICourse";
    import type IRole from "../../../interfaces/IRole";
    import getCourses from "../../../services/courses/getCourses";
    import "../../../styles/items.css";

    let role: IRole = {
        isAdmin: true,
        isProfessor: false,
        isStudent: false,
    };
    let isLoading = false;
    let errorMessage = "";

    let courses: ICourse[] = [];

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
                    </tr>
                </tbody>
            {/each}
        </table>
    {/if}
</div>
