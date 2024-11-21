<script>
    import "../styles/components/checkbox.css";
    export let fieldLabel;
    export let onChangeCheckbox;
    export let currentValue;
    export let inputID;
    export let isDisabled = false;
</script>

<div class="checkboxContainer">
    <span>{fieldLabel}</span>
    <input
        id={inputID}
        checked={currentValue}
        type="checkbox"
        disabled={isDisabled}
        on:keyup={onChangeCheckbox}
        on:change={onChangeCheckbox}
    />
</div>
