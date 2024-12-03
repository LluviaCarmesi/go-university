<script lang="ts">
    import Navigation from "../../../../components/Navigation.svelte";
    import TextField from "../../../../components/TextField.svelte";
    import type IRole from "../../../../interfaces/IRole";
    import "../../../../styles/items.css";
    import "../../../../styles/common.css";
    import DatePicker from "../../../../components/DatePicker.svelte";
    import editSemester from "../../../../services/semesters/editSemester";
    import getSemesterByName from "../../../../services/semesters/getSemesterByName";
    import formatDateTime from "../../../../utilities/formatDateTime";
    export let data;

    let role: IRole = {
        isAdmin: false,
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
        const editSemesterResponse = await editSemester({
            Name: semesterTextFields.Name,
            StartDate: semesterDateFields.StartDate,
            EndDate: semesterDateFields.EndDate,
        });
        isSuccessful = !editSemesterResponse.doesErrorExist;
        if (!isSuccessful) {
            errorMessage = editSemesterResponse.errorMessage;
        } else {
            errorMessage = "";
        }
        isLoading = false;
    }

    async function getSemesterByNameResponse() {
        const semesterResponse = await getSemesterByName(data.id);
        if (semesterResponse.isSuccessful) {
            const semester = semesterResponse.semester;
            semesterTextFields.Name = semester.Name;
            semesterDateFields.StartDate = semester.StartDate;
            semesterTextFields.StartDate = formatDateTime(
                semester.StartDate,
                false,
                false,
            );
            semesterDateFields.EndDate = semester.EndDate;
            semesterTextFields.EndDate = formatDateTime(
                semester.EndDate,
                false,
                false,
            );
        } else {
            errorMessage = semesterResponse.errorMessage;
        }
        isLoading = false;
    }
    getSemesterByNameResponse();
</script>

<Navigation {role} />

<div id="itemForm">
    <div class="mainHeadingContainer">
        <h2>Edit the Semester</h2>
    </div>
    <div class="descriptionContainer">
        <span>Edit the semester using the form below! </span>
    </div>
    {#if !!errorMessage}
        <div class="errorContainer">
            <span>{errorMessage}</span>
        </div>
    {/if}
    {#if isSuccessful}
        <div class="successContainer">
            <span>Semester was modified successfully!</span>
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
        isDisabled={true}
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
