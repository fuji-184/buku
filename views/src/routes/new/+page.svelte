<script>
        import Card from "$lib/Card.svelte"
        import { onMount } from "svelte"
        
        let buku = []
        
        onMount(async () => {
                try {
                      buku = await (await
                      fetch("http://localhost:3000/bukuterbaru")).json();
                      console.log(buku)
                } catch (error){
                      console.log(error)  
                }
        })

</script>

<div class="container">
        {#each buku as b}
        <Card nama={b.nama} gambar={b.gambar} tahun={b.tahun} rating={b.rating}
        harga={b.harga}/>
        {/each}
</div>

<style>
        .container {
                width: 100%;
                display: grid;
                justify-content: center;
                align-items: center;
                gap: 20px;
                margin-bottom: 100px;
        }
        .etc {
                display: flex;
        }
        
        @media only screen and (max-width: 768px) {
               .container {
                       grid-template-columns: auto;
               }
        }
        
        @media only screen and (min-width: 769px) {
                .container {
                        grid-template-columns: auto auto;
                }
        }
</style>