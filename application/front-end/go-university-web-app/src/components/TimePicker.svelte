<script>
    import "../styles/components/timePicker.css";
    export let fieldLabel;
    export let onChangeTimePicker;
    export let currentValue;
    export let inputID;
    export let isDisabled = false;
</script>

<div class="timePickerContainer">
    <span>{fieldLabel}</span>
    <input
        id={inputID}
        value={currentValue}
        type="time"
        disabled={isDisabled}
        on:change={onChangeTimePicker}
    />
</div>
