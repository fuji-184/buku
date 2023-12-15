<script>
  import { onMount } from 'svelte';

  const stars = [];

  onMount(() => {
    const canvas = document.getElementById('stars-canvas');
    const context = canvas.getContext('2d');

    canvas.width = window.innerWidth;
    canvas.height = window.innerHeight;

    createStars();

    function createStars() {
      for (let i = 0; i < 100; i++) {
        stars.push({
          x: Math.random() * canvas.width,
          y: Math.random() * canvas.height,
          size: Math.random() * 3 + 1,
          color: getRandomColor(),
          blinkRate: Math.random() * 2 + 1,
          alpha: Math.random(),
        });
      }
    }

    function getRandomColor() {
      const letters = '0123456789ABCDEF';
      let color = '#';
      for (let i = 0; i < 6; i++) {
        color += letters[Math.floor(Math.random() * 16)];
      }
      return color;
    }

    function drawStar(x, y, size, color, alpha) {
      context.beginPath();
      context.moveTo(x, y - size / 2); // Puncak bintang

      // Menggambar segitiga pertama
      for (let i = 1; i <= 4; i++) {
        const angle = ((i * 144 + 90) * Math.PI) / 180;
        const xPoint = x + Math.cos(angle) * size;
        const yPoint = y - Math.sin(angle) * size;
        context.lineTo(xPoint, yPoint);
      }

      // Menggambar segitiga kedua (menutupi bagian tengah)
      for (let i = 1; i <= 4; i++) {
        const angle = ((i * 144 + 90 + 72) * Math.PI) / 180;
        const xPoint = x + Math.cos(angle) * (size / 2);
        const yPoint = y - Math.sin(angle) * (size / 2);
        context.lineTo(xPoint, yPoint);
      }

      context.closePath();
      context.fillStyle = `${color}${Math.floor(alpha * 255).toString(16)}`;
      context.fill();
    }

    function animateStars() {
      context.clearRect(0, 0, canvas.width, canvas.height);

      stars.forEach((star) => {
        drawStar(star.x, star.y, star.size, star.color, star.alpha);

        // Mengatur kelap-kelip bintang
        star.alpha += star.blinkRate / 60;
        if (star.alpha > 1 || star.alpha < 0) {
          star.blinkRate = -star.blinkRate;
        }
      });

      requestAnimationFrame(animateStars);
    }

    animateStars();

    window.addEventListener('resize', () => {
      canvas.width = window.innerWidth;
      canvas.height = window.innerHeight;
      createStars();
    });
  });
</script>

<style>
  canvas {
    position: absolute;
    top: 0;
    left: 0;
    z-index: -1;
  }
</style>

<canvas id="stars-canvas"></canvas>
