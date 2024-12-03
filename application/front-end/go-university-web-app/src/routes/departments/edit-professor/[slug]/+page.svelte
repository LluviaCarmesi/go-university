<script lang="ts">
    import Navigation from "../../../../components/Navigation.svelte";
    import TextField from "../../../../components/TextField.svelte";
    import type IRole from "../../../../interfaces/IRole";
    import "../../../../styles/items.css";
    import "../../../../styles/common.css";
    import getProfessorInDepartmentByID from "../../../../services/departments/getProfessorInDepartmentByID";
    import type IDepartment from "../../../../interfaces/IDepartment";
    import createChoices from "../../../../utilities/createChoices";
    import Dropdown from "../../../../components/Dropdown.svelte";
    import Checkbox from "../../../../components/Checkbox.svelte";
    import editProfessorInDepartment from "../../../../services/departments/editProfessorInDepartment";
    import getDepartments from "../../../../services/departments/getDepartments";
    export let data;

    let role: IRole = {
        isAdmin: false,
        isProfessor: false,
        isStudent: false,
    };

    let isLoading = false;
    let errorMessage = "";
    let isSuccessful = false;

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
        const editProfessorInDepartmentResponse =
            await editProfessorInDepartment({
                ProfessorEmail: professorInDepartmentTextFields.ProfessorEmail,
                DepartmentID: professorInDepartmentDropdownFields.DepartmentID,
                DepartmentName: professorInDepartmentTextFields.DepartmentName,
                IsLeader: professorInDepartmentCheckboxFields.IsLeader,
            });
        isSuccessful = !editProfessorInDepartmentResponse.doesErrorExist;
        if (!isSuccessful) {
            errorMessage = editProfessorInDepartmentResponse.errorMessage;
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

    async function getProfessorInDepartmentByIDResponse() {
        const professorInDepartmentResponse =
            await getProfessorInDepartmentByID(data.id);
        if (professorInDepartmentResponse.isSuccessful) {
            const professorInDepartment =
                professorInDepartmentResponse.professorInDepartment;
            professorInDepartmentDropdownFields.DepartmentID =
                professorInDepartment.DepartmentID;
            professorInDepartmentTextFields.ProfessorEmail =
                professorInDepartment.ProfessorEmail;
            professorInDepartmentTextFields.ProfessorEmail =
                professorInDepartment.DepartmentName;
            professorInDepartmentCheckboxFields.IsLeader =
                professorInDepartment.IsLeader;
        } else {
            errorMessage = professorInDepartmentResponse.errorMessage;
        }
        isLoading = false;
    }
    getProfessorInDepartmentByIDResponse();
    getDepartmentsResponse();
</script>

<Navigation {role} />

<div id="itemForm">
    <div class="mainHeadingContainer">
        <h2>Edit the Professor Department</h2>
    </div>
    <div class="descriptionContainer">
        <span>Edit the professor department using the form below! </span>
    </div>
    {#if !!errorMessage}
        <div class="errorContainer">
            <span>{errorMessage}</span>
        </div>
    {/if}
    {#if isSuccessful}
        <div class="successContainer">
            <span>Professor department was modified successfully!</span>
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
