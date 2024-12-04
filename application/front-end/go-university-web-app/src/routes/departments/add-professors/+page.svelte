<script lang="ts">
    import Navigation from "../../../components/Navigation.svelte";
    import TextField from "../../../components/TextField.svelte";
    import type IRole from "../../../interfaces/IRole";
    import "../../../styles/items.css";
    import "../../../styles/common.css";
    import type IDepartment from "../../../interfaces/IDepartment";
    import createChoices from "../../../utilities/createChoices";
    import Dropdown from "../../../components/Dropdown.svelte";
    import Checkbox from "../../../components/Checkbox.svelte";
    import getDepartments from "../../../services/departments/getDepartments";
    import addProfessorInDepartment from "../../../services/departments/addProfessorInDepartment";
    import checkCurrentUser from "../../../services/users/checkCurrentUser";
    import { onMount } from "svelte";

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

    const professorInDepartmentDropdownFields: any = {
        DepartmentID: 0,
    };

    const professorInDepartmentTextFields: any = {
        ProfessorEmail: "",
        DepartmentName: "",
    };

    const professorInDepartmentCheckboxFields: any = {
        IsLeader: false,
    };

    function handleTextChange(event: any) {
        professorInDepartmentTextFields[event.target.id] = event.target.value;
    }

    function handleCheckboxChange(event: any) {
        professorInDepartmentCheckboxFields[event.target.id] =
            event.target.checked;
    }

    function handleDropdownChange(event: any) {
        professorInDepartmentTextFields[event.target.id] = event.target.value;
    }

    async function submitProfessorInDepartment() {
        isLoading = true;
        const addProfessorInDepartmentResponse = await addProfessorInDepartment(
            {
                ProfessorEmail: professorInDepartmentTextFields.ProfessorEmail,
                DepartmentID: professorInDepartmentDropdownFields.DepartmentID,
                DepartmentName: professorInDepartmentTextFields.DepartmentName,
                IsLeader: professorInDepartmentCheckboxFields.IsLeader,
            },
        );
        isSuccessful = !addProfessorInDepartmentResponse.doesErrorExist;
        if (!isSuccessful) {
            errorMessage = addProfessorInDepartmentResponse.errorMessage;
        } else {
            errorMessage = "";
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

<div id="itemForm">
    <div class="mainHeadingContainer">
        <h2>Add the Professor Department</h2>
    </div>
    <div class="descriptionContainer">
        <span>Add the professor department using the form below! </span>
    </div>
    {#if !!errorMessage}
        <div class="errorContainer">
            <span>{errorMessage}</span>
        </div>
    {/if}
    {#if isSuccessful}
        <div class="successContainer">
            <span>Professor department was added successfully!</span>
        </div>
    {/if}
    {#if !!errorMessage}
        <div>
            <span class="error">{errorMessage}</span>
        </div>
    {/if}
    <TextField
        fieldLabel="Professor Email"
        currentValue={professorInDepartmentTextFields.ProfessorEmail}
        onChangeTextField={handleTextChange}
        inputID="ProfessorEmail"
    />
    <Dropdown
        fieldLabel="Department"
        currentValue={professorInDepartmentDropdownFields.DepartmentID}
        onDropdownChange={handleDropdownChange}
        inputID="DepartmentID"
        dropdownOptions={createChoices(departments, "ID", "Name")}
    />
    <Checkbox
        fieldLabel="Is Head of Department"
        currentValue={professorInDepartmentCheckboxFields.IsLeader}
        onChangeCheckbox={handleCheckboxChange}
        inputID="IsLeader"
    />
    <div class="actionsContainer">
        <a href="/departments">Cancel</a>
        <button on:click={submitProfessorInDepartment}
            >Submit Professor Department</button
        >
    </div>
</div>
