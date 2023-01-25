<script>
    export let cls = 'bg-accent'
    export let type = 'button'
    export let enabled = true
</script>

<button 
    class={ cls } 
    type={ type }
    disabled = { !enabled }
    on:click
>
    <slot></slot>
</button>

<style>
	button {
		cursor: pointer;
        margin: 0 0.25rem;
        padding: 0.25rem 0.5rem;
        border-radius: 0.2rem;
        border: 0.1rem solid transparent;
        width: 100%;
	}
</style>