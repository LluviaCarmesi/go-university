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
            <a href="departments/add-departments">Add Department</a>
        </div>
        <div>
            <a href="departments/show-departments">Edit Departments</a>
        </div>
        <div>
            <a href="departments/add-professors">Add Professor Department</a>
        </div>
        <div>
            <a href="departments/show-professors">Edit Professors Departments</a>
        </div>
    {/if}
</div>
