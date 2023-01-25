<script>
    import Header from "./components/header/Header.svelte"

    import InputText from "./common/input_text/InputText.svelte";
    import Button from "./common/button/Button.svelte";

    import { ToolID } from "./toolData"

    let toolID = new ToolID(  )
    const getToolID = async( ) => {
        await fetch("/plummetron/tool-id")
        .then( d => d.text( ) )
        .then( d => {
            console.log( d )
            toolID = new ToolID( JSON.parse( d ).Data )
            console.log( JSON.stringify( toolID, null, 2 ))
        } )
    }

</script>

<main class="flx-col">

    <Header/>

    <div class="flx-row" id="content">

        <div class="flx-col" id="left-pane">
            <h3>Click Stuff</h3>
            <Button on:click={ getToolID } cls="bg-aqua" enabled={ true }>
                Get Tool ID
            </Button>

            { #if toolID }
                <InputText
                    bind:txt= { toolID.ID }
                    lbl="Tool ID:"
                    enabled={ false }
                />

                <InputText
                    bind:txt= { toolID.SubID }
                    lbl="Tool SubID:"
                    enabled={ false }
                />

                <InputText
                    bind:txt= { toolID.MemCount }
                    lbl="Memory Count:"
                    enabled={ false }
                />

                <InputText
                    bind:txt= { toolID.MemType }
                    lbl="Memory Type Code:"
                    enabled={ false }
                />

                <InputText
                    bind:txt= { toolID.MemCapacity }
                    lbl="Memory Capacity:"
                    enabled={ false }
                />

                <InputText
                    bind:txt= { toolID.FWYear }
                    lbl="Firmware Year:"
                    enabled={ false }
                />

                <InputText
                    bind:txt= { toolID.FWWeek }
                    lbl="Firmware Week:"
                    enabled={ false }
                />
            { /if }
        </div>

        <div class="flx-col" id="right-pane">


        </div>

    </div>

</main>

<style>

    main {
        background-color: var(--dark);
        color: var(--accent_a);
        padding-bottom: 1rem;
        height: 100vh;
    }
    #content {
        padding: 0 1rem;
        overflow: hidden;
        height: 100%;
    }
    #left-pane {
        background-color: var(--light_aa);
        padding: 1rem;
        min-width: 17rem;
        width: 17rem;
        height: 100%
    }
    #right-pane {
        background-color: var(--light_aa);
        padding: 1rem;
        overflow: hidden;
        height: 100%;
    }
</style>