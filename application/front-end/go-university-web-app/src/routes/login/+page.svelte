<script lang="ts">
    import Navigation from "../../components/Navigation.svelte";
    import TextField from "../../components/TextField.svelte";
    import type ILogin from "../../interfaces/ILogin";
    import type IRole from "../../interfaces/IRole";
    import loginUser from "../../services/users/loginUser";
    import "../../styles/common.css";

    let role: IRole = {
        isAdmin: true,
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

    function handleTextChange(event: any) {
        login[event.target.id as keyof ILogin] = event.target.value;
    }

    async function submitLogin() {
        isLoading = true;
        const submitLoginResponse = await loginUser(login);
        isSuccessful = submitLoginResponse.isSuccessful;
        if (!isSuccessful) {
            errorMessage = submitLoginResponse.errorMessage;
        } else {
            errorMessage = "";
        }
        isLoading = false;
    }
</script>

<Navigation {role} />

<div>
    <TextField
        fieldLabel="Email"
        currentValue={login.Email}
        onChangeTextField={handleTextChange}
        inputID="Email"
    />
    <TextField
        fieldLabel="Password"
        currentValue={login.Password}
        onChangeTextField={handleTextChange}
        inputID="Password"
    />
    <div class="actionsContainer">
        <a href="/courses">Cancel</a>
        <button on:click={submitLogin}>Login</button>
    </div>
</div>
