<script lang="ts">
    import Navigation from "../../../components/Navigation.svelte";
    import TextField from "../../../components/TextField.svelte";
    import type IRole from "../../../interfaces/IRole";
    import "../../../styles/items.css";
    import "../../../styles/common.css";
    import DatePicker from "../../../components/DatePicker.svelte";
    import addSemester from "../../../services/semesters/addSemester";
    import formatDateTime from "../../../utilities/formatDateTime";

    let role: IRole = {
        isAdmin: true,
        isProfessor: false,
        isStudent: false,
    };

    let isLoading = false;
    let errorMessage = "";
    let isSuccessful = false;

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
        const addUserResponse = await addSemester({
            Name: semesterTextFields.Name,
            StartDate: semesterDateFields.StartDate,
            EndDate: semesterDateFields.EndDate,
        });
        isSuccessful = !addUserResponse.doesErrorExist;
        if (!isSuccessful) {
            errorMessage = addUserResponse.errorMessage;
        } else {
            errorMessage = "";
            semesterTextFields.Name = "";
            semesterDateFields.StartDate = "";
            semesterDateFields.EndDate = "";
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
