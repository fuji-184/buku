export async function load({ fetch, params }){
        
        let { judul, penulis } = params
        
        let data = []
        
	const res = await
	fetch(`http://localhost:3000/buku/${judul}-${penulis}`)
	if (res.status == 200){
	        data = await res.json()
	}
	
	console.log(res)
	console.log(data)
	
	return {
	        status: res.status,
	        buku: data
	}
}