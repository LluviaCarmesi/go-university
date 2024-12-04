<script lang="ts">
    import Navigation from "../../../../components/Navigation.svelte";
    import TextField from "../../../../components/TextField.svelte";
    import type IRole from "../../../../interfaces/IRole";
    import "../../../../styles/items.css";
    import "../../../../styles/common.css";
    import Checkbox from "../../../../components/Checkbox.svelte";
    import editAppointment from "../../../../services/appointments/editAppointment";
    import getAppointmentByID from "../../../../services/appointments/getAppointmentByID";
    import formatDateTime from "../../../../utilities/formatDateTime";
    import DateTimePicker from "../../../../components/DateTimePicker.svelte";
    import { onMount } from "svelte";
    import checkCurrentUser from "../../../../services/users/checkCurrentUser";
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

    const appointmentTextFields: any = {
        StudentEmail: "",
        AdminEmail: "",
        StartTime: "",
        EndTime: "",
        Description: "",
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
        const editCourseResponse = await editAppointment({
            ID: appointmentNumberFields.ID,
            StudentEmail: appointmentTextFields.StudentEmail,
            AdminEmail: appointmentTextFields.AdminEmail,
            IsComplete: appointmentCheckboxFields.IsComplete,
            StartTime: appointmentDateTimeFields.StartTime,
            EndTime: appointmentDateTimeFields.EndTime,
            Description: appointmentTextFields.Description,
        });
        isSuccessful = !editCourseResponse.doesErrorExist;
        if (!isSuccessful) {
            errorMessage = editCourseResponse.errorMessage;
        } else {
            errorMessage = "";
        }
        isLoading = false;
    }

    async function getAppointmentByIDResponse() {
        const appointmentResponse = await getAppointmentByID(data.id);
        if (appointmentResponse.isSuccessful) {
            const appointment = appointmentResponse.appointment;
            appointmentNumberFields.ID = appointment.ID;
            appointmentTextFields.StudentEmail = appointment.StudentEmail;
            appointmentTextFields.AdminEmail = appointment.AdminEmail;
            appointmentCheckboxFields.IsComplete = appointment.IsComplete;
            appointmentDateTimeFields.StartTime = appointment.StartTime;
            appointmentTextFields.StartTime = formatDateTime(
                appointment.StartTime,
                true,
                false,
            );
            appointmentDateTimeFields.EndTime = appointment.EndTime;
            appointmentTextFields.EndTime = formatDateTime(
                appointment.EndTime,
                true,
                false,
            );
            appointmentTextFields.Description = appointment.Description;
        } else {
            errorMessage = appointmentResponse.errorMessage;
        }
        isLoading = false;
    }
    getAppointmentByIDResponse();
</script>

<Navigation {role} />

<div id="itemForm">
    <div class="mainHeadingContainer">
        <h2>Edit the Appointment</h2>
    </div>
    <div class="descriptionContainer">
        <span>Edit the appointment using the form below! </span>
    </div>
    {#if !!errorMessage}
        <div class="errorContainer">
            <span>{errorMessage}</span>
        </div>
    {/if}
    {#if isSuccessful}
        <div class="successContainer">
            <span>Appointment was modified successfully!</span>
        </div>
    {/if}
    {#if !!errorMessage}
        <div>
            <span class="error">{errorMessage}</span>
        </div>
    {/if}
    <TextField
        fieldLabel="Student Email"
        currentValue={appointmentTextFields.StudentEmail}
        onChangeTextField={handleTextChange}
        inputID="ID"
        isDisabled={!role.isAdmin}
    />
    <TextField
        fieldLabel="Admin Email"
        currentValue={appointmentTextFields.AdminEmail}
        onChangeTextField={handleTextChange}
        inputID="Name"
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
    <TextField
        fieldLabel="Description"
        currentValue={appointmentTextFields.Description}
        onChangeTextField={handleTextChange}
        inputID="Description"
        isDisabled={!role.isAdmin}
        isExpandable={true}
    />
    <div class="actionsContainer">
        <a href="/appointments">Cancel</a>
        <button on:click={submitAppointment}>Submit Appointment</button>
    </div>
</div>
