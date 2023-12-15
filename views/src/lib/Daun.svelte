<script>
  import { onMount } from 'svelte';

  const leaves = [];

  onMount(() => {
    const canvas = document.getElementById('leaves-canvas');
    const context = canvas.getContext('2d');

    canvas.width = window.innerWidth;
    canvas.height = window.innerHeight;

    createLeaves();

    function createLeaves() {
      for (let i = 0; i < 50; i++) {
        leaves.push({
          x: Math.random() * canvas.width,
          y: Math.random() * canvas.height,
          size: Math.random() * 20 + 5,
          speed: Math.random() * 2 + 1,
          rotation: Math.random() * 360,
        });
      }
    }

    function drawLeaf(x, y, size, rotation) {
      context.save();
      context.translate(x, y);
      context.rotate((rotation * Math.PI) / 180);

      // Gambar bentuk daun
      context.beginPath();
      context.moveTo(0, 0);
      context.quadraticCurveTo(-5, -5, -5, -10);
      context.quadraticCurveTo(-5, -15, 5, -15);
      context.quadraticCurveTo(5, -15, 5, -10);
      context.quadraticCurveTo(5, -5, 0, 0);
      context.fillStyle = '#ff7f50'; // Warna daun
      context.fill();

      context.restore();
    }

    function animateLeaves() {
      context.clearRect(0, 0, canvas.width, canvas.height);

      leaves.forEach((leaf) => {
        drawLeaf(leaf.x, leaf.y, leaf.size, leaf.rotation);

        leaf.y += leaf.speed;

        if (leaf.y > canvas.height) {
          leaf.y = 0;
        }
      });

      requestAnimationFrame(animateLeaves);
    }

    animateLeaves();

    window.addEventListener('resize', () => {
      canvas.width = window.innerWidth;
      canvas.height = window.innerHeight;
      createLeaves();
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

<canvas id="leaves-canvas"></canvas>
