<script lang="ts">
    import Navigation from "../../../../components/Navigation.svelte";
    import type ICourse from "../../../../interfaces/ICourse";
    import TextField from "../../../../components/TextField.svelte";
    import type IRole from "../../../../interfaces/IRole";
    import getCourseByID from "../../../../services/courses/getCourseByID";
    import editCourse from "../../../../services/courses/editCourse";
    import "../../../../styles/items.css";
    import "../../../../styles/common.css";
    export let data;

    let role: IRole = {
        isAdmin: false,
        isProfessor: false,
        isStudent: false,
    };

    let isLoading = false;
    let errorMessage = "";
    let isSuccessful = false;

    let course: ICourse = {
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
        const editCourseResponse = await editCourse(course);
        isSuccessful = !editCourseResponse.doesErrorExist;
        if (!isSuccessful) {
            errorMessage = editCourseResponse.errorMessage;
        } else {
            errorMessage = "";
        }
        isLoading = false;
    }

    async function getCourseByIDResponse() {
        const courseResponse = await getCourseByID(data.id);
        if (courseResponse.isSuccessful) {
            course = courseResponse.course;
        } else {
            errorMessage = courseResponse.errorMessage;
        }
        isLoading = false;
    }
    getCourseByIDResponse();
</script>

<Navigation {role} />

<div id="itemForm">
    <div class="mainHeadingContainer">
        <h2>Edit {course.ID} Course</h2>
    </div>
    <div class="descriptionContainer">
        <span>Edit the course using the form below!</span>
    </div>
    {#if !!errorMessage}
        <div class="errorContainer">
            <span>{errorMessage}</span>
        </div>
    {/if}
    {#if isSuccessful}
        <div class="successContainer">
            <span>Course was modified successfully!</span>
        </div>
    {/if}
    <TextField
        fieldLabel="Course ID"
        currentValue={course.ID}
        onChangeTextField={handleTextChange}
        inputID="ID"
        isDisabled={true}
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
