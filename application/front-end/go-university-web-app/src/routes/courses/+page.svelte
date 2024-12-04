<script lang="ts">
    import Navigation from "../../components/Navigation.svelte";
    import type IRole from "../../interfaces/IRole";
    import "../../styles/items.css";
    import checkCurrentUser from "../../services/users/checkCurrentUser";
    import { onMount } from "svelte";

    let role: IRole = {
        isAdmin: false,
        isProfessor: false,
        isStudent: false,
    };
    onMount(() => {
        async function getRole() {
            const checkCurrentUserResponse = await checkCurrentUser();
            role = checkCurrentUserResponse;
        }
        getRole();
    });
</script>

<Navigation {role} />

<div id="itemsLinks">
    {#if role.isAdmin}
        <div>
            <a href="courses/add">Add Course</a>
        </div>
        <div>
            <a href="courses/show">Edit Courses</a>
        </div>
        <div>
            <a href="taught-courses/add">Add Taught Course</a>
        </div>
        <div>
            <a href="taught-courses/show">Edit Taught Course</a>
        </div>
        <div>
            <a href="registrations/show">Edit Registrations</a>
        </div>
    {:else if role.isProfessor}
        <div>
            <a href="courses/stats">Stats</a>
        </div>
    {:else if role.isStudent || role.isProfessor}
        <div>
            <a href="taught-courses/grades-status">Grades and Status</a>
        </div>
        <div>
            <a href="taught-courses/schedule">Schedule</a>
        </div>
    {:else if role.isStudent}
        <div>
            <a href="registrations/add">Register for classes</a>
        </div>
        <div>
            <a href="registrations/show">Edit Registrations</a>
        </div>
    {/if}
</div>
