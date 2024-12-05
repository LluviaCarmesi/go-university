<script lang="ts">
    import Navigation from "../../../../components/Navigation.svelte";
    import type IRole from "../../../../interfaces/IRole";
    import "../../../../styles/items.css";
    import "../../../../styles/common.css";
    import type IRegistration from "../../../../interfaces/IRegistration";
    import checkCurrentUser from "../../../../services/users/checkCurrentUser";
    import { onMount } from "svelte";
    import deleteRegistration from "../../../../services/registrations/deleteRegistration";
    import getRegistrationsByProfessorEmailAndID from "../../../../services/registrations/getRegistrationsByProfessorEmailAndID";
    export let data;

    let role: IRole = {
        isAdmin: false,
        isProfessor: false,
        isStudent: false,
    };
    let isLoading = false;
    let errorMessage = "";
    let isSuccessful = false;

    let registrations: IRegistration[] = [];
    let grades: number[] = [];
    let average = 0;

    onMount(() => {
        async function getRole() {
            const checkCurrentUserResponse = await checkCurrentUser();
            role = checkCurrentUserResponse;
        }
        getRole();
    });

    async function removeRegistration(registration: IRegistration) {
        isLoading = true;
        const deleteRegistrationResponse =
            await deleteRegistration(registration);
        isSuccessful = !deleteRegistrationResponse.doesErrorExist;
        if (!isSuccessful) {
            errorMessage = deleteRegistrationResponse.errorMessage;
        } else {
            errorMessage = "";
            window.open(`${data.id}`, "_self");
        }
        isLoading = false;
    }

    function calculateAverage() {
        for (let i = 0; i < registrations.length; i++) {
            grades.push(registrations[i].FinalGrade);
            average += registrations[i].FinalGrade;
        }
        average = average / registrations.length;
    }

    async function getRegistrationsResponse() {
        const registrationsResponse =
            await getRegistrationsByProfessorEmailAndID(data.id);
        if (registrationsResponse.isSuccessful) {
            registrations = registrationsResponse.registrations;
            calculateAverage();
        } else {
            errorMessage = registrationsResponse.errorMessage;
        }
        isLoading = false;
    }
    getRegistrationsResponse();
</script>

<Navigation {role} />

<div id="itemsTable">
    {#if registrations.length === 0}
        <div class="descriptionContainer">
            <span>No registrations as of now!</span>
        </div>
    {:else}
    <div class="descriptionContainer">
        <span>Total number of students: {registrations.length}</span>
    </div>
        <table>
            <thead>
                <tr>
                    <th>Average</th>
                    <th># of A</th>
                    <th># of A-</th>
                    <th># of B+</th>
                    <th># of B</th>
                    <th># of B-</th>
                    <th># of C+</th>
                    <th># of C</th>
                    <th># of C-</th>
                    <th># of D+</th>
                    <th># of D</th>
                    <th># of D-</th>
                    <th># of F</th>
                </tr>
            </thead>
            <tbody>
                <tr>
                    <td>{average.toFixed(2)}</td>
                    <td>{grades.filter((grade) => grade >= 95).length}</td>
                    <td
                        >{grades.filter((grade) => grade < 95 && grade >= 90)
                            .length}</td
                    >
                    <td
                        >{grades.filter((grade) => grade < 90 && grade >= 87)
                            .length}</td
                    >
                    <td
                        >{grades.filter((grade) => grade < 87 && grade >= 84)
                            .length}</td
                    >
                    <td
                        >{grades.filter((grade) => grade < 84 && grade >= 80)
                            .length}</td
                    >
                    <td
                        >{grades.filter((grade) => grade < 80 && grade >= 77)
                            .length}</td
                    >
                    <td
                        >{grades.filter((grade) => grade < 77 && grade >= 74)
                            .length}</td
                    >
                    <td
                        >{grades.filter((grade) => grade < 74 && grade >= 70)
                            .length}</td
                    >
                    <td
                        >{grades.filter((grade) => grade < 70 && grade >= 67)
                            .length}</td
                    >
                    <td
                        >{grades.filter((grade) => grade < 67 && grade >= 64)
                            .length}</td
                    >
                    <td
                        >{grades.filter((grade) => grade < 64 && grade >= 60)
                            .length}</td
                    >
                    <td>{grades.filter((grade) => grade < 60).length}</td>
                </tr>
            </tbody>
        </table><br>
        <table>
            <thead>
                <tr>
                    <th>Taught Course ID</th>
                    <th>Student Email</th>
                    <th>Course ID</th>
                    <th>Semester</th>
                    <th>Final Grade</th>
                    <th>Status</th>
                    <th>Remove</th>
                </tr>
            </thead>
            {#each registrations as registration}
                <tbody>
                    <tr>
                        <td>
                            <a
                                href={"/registrations/edit/" +
                                    registration.StudentEmail +
                                    "-" +
                                    registration.TaughtCourseID}
                            >
                                <span>{registration.TaughtCourseID}</span>
                            </a>
                        </td>
                        <td>
                            <span>{registration.StudentEmail}</span>
                        </td>
                        <td>
                            <span>{registration.CourseID}</span>
                        </td>
                        <td>
                            <span>{registration.SemesterName}</span>
                        </td>
                        <td>
                            <span>{registration.FinalGrade}</span>
                        </td>
                        <td>
                            <span>{registration.Status}</span>
                        </td>
                        <td>
                            <button
                                on:click={() =>
                                    removeRegistration(registration)}
                                >Remove</button
                            >
                        </td>
                    </tr>
                </tbody>
            {/each}
        </table>
    {/if}
</div>
