<script lang="ts">
    import Navigation from "../../components/Navigation.svelte";
    import type ILogin from "../../interfaces/ILogin";
    import type IRole from "../../interfaces/IRole";
    import loginUser from "../../services/users/loginUser";
    import "../../styles/common.css";
    import "../../styles/items.css";
    import setCookie from "../../utilities/setCookie";

    let role: IRole = {
        isAdmin: false,
        isProfessor: false,
        isStudent: false,
    };

    const login: ILogin = {
        Email: "",
        Password: "",
    };

    let isLoading = false;
    let errorMessage = "";
    let isSuccessful = false;

    setCookie("token", "", 0);
</script>

<Navigation {role} />

<div id="itemForm">
    <div class="successContainer">
        <span>You have logged out successfully!</span>
    </div>
</div>
