<script lang="ts">
    import Navigation from "../../../components/Navigation.svelte";
    import type IRole from "../../../interfaces/IRole";
    import "../../../styles/items.css";
    import "../../../styles/common.css";
    import type IDepartment from "../../../interfaces/IDepartment";
    import getDepartments from "../../../services/departments/getDepartments";
    import checkCurrentUser from "../../../services/users/checkCurrentUser";
    import { onMount } from "svelte";
    import deleteDepartment from "../../../services/departments/deleteDepartment";

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

    let departments: IDepartment[] = [];

    async function removeDepartment(department: IDepartment) {
        isLoading = true;
        const deleteDepartmentResponse = await deleteDepartment(department);
        isSuccessful = !deleteDepartmentResponse.doesErrorExist;
        if (!isSuccessful) {
            errorMessage = deleteDepartmentResponse.errorMessage;
        } else {
            errorMessage = "";
            window.open("show-departments", "_self");
        }
        isLoading = false;
    }

    async function getDepartmentsResponse() {
        const departmentsResponse = await getDepartments();
        if (departmentsResponse.isSuccessful) {
            departments = departmentsResponse.departments;
        } else {
            errorMessage = departmentsResponse.errorMessage;
        }
        isLoading = false;
    }
    getDepartmentsResponse();
</script>

<Navigation {role} />

<div id="itemsTable">
    {#if departments.length === 0}
        <div class="descriptionContainer">
            <span>No departments as of now!</span>
        </div>
    {:else}
        <table>
            <thead>
                <tr>
                    <th>ID</th>
                    <th>Name</th>
                    <th>Remove</th>
                </tr>
            </thead>
            {#each departments as department}
                <tbody>
                    <tr>
                        <td>
                            <a href={"edit-department/" + department.ID}>
                                <span>{department.ID}</span>
                            </a>
                        </td>
                        <td>
                            <span>{department.Name}</span>
                        </td>
                        <td>
                            <button
                                on:click={() => removeDepartment(department)}
                                >Remove</button
                            >
                        </td>
                    </tr>
                </tbody>
            {/each}
        </table>
    {/if}
</div>
