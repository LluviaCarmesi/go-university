<script lang="ts">
    import type ICourse from "../interfaces/ICourse";
    import addCourse from "../services/addCourse";
    import TextField from "./TextField.svelte";

    let isLoading = true;
    let errorMessage = "";
    let isSuccessful = false;

    const course: ICourse = {
        ID: "",
        Name: "",
        Description: "",
    };

    function handleTextChange(event: any) {
        course[event.target.id as keyof ICourse] = event.target.value;
    }

    async function submitCourse() {
        const addCourseResponse = await addCourse(course);
        isSuccessful = !addCourseResponse.doesErrorExist;
        if (!isSuccessful) {
            errorMessage = addCourseResponse.errorMessage;
        } else {
            errorMessage = "";
        }
    }
</script>

<div>
    <div>
        <h2 class="mainHeading">Add a Video</h2>
    </div>
    <div>
        <span class="description">Add a video using the form below! </span>
    </div>
    {#if !!errorMessage}
        <div>
            <span class="error">{errorMessage}</span>
        </div>
    {/if}
    {#if isSuccessful}
        <div>
            <span class="success">Video was added</span>
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
    />
    <button on:click={submitCourse}>Submit Course</button>
    <a href="/courses">Cancel</a>
</div>
