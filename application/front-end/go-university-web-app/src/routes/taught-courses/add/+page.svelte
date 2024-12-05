<script lang="ts">
    import Navigation from "../../../components/Navigation.svelte";
    import TextField from "../../../components/TextField.svelte";
    import type IRole from "../../../interfaces/IRole";
    import "../../../styles/items.css";
    import "../../../styles/common.css";
    import addTaughtCourse from "../../../services/taught-courses/addTaughtCourse";
    import TimePicker from "../../../components/TimePicker.svelte";
    import getCourses from "../../../services/courses/getCourses";
    import type ICourse from "../../../interfaces/ICourse";
    import type ISemester from "../../../interfaces/ISemester";
    import getSemesters from "../../../services/semesters/getSemesters";
    import Dropdown from "../../../components/Dropdown.svelte";
    import createChoices from "../../../utilities/createChoices";
    import { DAY_OPTIONS } from "../../../appSettings";
    import NumberField from "../../../components/NumberField.svelte";
    import { onMount } from "svelte";
    import checkCurrentUser from "../../../services/users/checkCurrentUser";

    let role: IRole = {
        isAdmin: false,
        isProfessor: false,
        isStudent: false,
    };

    let isLoading = false;
    let errorMessage = "";
    let isSuccessful = false;
    let courses: ICourse[] = [];
    let semesters: ISemester[] = [];

    const taughtCourseTextFields: any = {
        ProfessorEmail: "",
        Location: "",
        MaxStudents: "",
        StartTime: "",
        EndTime: "",
    };

    const taughtCourseNumberFields: any = {
        ID: 0,
    };

    const taughtCourseDropdownFields: any = {
        CourseID: "",
        SemesterName: "",
        Day: "",
    };

    function handleTextChange(event: any) {
        taughtCourseTextFields[event.target.id] = event.target.value;
    }

    function handleDropdownChange(event: any) {
        taughtCourseDropdownFields[event.target.id] = event.target.value;
    }

    function handleTimePickerChange(event: any) {
        taughtCourseTextFields[event.target.id] = event.target.value;
    }

    onMount(() => {
        async function getRole() {
            const checkCurrentUserResponse = await checkCurrentUser();
            role = checkCurrentUserResponse;
        }
        getRole();
    });

    async function submitTaughtCourse() {
        isLoading = true;
        const addAppointmentResponse = await addTaughtCourse({
            ID: taughtCourseNumberFields.ID,
            CourseID: taughtCourseDropdownFields.CourseID,
            SemesterName: taughtCourseDropdownFields.SemesterName,
            ProfessorEmail: taughtCourseTextFields.ProfessorEmail,
            MaxStudents: parseInt(taughtCourseTextFields.MaxStudents),
            Location: taughtCourseTextFields.Location,
            Day: taughtCourseDropdownFields.Day,
            StartTime: taughtCourseTextFields.StartTime,
            EndTime: taughtCourseTextFields.EndTime,
        });
        isSuccessful = !addAppointmentResponse.doesErrorExist;
        if (!isSuccessful) {
            errorMessage = addAppointmentResponse.errorMessage;
        } else {
            errorMessage = "";
            taughtCourseDropdownFields.CourseID = "";
            taughtCourseDropdownFields.SemesterName = "";
            taughtCourseTextFields.ProfessorEmail = "";
            taughtCourseTextFields.MaxStudents = "";
            taughtCourseTextFields.Location = "";
            taughtCourseDropdownFields.Day = "";
            taughtCourseTextFields.StartTime = "";
            taughtCourseTextFields.EndTime = "";
        }
        isLoading = false;
    }

    async function getCoursesResponse() {
        const coursesResponse = await getCourses();
        if (coursesResponse.isSuccessful) {
            courses = coursesResponse.courses;
        } else {
            errorMessage = coursesResponse.errorMessage;
        }
        isLoading = false;
    }
    async function getSemestersResponse() {
        const semestersResponse = await getSemesters();
        if (semestersResponse.isSuccessful) {
            semesters = semestersResponse.semesters;
        } else {
            errorMessage = semestersResponse.errorMessage;
        }
        isLoading = false;
    }
    getCoursesResponse();
    getSemestersResponse();
</script>

<Navigation {role} />

<div id="itemForm">
    <div class="mainHeadingContainer">
        <h2>Add an Taught Course</h2>
    </div>
    <div class="descriptionContainer">
        <span>Add a taught course using the form below! </span>
    </div>
    {#if !!errorMessage}
        <div class="errorContainer">
            <span>{errorMessage}</span>
        </div>
    {/if}
    {#if isSuccessful}
        <div class="successContainer">
            <span>Taught Course was added successfully!</span>
        </div>
    {/if}
    {#if !!errorMessage}
        <div>
            <span class="error">{errorMessage}</span>
        </div>
    {/if}
    <Dropdown
        fieldLabel="Course ID"
        currentValue={taughtCourseDropdownFields.CourseID}
        onDropdownChange={handleDropdownChange}
        inputID="CourseID"
        dropdownOptions={createChoices(courses, "ID", "ID")}
    />
    <Dropdown
        fieldLabel="Semester Name"
        currentValue={taughtCourseDropdownFields.SemesterName}
        onDropdownChange={handleDropdownChange}
        inputID="SemesterName"
        dropdownOptions={createChoices(semesters, "Name", "Name")}
    />
    <TextField
        fieldLabel="Professor Email"
        currentValue={taughtCourseTextFields.ProfessorEmail}
        onChangeTextField={handleTextChange}
        inputID="ProfessorEmail"
    />
    <TextField
        fieldLabel="Location"
        currentValue={taughtCourseTextFields.Location}
        onChangeTextField={handleTextChange}
        inputID="Location"
    />
    <NumberField
        fieldLabel="Max Students"
        currentValue={taughtCourseTextFields.MaxStudents}
        onChangeNumberField={handleTextChange}
        inputID="MaxStudents"
        minNumber={5}
        maxNumber={30}
    />
    <Dropdown
        fieldLabel="Day"
        currentValue={taughtCourseDropdownFields.Day}
        onDropdownChange={handleDropdownChange}
        inputID="Day"
        dropdownOptions={DAY_OPTIONS}
    />
    <TimePicker
        fieldLabel="Start Time"
        currentValue={taughtCourseTextFields.StartTime}
        onChangeTimePicker={handleTimePickerChange}
        inputID="StartTime"
    />
    <TimePicker
        fieldLabel="End Time"
        currentValue={taughtCourseTextFields.EndTime}
        onChangeTimePicker={handleTimePickerChange}
        inputID="EndTime"
    />
    <div class="actionsContainer">
        <a href="/taught-courses">Cancel</a>
        <button on:click={submitTaughtCourse}>Submit Taught Course</button>
    </div>
</div>
