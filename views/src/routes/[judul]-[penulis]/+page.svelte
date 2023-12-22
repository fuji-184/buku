<script>
        import Card from "$lib/Card.svelte"
        
        export let data

</script>

<div class="container">
        {#if data.status === 200}
        {#each data.buku as b}
        <Card nama={b.nama} gambar={b.gambar} harga={b.harga} rating={b.rating}/>
        {/each}
        {/if}
        
        {#if data.status === 404}
        <div>Maaf Buku Tidak Ditemukan</div>
        {/if}
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