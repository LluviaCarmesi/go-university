<script lang="ts">
    import Navigation from "../../../components/Navigation.svelte";
    import type IRole from "../../../interfaces/IRole";
    import "../../../styles/items.css";
    import "../../../styles/common.css";
    import type IProfessorInDepartment from "../../../interfaces/IProfessorToDepartment";
    import getProfessorsInDepartments from "../../../services/departments/getProfessorsInDepartments";
    import checkCurrentUser from "../../../services/users/checkCurrentUser";
    import { onMount } from "svelte";
    import deleteProfessorInDepartment from "../../../services/departments/deleteProfessorInDepartment";

    let role: IRole = {
        isAdmin: false,
        isProfessor: false,
        isStudent: false,
    };
    let isLoading = false;
    let errorMessage = "";
    let isSuccessful = false;

    onMount(() => {
        async function getRole() {
            const checkCurrentUserResponse = await checkCurrentUser();
            role = checkCurrentUserResponse;
        }
        getRole();
    });

    let professorsInDepartments: IProfessorInDepartment[] = [];

    async function removeProfessorInDepartment(
        professorInDepartment: IProfessorInDepartment,
    ) {
        isLoading = true;
        const deleteProfessorInDepartmentResponse =
            await deleteProfessorInDepartment(professorInDepartment);
        isSuccessful = !deleteProfessorInDepartmentResponse.doesErrorExist;
        if (!isSuccessful) {
            errorMessage = deleteProfessorInDepartmentResponse.errorMessage;
        } else {
            errorMessage = "";
            window.open("show-professors", "_self");
        }
        isLoading = false;
    }

    async function getProfessorsInDepartmentsResponse() {
        const professorsInDepartmentsResponse =
            await getProfessorsInDepartments();
        if (professorsInDepartmentsResponse.isSuccessful) {
            professorsInDepartments =
                professorsInDepartmentsResponse.professorsInDepartments;
        } else {
            errorMessage = professorsInDepartmentsResponse.errorMessage;
        }
        isLoading = false;
    }
    getProfessorsInDepartmentsResponse();
</script>

<Navigation {role} />

<div id="itemsTable">
    {#if professorsInDepartments.length === 0}
        <div class="descriptionContainer">
            <span>No professors as of now!</span>
        </div>
    {:else}
        <table>
            <thead>
                <tr>
                    <th>Professor Email</th>
                    <th>Department Name</th>
                    <th>Is Head</th>
                    <th>Remove</th>
                </tr>
            </thead>
            {#each professorsInDepartments as professorInDepartment}
                <tbody>
                    <tr>
                        <td>
                            <a
                                href={"edit-professor/" +
                                    professorInDepartment.ProfessorEmail +
                                    "-" +
                                    professorInDepartment.DepartmentID}
                            >
                                <span
                                    >{professorInDepartment.ProfessorEmail}</span
                                >
                            </a>
                        </td>
                        <td>
                            <span>{professorInDepartment.DepartmentName}</span>
                        </td>
                        <td>
                            <span>{professorInDepartment.IsLeader}</span>
                        </td>
                        <td>
                            <button
                                on:click={() =>
                                    removeProfessorInDepartment(
                                        professorInDepartment,
                                    )}>Remove</button
                            >
                        </td>
                    </tr>
                </tbody>
            {/each}
        </table>
    {/if}
</div>
