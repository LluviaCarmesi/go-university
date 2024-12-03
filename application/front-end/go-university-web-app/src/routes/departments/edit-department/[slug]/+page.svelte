<script lang="ts">
    import Navigation from "../../../../components/Navigation.svelte";
    import TextField from "../../../../components/TextField.svelte";
    import type IRole from "../../../../interfaces/IRole";
    import "../../../../styles/items.css";
    import "../../../../styles/common.css";
    import getDepartmentByID from "../../../../services/departments/getDepartmentByID";
    import editDepartment from "../../../../services/departments/editDepartment";
    export let data;

    let role: IRole = {
        isAdmin: false,
        isProfessor: false,
        isStudent: false,
    };

    let isLoading = false;
    let errorMessage = "";
    let isSuccessful = false;

    const departmentNumberFields: any = {
        ID: 0,
    };

    const departmentTextFields: any = {
        Name: "",
    };

    function handleTextChange(event: any) {
        departmentTextFields[event.target.id] = event.target.value;
    }

    async function submitDepartment() {
        isLoading = true;
        const editDepartmentResponse = await editDepartment({
            ID: departmentNumberFields.ID,
            Name: departmentTextFields.Name,
        });
        isSuccessful = !editDepartmentResponse.doesErrorExist;
        if (!isSuccessful) {
            errorMessage = editDepartmentResponse.errorMessage;
        } else {
            errorMessage = "";
        }
        isLoading = false;
    }

    async function getDepartmentByIDResponse() {
        const departmentResponse = await getDepartmentByID(data.id);
        if (departmentResponse.isSuccessful) {
            const department = departmentResponse.department;
            departmentNumberFields.ID = department.ID;
            departmentTextFields.Name = department.Name;
        } else {
            errorMessage = departmentResponse.errorMessage;
        }
        isLoading = false;
    }
    getDepartmentByIDResponse();
</script>

<Navigation {role} />

<div id="itemForm">
    <div class="mainHeadingContainer">
        <h2>Edit the Department</h2>
    </div>
    <div class="descriptionContainer">
        <span>Edit the department using the form below! </span>
    </div>
    {#if !!errorMessage}
        <div class="errorContainer">
            <span>{errorMessage}</span>
        </div>
    {/if}
    {#if isSuccessful}
        <div class="successContainer">
            <span>Department was modified successfully!</span>
        </div>
    {/if}
    {#if !!errorMessage}
        <div>
            <span class="error">{errorMessage}</span>
        </div>
    {/if}
    <TextField
        fieldLabel="Name"
        currentValue={departmentTextFields.Name}
        onChangeTextField={handleTextChange}
        inputID="Name"
    />
    <div class="actionsContainer">
        <a href="/departments">Cancel</a>
        <button on:click={submitDepartment}>Submit Department</button>
    </div>
</div>
