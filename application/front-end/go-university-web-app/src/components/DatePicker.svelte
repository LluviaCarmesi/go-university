<script>
    import "../styles/components/datePicker.css";
    export let fieldLabel;
    export let onChangeDatePicker;
    export let currentValue;
    export let inputID;
    export let isDisabled = false;
</script>

<div class="datePickerContainer">
    <span>{fieldLabel}</span>
    <input
        id={inputID}
        value={currentValue}
        type="date"
        disabled={isDisabled}
        on:change={onChangeDatePicker}
    />
</div>
