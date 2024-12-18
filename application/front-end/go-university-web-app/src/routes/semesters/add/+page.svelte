<script lang="ts">
    import Navigation from "../../../components/Navigation.svelte";
    import TextField from "../../../components/TextField.svelte";
    import type IRole from "../../../interfaces/IRole";
    import "../../../styles/items.css";
    import "../../../styles/common.css";
    import DatePicker from "../../../components/DatePicker.svelte";
    import addSemester from "../../../services/semesters/addSemester";
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

    const semesterTextFields: any = {
        Name: "",
        StartDate: "",
        EndDate: "",
    };

    const semesterDateFields: any = {
        StartDate: new Date(),
        EndDate: new Date(),
    };

    function handleTextChange(event: any) {
        semesterTextFields[event.target.id] = event.target.value;
    }

    function handleDatePickerChange(event: any) {
        semesterTextFields[event.target.id] = event.target.value;
        semesterDateFields[event.target.id] = new Date(event.target.value);
    }

    async function submitSemester() {
        isLoading = true;
        const addSemesterResponse = await addSemester({
            Name: semesterTextFields.Name,
            StartDate: semesterDateFields.StartDate,
            EndDate: semesterDateFields.EndDate,
        });
        isSuccessful = !addSemesterResponse.doesErrorExist;
        if (!isSuccessful) {
            errorMessage = addSemesterResponse.errorMessage;
        } else {
            errorMessage = "";
            semesterTextFields.Name = "";
            semesterDateFields.StartDate = new Date();
            semesterTextFields.StartDate = "";
            semesterDateFields.EndDate = new Date();
            semesterTextFields.EndDate = "";
        }
        isLoading = false;
    }
</script>

<Navigation {role} />

<div id="itemForm">
    <div class="mainHeadingContainer">
        <h2>Add a Semester</h2>
    </div>
    <div class="descriptionContainer">
        <span>Add a semester using the form below! </span>
    </div>
    {#if !!errorMessage}
        <div class="errorContainer">
            <span>{errorMessage}</span>
        </div>
    {/if}
    {#if isSuccessful}
        <div class="successContainer">
            <span>Semester was added successfully!</span>
        </div>
    {/if}
    {#if !!errorMessage}
        <div>
            <span class="error">{errorMessage}</span>
        </div>
    {/if}
    <TextField
        fieldLabel="Name"
        currentValue={semesterTextFields.Name}
        onChangeTextField={handleTextChange}
        inputID="Name"
    />
    <DatePicker
        fieldLabel="Start Date"
        currentValue={semesterTextFields.StartDate}
        onChangeDatePicker={handleDatePickerChange}
        inputID="StartDate"
    />
    <DatePicker
        fieldLabel="End Date"
        currentValue={semesterTextFields.EndDate}
        onChangeDatePicker={handleDatePickerChange}
        inputID="EndDate"
    />
    <div class="actionsContainer">
        <a href="/semesters">Cancel</a>
        <button on:click={submitSemester}>Submit Semester</button>
    </div>
</div>
