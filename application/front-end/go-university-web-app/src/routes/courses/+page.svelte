<script lang="ts">
    import Navigation from "../../components/Navigation.svelte";
    import type IRole from "../../interfaces/IRole";

    let role: IRole = {
        isAdmin: true,
        isProfessor: false,
        isStudent: false,
    };
</script>

<Navigation {role} />

<div>
    {#if role.isAdmin}
        <a href="courses/add">Add Course</a>
        <a href="courses/edit">Edit Courses</a>
    {:else if role.isProfessor}
        <a href="courses/stats">Stats</a>
    {:else if role.isStudent || role.isProfessor}
        <a href="courses/grades">Grades</a>
        <a href="courses/schedule">Schedule</a>
    {/if}
</div>
