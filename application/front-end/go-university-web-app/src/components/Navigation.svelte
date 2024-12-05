<script lang="ts">
    import "../styles/components/navigation.css";
    import * as strings from "../strings/ENUSStrings";
    import type IRole from "../interfaces/IRole";
    import { onMount } from "svelte";
    import getCookie from "../utilities/getCookie";

    export let role: IRole;
    let email = "";
    onMount(() => {
        email = getCookie("email");
    });
</script>

<div class="linkContainer">
    <div class="titleLinkContainer">
        <a class="titleLink link" href="/">{strings.GO_UNIVERSITY}</a>
    </div>
    {#if role.isAdmin}
        <div>
            <a href="/courses">Courses</a>
            <a href="/appointments">Appointments</a>
            <a href="/accounts">Accounts</a>
            <a href="/semesters">Semesters</a>
            <a href="/departments">Departments</a>
            <a href="/logout">Logout</a>
        </div>
    {:else if role.isProfessor}
        <div>
            <a href="/courses">Courses</a>
            <a href={`/accounts/edit/${email}`}>Edit Account</a>
            <a href="/logout">Logout</a>
        </div>
    {:else if role.isStudent}
        <div>
            <a href="/courses">Courses</a>
            <a href="/appointments">Appointments</a>
            <a href={`/accounts/edit/${email}`}>Edit Account</a>
            <a href="/logout">Logout</a>
        </div>
    {/if}
</div>
