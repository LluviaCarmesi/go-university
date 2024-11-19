<script>
    import "../styles/components/textField.css";
    export let fieldLabel;
    export let onChangeTextField;
    export let currentValue;
    export let inputID;
    export let isExpandable = false;
    export let isDisabled = false;
</script>

<div class="textFieldContainer">
    <span>{fieldLabel}</span>
    {#if isExpandable}
        <textarea
            id={inputID}
            rows="6"
            cols="30"
            disabled={isDisabled}
            on:keyup={onChangeTextField}
            on:change={onChangeTextField}
            >{currentValue}</textarea>
    {:else}
        <input
            id={inputID}
            value={currentValue}
            type="text"
            disabled={isDisabled}
            on:keyup={onChangeTextField}
            on:change={onChangeTextField}
        />
    {/if}
</div>
