<script setup>
import { ref } from 'vue'
import MapView from './components/MapView.vue'

const showMap = ref(false)
</script>

<template>
  <div v-if="!showMap" class="landing">
    <h1>Jakarta Flood History</h1>
    <p>Explore historical flood reports across DKI Jakarta.</p>
    <button @click="showMap = true">View Map</button>
  </div>

  <MapView v-else />
</template>

<style scoped>
.landing {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  height: 100vh;
  gap: 16px;
  text-align: center;
}

h1 {
  font-size: 2.5rem;
  margin: 0;
}

p {
  color: #666;
  margin: 0;
}

button {
  margin-top: 8px;
  padding: 10px 24px;
  font-size: 1rem;
  cursor: pointer;
}
</style>
