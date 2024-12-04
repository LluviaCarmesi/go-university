<script lang="ts">
    import Navigation from "../../../components/Navigation.svelte";
    import type IAppointment from "../../../interfaces/IAppointment";
    import type IRole from "../../../interfaces/IRole";
    import getAppointments from "../../../services/appointments/getAppointments";
    import "../../../styles/items.css";
    import "../../../styles/common.css";
    import formatDateTime from "../../../utilities/formatDateTime";
    import checkCurrentUser from "../../../services/users/checkCurrentUser";
    import { onMount } from "svelte";
    import deleteAppointment from "../../../services/appointments/deleteAppointment";

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

    let appointments: IAppointment[] = [];

    async function removeAppointment(appointment: IAppointment) {
        isLoading = true;
        const deleteAppointmentResponse = await deleteAppointment(appointment);
        isSuccessful = !deleteAppointmentResponse.doesErrorExist;
        if (!isSuccessful) {
            errorMessage = deleteAppointmentResponse.errorMessage;
        } else {
            errorMessage = "";
            window.open("show", "_self");
        }
        isLoading = false;
    }

    async function getAppointmentsResponse() {
        const appointmentsResponse = await getAppointments();
        if (appointmentsResponse.isSuccessful) {
            appointments = appointmentsResponse.appointments;
        } else {
            errorMessage = appointmentsResponse.errorMessage;
        }
        isLoading = false;
    }
    getAppointmentsResponse();
</script>

<Navigation {role} />

<div id="itemsTable">
    {#if appointments.length === 0}
        <div class="descriptionContainer">
            <span>No appointments as of now!</span>
        </div>
    {:else}
        <table>
            <thead>
                <tr>
                    <th>ID</th>
                    <th>Student Email</th>
                    <th>Admin Email</th>
                    <th>Is Complete</th>
                    <th>Start Time</th>
                    <th>End Time</th>
                    <th>Description</th>
                    <th>Remove</th>
                </tr>
            </thead>
            {#each appointments as appointment}
                <tbody>
                    <tr>
                        <td>
                            <a href={"edit/" + appointment.ID}>
                                <span>{appointment.ID}</span>
                            </a>
                        </td>
                        <td>
                            <span>{appointment.StudentEmail}</span>
                        </td>
                        <td>
                            <span>{appointment.AdminEmail}</span>
                        </td>
                        <td>
                            <span>{appointment.IsComplete}</span>
                        </td>
                        <td>
                            <span
                                >{formatDateTime(
                                    appointment.StartTime,
                                    true,
                                    true,
                                )}</span
                            >
                        </td>
                        <td>
                            <span
                                >{formatDateTime(
                                    appointment.EndTime,
                                    true,
                                    true,
                                )}</span
                            >
                        </td>
                        <td>
                            <span>{appointment.Description}</span>
                        </td>
                        <td>
                            <button
                                on:click={() => removeAppointment(appointment)}
                                >Remove</button
                            >
                        </td>
                    </tr>
                </tbody>
            {/each}
        </table>
    {/if}
</div>
