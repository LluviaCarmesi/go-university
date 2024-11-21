<script lang="ts">
    import { onMount } from "svelte";
    import Navigation from "../components/Navigation.svelte";
    import type IRole from "../interfaces/IRole";
    import checkCurrentUser from "../services/users/checkCurrentUser";
    import "../styles/pages/home.css";

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

<div id="homePageContainer">
    <p>Welcome to Go University!</p>
    {#if role.isAdmin}
        <p>
            You are an admin! Use the menus above to change anything in the
            system.
        </p>
    {:else if role.isProfessor}
        <p>
            You are a professor! Use the menus above to accept pending requests
            or view stats on your courses.
        </p>
    {:else if role.isStudent}
        <p>You are student! Use the menus above to register for classes.</p>
    {/if}
</div>
