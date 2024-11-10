<script lang="ts">
    import AddCourse from "../../../components/AddCourse.svelte";
    import Navigation from "../../../components/Navigation.svelte";
    import type IRole from "../../../interfaces/IRole";

    let role: IRole = {
        isAdmin: true,
        isProfessor: false,
        isStudent: false,
    };
</script>

<Navigation {role} />

<AddCourse />
