<script lang="ts">
    import Navigation from "../../../components/Navigation.svelte";
    import type IRole from "../../../interfaces/IRole";
    import "../../../styles/items.css";
    import "../../../styles/common.css";
    import getUsers from "../../../services/users/getUsers";
    import type IUser from "../../../interfaces/IUser";
    import checkCurrentUser from "../../../services/users/checkCurrentUser";
    import { onMount } from "svelte";
    import deleteUser from "../../../services/users/deleteUser";

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

    let users: IUser[] = [];

    async function removeAccount(account: IUser) {
        isLoading = true;
        const deleteAppointmentResponse = await deleteUser(account);
        isSuccessful = !deleteAppointmentResponse.doesErrorExist;
        if (!isSuccessful) {
            errorMessage = deleteAppointmentResponse.errorMessage;
        } else {
            errorMessage = "";
            window.open("show", "_self");
        }
        isLoading = false;
    }

    async function getUsersResponse() {
        const usersResponse = await getUsers();
        if (usersResponse.isSuccessful) {
            users = usersResponse.users;
        } else {
            errorMessage = usersResponse.errorMessage;
        }
        isLoading = false;
    }
    getUsersResponse();
</script>

<Navigation {role} />

<div id="itemsTable">
    {#if users.length === 0}
        <div class="descriptionContainer">
            <span>No users as of now!</span>
        </div>
    {:else}
        <table>
            <thead>
                <tr>
                    <th>Email</th>
                    <th>Email Alias</th>
                    <th>First Name</th>
                    <th>Last Name</th>
                    <th>Phone Number</th>
                    <th>Home Address</th>
                    <th>Role</th>
                    <th>Must Change Password</th>
                    <th>Remove</th>
                </tr>
            </thead>
            {#each users as user}
                <tbody>
                    <tr>
                        <td>
                            <a href={"edit/" + user.Email}>
                                <span>{user.Email}</span>
                            </a>
                        </td>
                        <td>
                            <span>{user.EmailAlias}</span>
                        </td>
                        <td>
                            <span>{user.FirstName}</span>
                        </td>
                        <td>
                            <span>{user.LastName}</span>
                        </td>
                        <td>
                            <span>{user.PhoneNumber}</span>
                        </td>
                        <td>
                            <span>{user.HomeAddress}</span>
                        </td>
                        <td>
                            <span>{user.Role}</span>
                        </td>
                        <td>
                            <span>{user.MustChangePW}</span>
                        </td>
                        <td>
                            <button on:click={() => removeAccount(user)}
                                >Remove</button
                            >
                        </td>
                    </tr>
                </tbody>
            {/each}
        </table>
    {/if}
</div>
