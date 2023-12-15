<script>
  import { onMount } from 'svelte';

  onMount(() => {
    const canvas = document.getElementById('walker-canvas');
    const context = canvas.getContext('2d');

    const walker = {
      x: 50, // Posisi awal x
      y: 50, // Posisi awal y
      speed: 2, // Kecepatan berjalan
    };

    function drawWalker() {
      // Hapus canvas untuk menggambar kembali
      context.clearRect(0, 0, canvas.width, canvas.height);

      // Gambar kepala
      context.beginPath();
      context.arc(walker.x, walker.y - 20, 10, 0, Math.PI * 2);
      context.fillStyle = '#3498db';
      context.fill();
      context.closePath();

      // Gambar tubuh
      context.fillRect(walker.x - 5, walker.y - 10, 10, 20);

      // Gambar kaki
      context.fillRect(walker.x - 5, walker.y + 10, 5, 10);
      context.fillRect(walker.x, walker.y + 10, 5, 10);

      // Pindahkan posisi walker
      walker.x += walker.speed;

      // Reset posisi jika melewati batas layar
      if (walker.x > canvas.width) {
        walker.x = -10;
      }

      requestAnimationFrame(drawWalker);
    }

    drawWalker();
  });
</script>

<style>
  canvas {
    background-color: #fff; /* Warna latar belakang canvas */
  }
</style>

<canvas id="walker-canvas" width="400" height="200"></canvas>
