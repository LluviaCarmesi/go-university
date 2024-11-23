<script lang="ts">
    import Navigation from "../../../components/Navigation.svelte";
    import type IRole from "../../../interfaces/IRole";
    import "../../../styles/items.css";
    import "../../../styles/common.css";
    import getSemesters from "../../../services/semesters/getSemesters";
    import type ISemester from "../../../interfaces/ISemester";
    import formatDateTime from "../../../utilities/formatDateTime";

    let role: IRole = {
        isAdmin: true,
        isProfessor: false,
        isStudent: false,
    };
    let isLoading = false;
    let errorMessage = "";

    let semesters: ISemester[] = [];

    async function getSemestersResponse() {
        const semestersResponse = await getSemesters();
        if (semestersResponse.isSuccessful) {
            semesters = semestersResponse.semesters;
        } else {
            errorMessage = semestersResponse.errorMessage;
        }
        isLoading = false;
    }
    getSemestersResponse();
</script>

<Navigation {role} />

<div id="itemsTable">
    {#if semesters.length === 0}
        <div class="descriptionContainer">
            <span>No semesters as of now!</span>
        </div>
    {:else}
        <table>
            <thead>
                <tr>
                    <th>Name</th>
                    <th>Start Date</th>
                    <th>End Date</th>
                </tr>
            </thead>
            {#each semesters as semester}
                <tbody>
                    <tr>
                        <td>
                            <a href={"edit/" + semester.Name}>
                                <span>{semester.Name}</span>
                            </a>
                        </td>
                        <td>
                            <span
                                >{formatDateTime(
                                    semester.StartDate,
                                    false,
                                    false,
                                )}</span
                            >
                        </td>
                        <td>
                            <span
                                >{formatDateTime(
                                    semester.EndDate,
                                    false,
                                    false,
                                )}</span
                            >
                        </td>
                    </tr>
                </tbody>
            {/each}
        </table>
    {/if}
</div>
