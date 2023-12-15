export async function load({params}){
        let data = []
	const res = await
	fetch(`http://localhost:3000/laporan/${params.slug}`)
	if (res.status == 200){
	        data = res.json()
	}
	
	return {
	        status: res.status,
	        laporan: data
	}
}