<script>
    import "../styles/components/datePicker.css";
    export let fieldLabel;
    export let onChangeDateTimePicker;
    export let currentValue;
    export let inputID;
    export let isDisabled = false;
</script>

<div class="datePickerContainer">
    <span>{fieldLabel}</span>
    <input
        id={inputID}
        value={currentValue}
        type="datetime-local"
        disabled={isDisabled}
        on:change={onChangeDateTimePicker}
    />
</div>
