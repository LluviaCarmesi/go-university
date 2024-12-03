<script lang="ts">
    import Navigation from "../../../components/Navigation.svelte";
    import type IRole from "../../../interfaces/IRole";
    import "../../../styles/items.css";
    import "../../../styles/common.css";
    import getRegistrations from "../../../services/registrations/getRegistrations";
    import type IRegistration from "../../../interfaces/IRegistration";
    import checkCurrentUser from "../../../services/users/checkCurrentUser";
    import { onMount } from "svelte";

    let role: IRole = {
        isAdmin: false,
        isProfessor: false,
        isStudent: false,
    };
    let isLoading = false;
    let errorMessage = "";

    let registrations: IRegistration[] = [];

    onMount(() => {
        async function getRole() {
            const checkCurrentUserResponse = await checkCurrentUser();
            role = checkCurrentUserResponse;
        }
        getRole();
    });

    async function getRegistrationsResponse() {
        const registrationsResponse = await getRegistrations();
        if (registrationsResponse.isSuccessful) {
            registrations = registrationsResponse.registrations;
        } else {
            errorMessage = registrationsResponse.errorMessage;
        }
        isLoading = false;
    }
    getRegistrationsResponse();
</script>

<Navigation {role} />

<div id="itemsTable">
    {#if registrations.length === 0}
        <div class="descriptionContainer">
            <span>No registrations as of now!</span>
        </div>
    {:else}
        <table>
            <thead>
                <tr>
                    <th>Taught Course ID</th>
                    <th>Student Email</th>
                    <th>Course ID</th>
                    <th>Final Grade</th>
                    <th>Status</th>
                </tr>
            </thead>
            {#each registrations as registration}
                <tbody>
                    <tr>
                        <td>
                            <span>{registration.TaughtCourseID}</span>
                        </td>
                        <td>
                            <span>{registration.StudentEmail}</span>
                        </td>
                        <td>
                            <span>{registration.CourseID}</span>
                        </td>
                        <td>
                            <span>{registration.FinalGrade}</span>
                        </td>
                        <td>
                            <span>{registration.Status}</span>
                        </td>
                    </tr>
                </tbody>
            {/each}
        </table>
    {/if}
</div>
