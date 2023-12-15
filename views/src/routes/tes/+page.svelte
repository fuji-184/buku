<script>
  import { onMount } from 'svelte';

  const particles = [];

  onMount(() => {
    const canvas = document.getElementById('particle-canvas');
    const context = canvas.getContext('2d');

    canvas.width = window.innerWidth;
    canvas.height = window.innerHeight;

    createParticles();

    function createParticles() {
      for (let i = 0; i < 100; i++) {
        particles.push({
          x: Math.random() * canvas.width,
          y: Math.random() * canvas.height,
          radius: Math.random() * 4 + 1,
          color: 'rgba(255, 255, 255, 0.6)',
          speedX: Math.random() * 3 - 1.5,
          speedY: Math.random() * 3 - 1.5,
        });
      }
    }

    function animateParticles() {
      context.clearRect(0, 0, canvas.width, canvas.height);

      particles.forEach((particle) => {
        context.beginPath();
        context.arc(particle.x, particle.y, particle.radius, 0, Math.PI * 2);
        context.fillStyle = particle.color;
        context.fill();

        particle.x += particle.speedX;
        particle.y += particle.speedY;

        if (particle.radius > 0.2) particle.radius -= 0.1;
      });

      requestAnimationFrame(animateParticles);
    }

    animateParticles();

    window.addEventListener('resize', () => {
      canvas.width = window.innerWidth;
      canvas.height = window.innerHeight;
      createParticles();
    });
  });
</script>

<style>
  canvas {
    position: absolute;
    top: 0;
    left: 0;
  }
</style>

<canvas id="particle-canvas"></canvas>
