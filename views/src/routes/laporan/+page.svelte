<script>
        import DaunHijau from "$lib/DaunHijau.svelte"
        
        let tanggal = ""
        
        function submit(){
                console.log("Submit")
                window.history.replaceState(history.state, '',
                `laporan/${tanggal}`)
        }
        
        document.body.style.backgroundImage = "linear-gradient(315deg, #1dd1a1 0%, #d3d3d3 74%)"


</script>

<DaunHijau />   

<div class="flex justify-center h-screen items-center">
        <form class="w-80 flex justify-center flex-col gap-2" on:submit={submit}>
                                
                <label class="label">
                        <center><span>Tanggal</span></center>
                	<input class="input" title="Input (date)" type="date" placeholder="Klik Untuk
                	Memilih Tanggal" bind:value={tanggal} />
                </label>
				
                <center><button type="submit" class="w-[100px] btn
                variant-filled">Submit</button></center>
        </form>
</div>