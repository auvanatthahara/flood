<script setup>
import { ref } from 'vue'
import MapView from './components/MapView.vue'

const showMap = ref(false)
</script>

<template>
  <div v-if="!showMap" class="landing">
    <div class="landing-content">
      <div class="badge">DKI Jakarta</div>
      <h1>Jakarta Flood History</h1>
      <p>Explore historical flood reports across DKI Jakarta — where floods happen, how deep, and when.</p>
      <button class="cta-btn" @click="showMap = true">View Map</button>
    </div>
    <footer class="landing-footer">
      &copy; 2026 Auvan Atthahara
    </footer>
  </div>

  <MapView v-else />
</template>

<style scoped>
.landing {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  min-height: 100vh;
  background: linear-gradient(160deg, #0f2942 0%, #1a4b78 60%, #1e6fa8 100%);
  color: #fff;
  padding: 24px;
  text-align: center;
}

.landing-content {
  max-width: 560px;
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 20px;
}

.badge {
  display: inline-block;
  padding: 4px 12px;
  border: 1px solid rgba(255, 255, 255, 0.35);
  border-radius: 999px;
  font-size: 0.75rem;
  letter-spacing: 0.08em;
  text-transform: uppercase;
  color: rgba(255, 255, 255, 0.75);
}

h1 {
  font-size: clamp(2rem, 5vw, 3rem);
  font-weight: 700;
  margin: 0;
  line-height: 1.15;
  letter-spacing: -0.02em;
}

p {
  margin: 0;
  font-size: 1.05rem;
  color: rgba(255, 255, 255, 0.7);
  line-height: 1.6;
}

.cta-btn {
  margin-top: 8px;
  padding: 12px 32px;
  font-size: 1rem;
  font-weight: 600;
  border: none;
  border-radius: 8px;
  background: #f97316;
  color: #fff;
  cursor: pointer;
  transition: background 0.15s, transform 0.1s;
}

.cta-btn:hover {
  background: #ea6c0a;
  transform: translateY(-1px);
}

.cta-btn:active {
  transform: translateY(0);
}

.landing-footer {
  position: absolute;
  bottom: 20px;
  font-size: 0.8rem;
  color: rgba(255, 255, 255, 0.4);
}

.landing-footer a {
  color: rgba(255, 255, 255, 0.6);
  text-decoration: none;
}

.landing-footer a:hover {
  color: rgba(255, 255, 255, 0.9);
}
</style>
