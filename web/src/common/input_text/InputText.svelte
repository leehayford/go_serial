<script>

    export let txt = null
    export let place = ""
    export let lbl = null
    export let enabled = false

    
</script>

<div class="flx-col" id="container">

    { #if lbl != null }
        <label for="txt-in" class="accent" id="lbl">
            { lbl }
        </label>
    { /if }

    <input 
        id="txt-in"
        type="text" 
        bind:value={ txt }
        spellcheck=value
        placeholder= { place }
        disabled = { !enabled }
    >

</div>

<style>

    #container {
        gap: 0.25rem;
    }

    #lbl {
        font-size: 0.9rem;
    }

    input {
        color: var(--light);
        background-color: var(--dark);
        padding: 0.25rem 0.5rem;
        border-radius: 0.2rem;
		border: 0.1rem solid var(--light_a);
        width: 100%;
    }

    input:disabled {
        color: var(--grey);
		border: 0.1rem solid transparent;
    }

</style>