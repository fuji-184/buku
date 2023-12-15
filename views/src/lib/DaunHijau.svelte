<script>
  import { onMount } from 'svelte';

  onMount(() => {
    const canvas = document.getElementById('leaves-canvas');
    const context = canvas.getContext('2d');
    const leaves = [];

    canvas.width = window.innerWidth;
    canvas.height = window.innerHeight;

    createLeaves();

    function createLeaves() {
      for (let i = 0; i < 10; i++) {
        leaves.push({
          x: Math.random() * canvas.width,
          y: Math.random() * canvas.height,
          size: 30,
          rotation: Math.random() * 360,
          swayAmplitude: 20 + Math.random() * 20,
          swayFrequency: 0.02 + Math.random() * 0.02,
        });
      }
    }

    function drawLeaf(x, y, size, rotation) {
      context.save();
      context.translate(x, y);
      context.rotate((rotation * Math.PI) / 180);

      context.beginPath();
      context.moveTo(0, 0);
      context.quadraticCurveTo(size / 2, -size / 4, 0, -size / 2);
      context.quadraticCurveTo(-size / 2, -size / 4, 0, 0);
      context.fillStyle = '#4caf50'; // Warna hijau daun
      context.fill();

      context.restore();
    }

    function animateLeaves() {
      context.clearRect(0, 0, canvas.width, canvas.height);

      leaves.forEach((leaf) => {
        leaf.y -= 1; // Pergerakan vertikal ke atas

        // Efek sway atau bergoyang
        leaf.x += leaf.swayAmplitude * Math.sin(leaf.swayFrequency * leaf.y);

        drawLeaf(leaf.x, leaf.y, leaf.size, leaf.rotation);

        // Reset daun jika di luar layar
        if (leaf.y + leaf.size < 0) {
          leaf.y = canvas.height + leaf.size;
          leaf.x = Math.random() * canvas.width;
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
