<script>

    import { Guid } from 'js-guid'
    import { onMount } from 'svelte'
    
    export let txt = null
    export let place = ""
    export let lbl = null
    export let enabled = false
    export let remMinHeigh = 1.5  
    export let guid = Guid.newGuid( ).StringGuid
    
    let init = null

    onMount( ( ) => { init = true } )
    
    $: { if( init != false & txt != null ) { resize( ) } }
    
    const resize = async( ) => { // console.log( `Resizing textarea: ${ guid }`)
        let txtArea = document.getElementById( guid )
        txtArea.parentNode.dataset.replicatedValue = txt
    }
    
</script>

<div class="flx-col" id="container">

    { #if lbl != null }
        <label for={ guid } class="accent" id="lbl">
            { lbl }
        </label>
    { /if }

	<div class="disp">

        <div class="grow-wrap">

            <textarea id={ guid } 
                type="text" 
                bind:value={ txt }
                spellcheck=value
                placeholder= { place }
                disabled = { !enabled }
                style="min-height: { remMinHeigh }rem;"
                ></textarea>

        </div>

    </div>

</div>

<style>

    #container {    
        gap: 0.25rem;    
    }

    #lbl {
        font-size: 0.9rem;
    }

    textarea {
        color: var(--light);
    	font-size: 0.9rem;
        background-color: var(--dark);
        padding: 0.25rem 0.5rem;
        border-radius: 0.2rem;
		border: 0.1rem solid var(--light_a);
        width: 100%;
        height: 100%;
        outline: none;
		overflow: visible;
        white-space: pre-wrap;
    }
    textarea:disabled {
        color: var(--grey);
		border: 0.1rem solid transparent;        
    }

	.disp {
        background-color: transparent;
		border-radius: 0.5rem;
	}
    .grow-wrap {
        display: grid;
    }
    .grow-wrap::after {
        content: attr(data-replicated-value) " ";
        white-space: pre-wrap;
        visibility: hidden;
    }
    .grow-wrap > textarea {
        resize: none;
        overflow: hidden;
    }
    .grow-wrap > textarea,
    .grow-wrap::after {
        padding: 0.5rem;
        font: inherit;
        grid-area: 1 / 1 / 2 / 2;
    }

</style>
