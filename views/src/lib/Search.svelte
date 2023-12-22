<script>
        let judul = ""
        let penulis = ""
        let open = false
        
        function click(){
                open = !open
        }
     
</script>

        <button on:click={click} style="z-index: 11">
                {#if !open}
                Cari
                {:else}
                Close
                {/if}
        </button>

        {#if open}
        <span class="cari">
                <span>
                        <input type="text" bind:value={judul} placeholder="Cari berdasarkan judul...">
                        <a href={judul + "-1"} class="btn-search" >
                                <button class="btn-search" on:click={() => open = !open}>Cari</button>
                        </a>
                </span>
                <span>
                        <input type="text" bind:value={penulis} placeholder="Cari berdasarkan penulis...">
                                             <a href={"1-" + penulis} class="btn-search" ><button class="btn-search" on:click={() => open = !open}>Cari</button>
                        </a>
                </span>
        </span>
        {/if}


<style>
        button, input {
                border-radius: 5px;
                padding: 10px;
                
        }
        
        button {
                z-index: 10;
                border: none;
                color: white;
                background-color: black;
        }
        
        .btn-search {
                background-color: #222;
                border: none;
                color: white;
        }
        
        button:hover {
                background-color: #222;
        }
        
        input {
                border: none;
                outline: none;
                background-color: #222;
                text-align: center;
                color: white;
        }
        
        .cari {
                display: flex;
                justify-content: center;
                padding: 0 20px;
                gap: 20px;
                position: absolute;
                top: 0;
                bottom: 0;
                left: 0;
                right: 0;
                align-items: center;
                background-color: rgba(0,0,0, 0.9);
                z-index: 10;
                flex-direction: column;
        }
        a {
                text-decoration: none;
        }
</style>