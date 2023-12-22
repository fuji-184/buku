<script>
        import Star from "./Star.svelte"
        
        export let nama = ""
        export let gambar = ""
        export let rating = 0
        export let tahun = ""
        export let harga = 0
</script>

<div class="card">
        <img src={gambar}/>
        <div class="info">
                <h2>{nama} | {tahun}</h2>
                <p style="margin-top: 20px" class="harga">Rp{harga}</p>
                <button class="beli">Beli</button>
        </div>
        <span class="rating">{rating}</span>
        
</div>

<style>
        .card {
                
                background-color: black;
                color: white;
                border-radius: 5px;
                position: relative;
                width: 100%;
        }
        img {
                
                border-top-left-radius: 5px;
                border-top-right-radius: 5px;
                width: 100%;
                height: 100%;
        }
        .info {
                
                position: absolute;
                bottom: -17px;
                background-color: rgba(0,0,0, 0.5);
                width: 100%;
                border-bottom-left-radius: 5px;
                border-bottom-right-radius: 5px;
                color: #f1f1f1;
                font-family: 'Comic Sans MS';
                
        }
        
        .rating {
                position: absolute;
                padding: 10px;
                border-radius: 50%;
                background-color: rgba(0,0,0, 1);
                top: -10px;
                left: -10px;
                width: 20px;
                height: 20px;
                display: flex;
                justify-content: center;
                align-items: center;
                border: 2px solid #999;
        }
        .harga {
                font-size: 25px;
        }
        .beli {
                padding: 10px;
                border-bottom-left-radius: 5px;
                border-bottom-right-radius: 5px;
                width: 100%;
                background-color: black;
                color: white;
                font-size: 20px;
                font-weight: 700;
                margin-top: 20px;
                border: none;
        }
        h2, p {
                margin: 10px;
        }

</style>