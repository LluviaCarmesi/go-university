<script lang="ts">
    import type ICourse from "../interfaces/ICourse";
    import TextField from "./TextField.svelte";

    let courseID = "";
    let courseName = "";
    let courseDescription = "";

    let isLoading = true;
    let errorMessage = "";
    let isSuccessful = false;

    const course: ICourse = {
        courseID: "",
        courseName: "",
        courseDescription: "",
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
        currentValue={courseID}
        onChangeTextField={handleTextChange}
        inputID="courseID"
    />
    <TextField
        fieldLabel="Course Name"
        currentValue={courseName}
        onChangeTextField={handleTextChange}
        inputID="courseName"
    />
    <TextField
        fieldLabel="Course Description"
        currentValue={courseDescription}
        onChangeTextField={handleTextChange}
        inputID="courseDescription"
    />
    <button >Submit Course</button>
    <a href="/courses">Cancel</a>
</div>
