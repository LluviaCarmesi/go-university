<script lang="ts">
    import Navigation from "../../../components/Navigation.svelte";
    import type IRole from "../../../interfaces/IRole";
    import "../../../styles/items.css";
    import "../../../styles/common.css";
    import type IProfessorInDepartment from "../../../interfaces/IProfessorToDepartment";
    import getProfessorsInDepartments from "../../../services/departments/getProfessorsInDepartments";

    let role: IRole = {
        isAdmin: true,
        isProfessor: false,
        isStudent: false,
    };
    let isLoading = false;
    let errorMessage = "";

    let professorsInDepartments: IProfessorInDepartment[] = [];

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
                </tr>
            </thead>
            {#each professorsInDepartments as professorsInDepartment}
                <tbody>
                    <tr>
                        <td>
                            <a
                                href={"edit-professor/" +
                                    professorsInDepartment.ProfessorEmail +
                                    "-" +
                                    professorsInDepartment.DepartmentID}
                            >
                                <span
                                    >{professorsInDepartment.ProfessorEmail}</span
                                >
                            </a>
                        </td>
                        <td>
                            <span>{professorsInDepartment.DepartmentName}</span>
                        </td>
                        <td>
                            <span>{professorsInDepartment.IsLeader}</span>
                        </td>
                    </tr>
                </tbody>
            {/each}
        </table>
    {/if}
</div>
