<script lang="ts">
    import Navigation from "../../../components/Navigation.svelte";
    import type ICourse from "../../../interfaces/ICourse";
    import addCourse from "../../../services/courses/addCourse";
    import TextField from "../../../components/TextField.svelte";
    import type IRole from "../../../interfaces/IRole";
    import "../../../styles/items.css";
    import "../../../styles/common.css";
    import checkCurrentUser from "../../../services/users/checkCurrentUser";
    import { onMount } from "svelte";

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

    const course: ICourse = {
        ID: "",
        Name: "",
        Description: "",
        Credits: "",
    };

    function handleTextChange(event: any) {
        course[event.target.id as keyof ICourse] = event.target.value;
    }

    async function submitCourse() {
        isLoading = true;
        const addCourseResponse = await addCourse(course);
        isSuccessful = !addCourseResponse.doesErrorExist;
        if (!isSuccessful) {
            errorMessage = addCourseResponse.errorMessage;
        } else {
            errorMessage = "";
        }
        isLoading = false;
    }
</script>

<Navigation {role} />

<div id="itemForm">
    <div class="mainHeadingContainer">
        <h2>Add a Course</h2>
    </div>
    <div class="descriptionContainer">
        <span>Add a course using the form below! </span>
    </div>
    {#if !!errorMessage}
        <div class="errorContainer">
            <span>{errorMessage}</span>
        </div>
    {/if}
    {#if isSuccessful}
        <div class="successContainer">
            <span>Course was added successfully!</span>
        </div>
    {/if}
    {#if !!errorMessage}
        <div>
            <span class="error">{errorMessage}</span>
        </div>
    {/if}
    <TextField
        fieldLabel="Course ID"
        currentValue={course.ID}
        onChangeTextField={handleTextChange}
        inputID="ID"
    />
    <TextField
        fieldLabel="Course Name"
        currentValue={course.Name}
        onChangeTextField={handleTextChange}
        inputID="Name"
    />
    <TextField
        fieldLabel="Course Description"
        currentValue={course.Description}
        onChangeTextField={handleTextChange}
        inputID="Description"
        isExpandable={true}
    />
    <TextField
        fieldLabel="Course Credits"
        currentValue={course.Credits}
        onChangeTextField={handleTextChange}
        inputID="Credits"
    />
    <div class="actionsContainer">
        <a href="/courses">Cancel</a>
        <button on:click={submitCourse}>Submit Course</button>
    </div>
</div>
