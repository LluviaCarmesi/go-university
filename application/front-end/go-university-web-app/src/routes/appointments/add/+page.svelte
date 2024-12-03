<script lang="ts">
    import Navigation from "../../../components/Navigation.svelte";
    import TextField from "../../../components/TextField.svelte";
    import type IRole from "../../../interfaces/IRole";
    import "../../../styles/items.css";
    import "../../../styles/common.css";
    import Checkbox from "../../../components/Checkbox.svelte";
    import { onMount } from "svelte";
    import getCookie from "../../../utilities/getCookie";
    import addAppointment from "../../../services/appointments/addAppointment";
    import DateTimePicker from "../../../components/DateTimePicker.svelte";

    let role: IRole = {
        isAdmin: false,
        isProfessor: false,
        isStudent: false,
    };

    let isLoading = false;
    let errorMessage = "";
    let isSuccessful = false;

    const appointmentTextFields: any = {
        StudentEmail: "",
        AdminEmail: "",
        StartTime: "",
        EndTime: "",
    };

    const appointmentNumberFields: any = {
        ID: 0,
    };

    const appointmentCheckboxFields: any = {
        IsComplete: false,
    };

    const appointmentDateTimeFields: any = {
        StartTime: new Date(),
        EndTime: new Date(),
    };

    onMount(() => {
        appointmentTextFields.StudentEmail = getCookie("email");
    });

    function handleTextChange(event: any) {
        appointmentTextFields[event.target.id] = event.target.value;
    }

    function handleCheckboxChange(event: any) {
        appointmentCheckboxFields[event.target.id] = event.target.checked;
    }

    function handleDateTimePickerChange(event: any) {
        appointmentTextFields[event.target.id] = event.target.value;
        appointmentDateTimeFields[event.target.id] = new Date(
            event.target.value,
        );
    }

    async function submitAppointment() {
        isLoading = true;
        const addAppointmentResponse = await addAppointment({
            ID: appointmentNumberFields.ID,
            StudentEmail: appointmentTextFields.StudentEmail,
            AdminEmail: appointmentTextFields.AdminEmail,
            IsComplete: appointmentCheckboxFields.IsComplete,
            StartTime: appointmentDateTimeFields.StartTime,
            EndTime: appointmentDateTimeFields.EndTime,
        });
        isSuccessful = !addAppointmentResponse.doesErrorExist;
        if (!isSuccessful) {
            errorMessage = addAppointmentResponse.errorMessage;
        } else {
            errorMessage = "";
        }
        isLoading = false;
    }
</script>

<Navigation {role} />

<div id="itemForm">
    <div class="mainHeadingContainer">
        <h2>Add an Appointment</h2>
    </div>
    <div class="descriptionContainer">
        <span>Add an appointment using the form below! </span>
    </div>
    {#if !!errorMessage}
        <div class="errorContainer">
            <span>{errorMessage}</span>
        </div>
    {/if}
    {#if isSuccessful}
        <div class="successContainer">
            <span>Appointment was added successfully!</span>
        </div>
    {/if}
    <TextField
        fieldLabel="Student Email"
        currentValue={appointmentTextFields.StudentEmail}
        onChangeTextField={handleTextChange}
        inputID="StudentEmail"
        isDisabled={role.isStudent}
    />
    <TextField
        fieldLabel="Admin Email"
        currentValue={appointmentTextFields.AdminEmail}
        onChangeTextField={handleTextChange}
        inputID="AdminEmail"
    />
    <Checkbox
        fieldLabel="Is Complete"
        currentValue={appointmentCheckboxFields.IsComplete}
        onChangeCheckbox={handleCheckboxChange}
        inputID="IsComplete"
        isDisabled={!role.isAdmin}
    />
    <DateTimePicker
        fieldLabel="Start Time"
        currentValue={appointmentTextFields.StartTime}
        onChangeDateTimePicker={handleDateTimePickerChange}
        inputID="StartTime"
    />
    <DateTimePicker
        fieldLabel="End Time"
        currentValue={appointmentTextFields.EndTime}
        onChangeDateTimePicker={handleDateTimePickerChange}
        inputID="EndTime"
    />
    <div class="actionsContainer">
        <a href="/appointments">Cancel</a>
        <button on:click={submitAppointment}>Submit Appointment</button>
    </div>
</div>
