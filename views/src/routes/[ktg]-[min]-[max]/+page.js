export async function load({params, fetch}){
        const { ktg, min, max } = params
        
        let res
        let buku = []
        
        try {
                res = await fetch(`http://localhost:3000/search/${ktg}-${min}-${max}`)
        
                if (res.status == 200){
                        buku = await res.json()
                }
        } catch (error){
                console.log(error)
        }
        
        console.log(buku)
        
        return {
                status: res.status,
                buku: buku
        }
}