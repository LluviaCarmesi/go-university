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
    {#if role.isAdmin || role.isStudent}
        <div>
            <a href="appointments/add">Add Appointment</a>
        </div>
        <div>
            <a href="appointments/show">Edit Appointment</a>
        </div>
    {/if}
</div>
