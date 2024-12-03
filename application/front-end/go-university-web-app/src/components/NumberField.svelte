<script>
    import "../styles/components/numberField.css";
    export let fieldLabel;
    export let onChangeNumberField;
    export let currentValue;
    export let inputID;
    export let isDisabled = false;
    export let minNumber;
    export let maxNumber;
</script>

<div class="numberFieldContainer">
    <span>{fieldLabel}</span>
    <input
        id={inputID}
        value={currentValue}
        type="number"
        disabled={isDisabled}
        min={minNumber}
        max={maxNumber}
        on:change={onChangeNumberField}
    />
</div>
