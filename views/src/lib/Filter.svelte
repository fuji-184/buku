<script>
        import { onMount } from "svelte"
        let open = false
        
        function click (){
                open = !open
        }
        
        let kategori = []
        
        onMount(async () => {
                try {
                      kategori = await (await fetch("http://localhost:3000/kategori")).json();  
                } catch (error){
                      console.log(error)  
                }
        })
        
        let ktg = 0
        let min = 0
        let max = 0
        
        function submit (){
              open = !open  
        }
        
</script>

<button class="open" on:click={click}>Filter</button>

{#if open}
        <form class="container">
                <span class="input">
                        <label>Kategori</label>
                <select bind:value={ktg}>
                  <option value=""></option>
                  {#each kategori as { id, nama }}
                    <option value={id}>{nama}</option>
                  {/each}
                </select></span>
                
                <span class="input">
                <label>Harga Minimum</label>
                <input bind:value={min} type="number" placeholder="min">
                <span class="input">
                        </span>
                <span class="input">
                <label>Harga Maximum</label>
                <input bind:value={max} type="number" placeholder="max">
                <span class="input">
                        </span>

                        <a href={ktg + "-" + min + "-" + max}>
                                <button on:click={submit}>Submit</button>
                        </a>
                        <button on:click={click}>Cancel</button>

                
        </form>
{/if}

<style>
        button {
                border-radius: 5px;
                padding: 10px;
                border: none;
                color: white;
                background-color: #222;
        }
        .container {
                position: absolute;
                top: 0;
                bottom: 0;
                left: 0;
                right: 0;
                background-color: rgba(0,0,0, 0.9);
                color: #999;
                z-index: 20;
                display: flex;
                justify-content: center;
                align-items: center;
                flex-direction: column;
                gap: 20px;
        }
        input, select, option {
                border: none;
                outline: none;
                border-radius: 5px;
                padding: 10px;
                background-color: #222;
                color: white;
                width: 100%;
        }
        
        .open {
                background-color: black;
        }
        .open:hover {
                background-color: #222;
                color: white;
        }
        .input {
                display: flex;
                flex-direction: column;
                width: 50%;
                gap: 20px;
        }
</style>