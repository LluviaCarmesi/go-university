<script lang="ts">
    import Navigation from "../../../../components/Navigation.svelte";
    import TextField from "../../../../components/TextField.svelte";
    import type IRole from "../../../../interfaces/IRole";
    import "../../../../styles/items.css";
    import "../../../../styles/common.css";
    import Dropdown from "../../../../components/Dropdown.svelte";
    import { REGISTRATION_STATUS_OPTIONS } from "../../../../appSettings";
    import NumberField from "../../../../components/NumberField.svelte";
    import { onMount } from "svelte";
    import checkCurrentUser from "../../../../services/users/checkCurrentUser";
    import editRegistration from "../../../../services/registrations/editRegistration";
    import getRegistrationByID from "../../../../services/registrations/getRegistrationByID";
    export let data;

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

    const registrationTextFields: any = {
        StudentEmail: "",
        FinalGrade: "",
        CourseID: "",
        SemesterName: "",
    };

    const registrationNumberFields: any = {
        TaughtCourseID: 0,
    };

    const registrationDropdownFields: any = {
        Status: "",
    };

    function handleTextChange(event: any) {
        registrationTextFields[event.target.id] = event.target.value;
    }

    function handleDropdownChange(event: any) {
        registrationDropdownFields[event.target.id] = event.target.value;
    }

    async function submitRegistration() {
        isLoading = true;
        const editRegistrationResponse = await editRegistration({
            StudentEmail: registrationTextFields.StudentEmail,
            FinalGrade: parseFloat(registrationTextFields.FinalGrade),
            CourseID: registrationTextFields.CourseID,
            SemesterName: registrationTextFields.SemesterName,
            TaughtCourseID: registrationNumberFields.TaughtCourseID,
            Status: registrationDropdownFields.Status,
        });
        isSuccessful = !editRegistrationResponse.doesErrorExist;
        if (!isSuccessful) {
            errorMessage = editRegistrationResponse.errorMessage;
        } else {
            errorMessage = "";
        }
        isLoading = false;
    }

    async function getRegistrationByIDResponse() {
        const registrationResponse = await getRegistrationByID(data.id);
        if (registrationResponse.isSuccessful) {
            const registration = registrationResponse.registration;
            registrationNumberFields.TaughtCourseID =
                registration.TaughtCourseID;
            registrationDropdownFields.Status = registration.Status;
            registrationTextFields.StudentEmail = registration.StudentEmail;
            registrationTextFields.FinalGrade = registration.FinalGrade;
            registrationTextFields.CourseID = registration.CourseID;
            registrationTextFields.SemesterName = registration.SemesterName;
        } else {
            errorMessage = registrationResponse.errorMessage;
        }
        isLoading = false;
    }

    getRegistrationByIDResponse();
</script>

<Navigation {role} />

<div id="itemForm">
    <div class="mainHeadingContainer">
        <h2>Edit the Registration</h2>
    </div>
    <div class="descriptionContainer">
        <span>Edit the taught course using the form below! </span>
    </div>
    {#if !!errorMessage}
        <div class="errorContainer">
            <span>{errorMessage}</span>
        </div>
    {/if}
    {#if isSuccessful}
        <div class="successContainer">
            <span>Taught Course was modified successfully!</span>
        </div>
    {/if}
    {#if !!errorMessage}
        <div>
            <span class="error">{errorMessage}</span>
        </div>
    {/if}
    <TextField
        fieldLabel="Student Email"
        currentValue={registrationTextFields.StudentEmail}
        onChangeTextField={handleTextChange}
        inputID="StudentEmail"
        isDisabled={true}
    />
    <TextField
        fieldLabel="Course ID"
        currentValue={registrationTextFields.CourseID}
        onChangeTextField={handleTextChange}
        inputID="CourseID"
        isDisabled={true}
    />
    <TextField
        fieldLabel="Semester Name"
        currentValue={registrationTextFields.SemesterName}
        onChangeTextField={handleTextChange}
        inputID="SemesterName"
        isDisabled={true}
    />
    <Dropdown
        fieldLabel="Status"
        currentValue={registrationDropdownFields.Status}
        onDropdownChange={handleDropdownChange}
        inputID="Status"
        dropdownOptions={REGISTRATION_STATUS_OPTIONS}
    />
    <NumberField
        fieldLabel="Final Grade"
        currentValue={registrationNumberFields.FinalGrade}
        onChangeNumberField={handleTextChange}
        inputID="FinalGrade"
        minNumber={0}
        maxNumber={100}
    />
    <div class="actionsContainer">
        <a href="/taught-courses">Cancel</a>
        <button on:click={submitRegistration}>Submit Registration</button>
    </div>
</div>
