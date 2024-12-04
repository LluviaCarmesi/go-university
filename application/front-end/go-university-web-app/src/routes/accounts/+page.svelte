<script lang="ts">
    import Navigation from "../../components/Navigation.svelte";
    import type IRole from "../../interfaces/IRole";
    import "../../styles/items.css";
    import checkCurrentUser from "../../services/users/checkCurrentUser";
    import { onMount } from "svelte";
    import getCookie from "../../utilities/getCookie";

    let role: IRole = {
        isAdmin: false,
        isProfessor: false,
        isStudent: false,
    };
    let currentUserEmail = "";
    onMount(() => {
        async function getRole() {
            const checkCurrentUserResponse = await checkCurrentUser();
            role = checkCurrentUserResponse;
        }
        getRole();
        currentUserEmail = getCookie("email");
    });
</script>

<Navigation {role} />

<div id="itemsLinks">
    {#if role.isAdmin}
        <div>
            <a href="accounts/add">Add Account</a>
        </div>
        <div>
            <a href="accounts/show">Edit Accounts</a>
        </div>
    {:else if role.isStudent || role.isProfessor}
        <div>
            <a href={`accounts/edit/${currentUserEmail}`}>Edit Account Details</a>
        </div>
    {/if}
</div>
