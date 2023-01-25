<script>

    /* The Menu components exoprts two variables. */

    export let bg = 'bg-accent'; 
    /* The string name of a css class 
            Defines the background-color and hover background-color
            See ~/public/global.css 
    */

    export let ops; 
    /* An array of options objects:  { Op, Args, Label, Tip }
            
        Op: --> The function to call when this menu options is clicked.
            Provide ONLY the name of the function.
                func NOT func( ) 
            Use anonymous functions to handle return values.
                ( ) => { x = func( ) } 
            
        Args: --> An array of arguments to be passed to Op. 
            Mixing data types is OK...
                [ 1000, "Yermom" ]
            Pass an empty array when Op takes no arguments.
                [ ]  
            
        Label: -->The name of this menu option.
    
        Tip: --> A hint displayed when users hover over this menu option.
*/
</script>

<div class="exp">
    
    <div class="icon { bg }">
        <!-- <img src="/img/menu.svg" alt="menu icon" /> -->
        <svg width="1.8rem" height="1.8rem" viewBox="0 0 6 6" version="1.1" id="menu_svg">
            <g  transform="translate(0,-291)">
                <path class="menu-svg"
                        d=" 
                        M 0 0 
                            L 0 27 
                            L 27 27 
                            L 27 0 
                            L 0 0 
                        z 
                        M 4.5 4 
                            L 18.5 4 
                                C 19.331 4 20 4.669 20 5.5 
                                C 20 6.331 19.331 7 18.5 7 
                            L 4.5 7 
                                C 3.669 7 3 6.331 3 5.5 
                                C 3 4.669 3.669 4 4.5 4 
                        z 
                        M 4.5 10 
                            L 18.5 10 
                                C 19.331 10 20 10.669 20 11.5 
                                C 20 12.331 19.331 13 18.5 13 
                            L 4.5 13 
                                C 3.669 13 3 12.331 3 11.5 
                                C 3 10.669 3.669 10 4.5 10 
                        z 
                        M 4.5 16 
                            L 18.5 16 
                                C 19.331 16 20 16.669 20 17.5 
                                C 20 18.331 19.331 19 18.5 19 
                            L 4.5 19 
                                C 3.669 19 3 18.331 3 17.5 
                                C 3 16.669 3.669 16 4.5 16 
                        z 
                        "
                        transform="matrix(0.27,0,0,0.27,0,291)"
                    />
            </g>
        </svg>
    </div>

    <div class="ops">
        { #each ops as op ( op.Label ) }

            <div class="op" on:click={ ( ) => { op.Op( ...op.Args ) } } on:keyup >
                { op.Label }
                { #if op.Tip != null }
                    <div class="tip">{ op.Tip }</div>
                { /if }
            </div>

        {/each}
    </div>

</div>

<style>
    .icon {
        display: flex;
        flex-direction: column;
        align-items: flex-end;
        justify-content: flex-end;
        height: 1.7rem;
        width: 1.7rem;
        border-radius: 0.19rem;
    }
    .menu-svg {
        fill:var(--dark);
        fill-opacity:1;
        stroke:none;
        stroke-width:6;
        stroke-miterlimit:4;
        stroke-dasharray:none;
        stroke-opacity:1
    }
    .exp {
        display: flex;
        flex-direction: column;
        position: relative;
		cursor: pointer;
    }
    .exp:hover .ops {
        display: block;
        left: 0.8rem;
        top: 0.9rem;
        z-index: 1000;
    }
    .ops {
        display: none;
        flex-direction: column; 
        position: absolute;
        padding: 0.23rem;
        border: 1px solid var(--light_a);
        border-radius: 0.19rem;
        background-color: var(--dark);
        z-index: 1000;
    }    
    .op {
        white-space: nowrap;
        border-radius: 0.19rem;
        color: var(--light_a);
        z-index: 1000;
    }
    .op:hover {
        color: var(--light);
    }
    .op .tip {
        visibility: hidden;
        background-color: var(--grey);
        color: var(--dark);
        text-align: left;
        position: absolute;
        left: 23px;
        height: 20px;
        padding: 0 8px;
        border-radius: 10px;
    }
    .op:hover .tip {
        visibility: visible;
    }


</style>

