<script lang='ts'>
        export let data
        
</script>

<div class="flex justify-center h-screen items-center">
<table class="w-[70%] bg-black rounded">
        {#if data.status === 200}
        <thead>
                <tr>
                        <td class="p-5 f">No</td>
                        <td class="p-5 f">Isi</td>
                        <td class="p-5 f">Tanggal</td>
                </tr>
        </thead>
        <tbody>
        {#each data.laporan as lp, i }
                <tr>
                        <td class="p-5">{i+1}</td>
                        <td class="p-5">{lp.isi}</td>
                        <td class="p-5">{lp.tanggal}</td>
                </tr>
        {/each}
        </tbody>
        {/if}
        {#if data.status === 404}
        <thead>
                <tr>
                        <th class="p-5 f">Laporan Tidak Tersedia</th>
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
</style>