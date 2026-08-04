<script setup>
import { onMounted } from 'vue'
import L from 'leaflet'
import 'leaflet/dist/leaflet.css'

onMounted(() => {
    const map = L.map('map').setView([-6.2, 106.8], 11)
    const corner1 = L.latLng(-6.074000, 106.682000)
    const corner2 = L.latLng(-6.398333, 106.971667)
    const bounds = L.latLngBounds(corner1, corner2)

    map.setMaxBounds(bounds)
    L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
        attribution: '© OpenStreetMap contributors'
    }).addTo(map)
})
</script>

<template>
    <div id="map"></div>
</template>

<style scoped>
#map {
    width: 100%;
    height: 100vh;
}
</style>
