<script lang="ts">
        import Partikel from "$lib/Partikel.svelte"
        import Komet from "$lib/Komet.svelte"
        import { fly } from 'svelte/transition';
        
        document.body.style.backgroundImage = 'radial-gradient(at 50% 0%, rgba(var(--color-secondary-500) / 0.18) 0px, transparent 75%), radial-gradient(at 100% 0%, rgba(var(--color-tertiary-500) / 0.18) 0px, transparent 50%)'
</script>

<Partikel warna="#2196F3" />

<Komet />

<div class="grid grid-cols-1 lg:grid-cols-3p h-screen place-items-center
place-content-center">
        
        <a href="/laporan" transition:fly={{ x:200, duration: 300 }}
        class="menu border border-black bg-black/50 w-full
                        h-[100px] flex justify-center items-center
                        font-comic-sans animasi">
        Laporan
        </a>
        <a href="/pegawai" transition:fly={{ x:200, duration: 300 }}
        class="menu border border-black bg-black/50 w-full
                        h-[100px] flex justify-center items-center
                        font-comic-sans animasi">
        Pegawai
        </a>
                
        <span class="top"/>
        <span class="right"/>
        <span class="bottom"/>
        <span class="left"/>
        
</div>

<style>
        @keyframes border {
                0% {
                    transform: rotate(-90deg);
                    border-color: white;
                    opacity: 0.2;
                    width: 100px;
                }
                100% {
                    transform: rotate(0deg);
                    opacity: 1;
                    width: 100%;
                }
        }
        
        .animasi {
                animation: border 1.1s;
        }
        
        .menu {
                font-size: 30px;
                font-weight: 1000;
                color: #3498db;
                -webkit-text-stroke-width: 0.5px;
                -webkit-text-stroke-color: black;
        }
        
        @keyframes tes {
                0% {
                        border-top: 3px solid white;
                        border-bottom: 3px solid white;
                }
                100% {
                        border-left: 3px solid white;
                        border-right: 3px solid white;
                }
        }
        
        .menu:hover {
                background-color: #3498db;
                color: black;
                -webkit-text-stroke-width: 0.5px;
                -webkit-text-stroke-color: black;
                font-size: 45px;
        }

        .f{
                font-size: 30px;
                font-weight: 1000;
                background-image: linear-gradient(to right, #ff9966, #ff5e62);
                -webkit-background-clip: text;
                color: transparent;
                display: inline-block;
}

.top {
        position: fixed;
  border-radius: 100vmax;
  top: 0;
  left: 0;
  width: 0;
  height: 5px;
  background: linear-gradient(
    90deg,
    transparent 50%,
    rgba(52, 152, 219, 0.5),
    rgb(52, 152, 219)
  );
}

.bottom {
        position: fixed;
  border-radius: 100vmax;
  right: 0;
  bottom: 0;
  height: 5px;
  background: linear-gradient(
    90deg,
    rgba(52, 152, 219, 0.5),
    rgb(52, 152, 219),
    transparent 50%
  );
}

.right {
        position: fixed;
  border-radius: 100vmax;
  top: 0;
  right: 0;
  width: 5px;
  height: 0;
  background: linear-gradient(
    180deg,
    transparent 30%,
    rgba(52, 152, 219, 0.5),
    rgb(52, 152, 219)
  );
}

.left {
        position: fixed;
  border-radius: 100vmax;
  left: 0;
  bottom: 0;
  width: 5px;
  height: 0;
  background: linear-gradient(
    180deg,
    rgba(52, 152, 219, 0.5),
    rgb(52, 152, 219),
    transparent 70%
  );
}

.top {
  animation: animateTop 3s ease-in-out infinite;
}

.bottom {
  animation: animateBottom 3s ease-in-out infinite;
}

.right {
  animation: animateRight 3s ease-in-out infinite;
}

.left {
  animation: animateLeft 3s ease-in-out infinite;
}

@keyframes animateTop {
  25% {
    width: 100%;
    opacity: 1;
  }

  30%,
  100% {
    opacity: 0;
  }
}

@keyframes animateBottom {
  0%,
  50% {
    opacity: 0;
    width: 0;
  }

  75% {
    opacity: 1;
    width: 100%;
  }

  76%,
  100% {
    opacity: 0;
  }
}

@keyframes animateRight {
  0%,
  25% {
    opacity: 0;
    height: 0;
  }

  50% {
    opacity: 1;
    height: 100%;
  }

  55%,
  100% {
    height: 100%;
    opacity: 0;
  }
}

@keyframes animateLeft {
  0%,
  75% {
    opacity: 0;
    bottom: 0;
    height: 0;
  }

  100% {
    opacity: 1;
    height: 100%;
  }
}


    
</style>