<script lang="ts">
    import Navigation from "../../components/Navigation.svelte";
    import type IRole from "../../interfaces/IRole";
    import "../../styles/pages/courses/courses.css";

    let role: IRole = {
        isAdmin: true,
        isProfessor: false,
        isStudent: false,
    };
</script>

<Navigation {role} />

<div id="coursesLinks">
    {#if role.isAdmin}
        <div>
            <a href="courses/add">Add Course</a>
        </div>
        <div>
            <a href="courses/show-courses">Edit Courses</a>
        </div>
    {:else if role.isProfessor}
        <div>
            <a href="courses/stats">Stats</a>
        </div>
    {:else if role.isStudent || role.isProfessor}
        <div>
            <a href="courses/grades">Grades</a>
        </div>
        <div>
            <a href="courses/schedule">Schedule</a>
        </div>
    {/if}
</div>
