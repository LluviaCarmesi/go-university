<script lang="ts">
    import Navigation from "../../components/Navigation.svelte";
    import type IRole from "../../interfaces/IRole";

    let role: IRole = {
        isAdmin: true,
        isProfessor: false,
        isStudent: false,
    };
</script>

<Navigation role={role}/>