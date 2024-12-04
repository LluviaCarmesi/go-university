<script lang="ts">
    import Navigation from "../../../components/Navigation.svelte";
    import TextField from "../../../components/TextField.svelte";
    import type IRole from "../../../interfaces/IRole";
    import "../../../styles/items.css";
    import "../../../styles/common.css";
    import Checkbox from "../../../components/Checkbox.svelte";
    import Dropdown from "../../../components/Dropdown.svelte";
    import { ROLE_OPTIONS } from "../../../appSettings";
    import addUser from "../../../services/users/addUser";
    import { onMount } from "svelte";
    import checkCurrentUser from "../../../services/users/checkCurrentUser";

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

    const userTextFields: any = {
        Email: "",
        EmailAlias: "",
        Password: "",
        FirstName: "",
        LastName: "",
        PhoneNumber: "",
        HomeAddress: "",
    };

    const userDropdownFields: any = {
        Role: "",
    };

    const userCheckboxFields: any = {
        MustChangePW: false,
    };

    function handleTextChange(event: any) {
        userTextFields[event.target.id] = event.target.value;
    }

    function handleDropdownChange(event: any) {
        userDropdownFields[event.target.id] = event.target.value;
    }

    function handleCheckboxChange(event: any) {
        userCheckboxFields[event.target.id] = event.target.checked;
    }

    async function submitUser() {
        isLoading = true;
        const addUserResponse = await addUser({
            Email: userTextFields.Email,
            EmailAlias: userTextFields.EmailAlias,
            Password: userTextFields.Password,
            FirstName: userTextFields.FirstName,
            LastName: userTextFields.LastName,
            PhoneNumber: userTextFields.PhoneNumber,
            HomeAddress: userTextFields.HomeAddress,
            Role: userDropdownFields.Role,
            Token: "",
            MustChangePW: userCheckboxFields.MustChangePW,
        });
        isSuccessful = !addUserResponse.doesErrorExist;
        if (!isSuccessful) {
            errorMessage = addUserResponse.errorMessage;
        } else {
            errorMessage = "";
            userTextFields.Email = "";
            userTextFields.EmailAlias = "";
            userTextFields.Password = "";
            userTextFields.FirstName = "";
            userTextFields.LastName = "";
            userTextFields.PhoneNumber = "";
            userTextFields.HomeAddress = "";
            userCheckboxFields.MustChangePW = false;
        }
        isLoading = false;
    }
</script>

<Navigation {role} />

<div id="itemForm">
    <div class="mainHeadingContainer">
        <h2>Add a User</h2>
    </div>
    <div class="descriptionContainer">
        <span>Add a user using the form below! </span>
    </div>
    {#if !!errorMessage}
        <div class="errorContainer">
            <span>{errorMessage}</span>
        </div>
    {/if}
    {#if isSuccessful}
        <div class="successContainer">
            <span>User was added successfully!</span>
        </div>
    {/if}
    {#if !!errorMessage}
        <div>
            <span class="error">{errorMessage}</span>
        </div>
    {/if}
    <TextField
        fieldLabel="Email"
        currentValue={userTextFields.Email}
        onChangeTextField={handleTextChange}
        inputID="Email"
    />
    <TextField
        fieldLabel="Email Alias"
        currentValue={userTextFields.EmailAlias}
        onChangeTextField={handleTextChange}
        inputID="EmailAlias"
    />
    <TextField
        fieldLabel="Password"
        currentValue={userTextFields.Password}
        onChangeTextField={handleTextChange}
        inputID="Password"
        isDisabled={!role.isAdmin}
    />
    <TextField
        fieldLabel="First Name"
        currentValue={userTextFields.FirstName}
        onChangeTextField={handleTextChange}
        inputID="FirstName"
    />
    <TextField
        fieldLabel="Last Name"
        currentValue={userTextFields.LastName}
        onChangeTextField={handleTextChange}
        inputID="LastName"
    />
    <TextField
        fieldLabel="Phone Number"
        currentValue={userTextFields.PhoneNumber}
        onChangeTextField={handleTextChange}
        inputID="PhoneNumber"
    />
    <TextField
        fieldLabel="Home Address"
        currentValue={userTextFields.HomeAddress}
        onChangeTextField={handleTextChange}
        inputID="HomeAddress"
        isExpandable={true}
    />
    <Dropdown
        fieldLabel="Role"
        currentValue={userDropdownFields.Role}
        onDropdownChange={handleDropdownChange}
        inputID="Role"
        dropdownOptions={ROLE_OPTIONS}
    />
    <Checkbox
        fieldLabel="Must Change Password"
        currentValue={userCheckboxFields.MustChangePW}
        onChangeCheckbox={handleCheckboxChange}
        inputID="MustChangePW"
    />
    <div class="actionsContainer">
        <a href="/accounts">Cancel</a>
        <button on:click={submitUser}>Submit User</button>
    </div>
</div>
