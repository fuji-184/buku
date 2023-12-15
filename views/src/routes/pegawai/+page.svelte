<script>
import Partikel from "$lib/Partikel.svelte"
import Daun from "$lib/Daun.svelte"

        let umur = 0
        let nama = ""
        
        let data = []
        
        let pesan = null
        
        async function cari(){
                try {

              const res = await fetch(`http://localhost:3000/pegawai/${umur}-${nama}`)
              console.log(res)
              
              if (res.status == 200) {
                      data = await res.json()
                      pesan = null
              } else {
                      pesan = "Data Tidak Ditemukan"
                      data = []
              }
              
                } catch (err){
                       console.log(err)
                }
        }
        
document.body.style.backgroundImage = 'linear-gradient(to right, #ff9966, #ff5e62)';
</script>

<Partikel warna="#ff9966" />

<Daun />

<div class="flex justify-center h-screen items-center flex-col">
        <div>
                <label class="label">
                        <span>Umur</span>
                        <input class="input" type="number" on:input={cari}
                        bind:value={umur} />
                </label>
                <label class="label">
                        <span>Nama</span>
                        <input class="input" type="text" on:input={cari}  bind:value={nama} />
                </label>
                </div>
        
        
                <table class="w-[80%] bg-black rounded mt-12">
                        {#if data.length != 0}
                        <thead>
                                <tr>
                                        <th class="p-5 f">No</th>
                                        <th class="p-5 f">Nama</th>
                                        <th class="p-5 f">Umur</th>
                              
                                </tr>
                        </thead>
                        <tbody>

                        {#each data as d, i }
                                <tr>
                                        <td class="p-5">{i+1}</td>
                                        <td class="p-5">{d.nama}</td>
                                        <td class="p-5">{d.umur}</td>
                                </tr>
                        {/each}

                        

                        
                        </tbody>
                        {/if}
                        
                                             {#if pesan != null}
                        <thead>
                        <tr>
                        <th class="p-5 f">{pesan}</th>
                        </tr>
                        </thead>
                        {/if}
                </table>
        
</div>

<style>
   .f{
                font-size: 30px;
                font-weight: 1000;
                background-image: linear-gradient(to right, #ff9966, #ff5e62);
                -webkit-background-clip: text;
                color: transparent;
                
}

.input {
        background-color: black;
        border: none;
}
</style>