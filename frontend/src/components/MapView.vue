<script setup>
import { onMounted } from 'vue'
import L from 'leaflet'
import 'leaflet/dist/leaflet.css'
import { useFloodReports } from '../composables/useFloodReports'

const { fetchAll } = useFloodReports()

const greenIcon = new L.Icon({
    iconUrl: 'https://raw.githubusercontent.com/pointhi/leaflet-color-markers/master/img/marker-icon-2x-green.png',
    shadowUrl: 'https://cdnjs.cloudflare.com/ajax/libs/leaflet/0.7.7/images/marker-shadow.png',
    iconSize: [25, 41],
    iconAnchor: [12, 41],
    popupAnchor: [1, -34],
    shadowSize: [41, 41]
})

onMounted(() => {
    const map = L.map('map').setView([-6.2, 106.8], 11)
    const corner1 = L.latLng(-6.074000, 106.682000)
    const corner2 = L.latLng(-6.398333, 106.971667)
    const bounds = L.latLngBounds(corner1, corner2)

    fetchAll().then(reports => {
        reports.forEach(report => {
            L.marker([report.latitude, report.longitude], { icon: greenIcon }).addTo(map)
        })
    })

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
