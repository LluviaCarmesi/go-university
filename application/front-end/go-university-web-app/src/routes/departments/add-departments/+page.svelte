<script lang="ts">
    import Navigation from "../../../components/Navigation.svelte";
    import TextField from "../../../components/TextField.svelte";
    import type IRole from "../../../interfaces/IRole";
    import "../../../styles/items.css";
    import "../../../styles/common.css";
    import addDepartment from "../../../services/departments/addDepartment";
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
        const addDepartmentResponse = await addDepartment({
            ID: departmentNumberFields.ID,
            Name: departmentTextFields.Name,
        });
        isSuccessful = !addDepartmentResponse.doesErrorExist;
        if (!isSuccessful) {
            errorMessage = addDepartmentResponse.errorMessage;
        } else {
            errorMessage = "";
            departmentTextFields.Name = "";
        }
        isLoading = false;
    }
</script>

<Navigation {role} />

<div id="itemForm">
    <div class="mainHeadingContainer">
        <h2>Add a Department</h2>
    </div>
    <div class="descriptionContainer">
        <span>Add a department using the form below! </span>
    </div>
    {#if !!errorMessage}
        <div class="errorContainer">
            <span>{errorMessage}</span>
        </div>
    {/if}
    {#if isSuccessful}
        <div class="successContainer">
            <span>Department was added successfully!</span>
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
