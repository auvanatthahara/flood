<script setup>
import { ref, onMounted } from 'vue'
import L from 'leaflet'
import 'leaflet/dist/leaflet.css'
import { useFloodReports } from '../composables/useFloodReports'
import { useNominatim } from '../composables/useNominatim'
import FloodSidebar from './FloodSidebar.vue'

const { fetchAll } = useFloodReports()
const { reverseGeocode } = useNominatim()

const selectedReport = ref(null)

const greenIcon = new L.Icon({
    iconUrl: 'https://raw.githubusercontent.com/pointhi/leaflet-color-markers/master/img/marker-icon-2x-green.png',
    shadowUrl: 'https://cdnjs.cloudflare.com/ajax/libs/leaflet/0.7.7/images/marker-shadow.png',
    iconSize: [25, 41],
    iconAnchor: [12, 41],
    popupAnchor: [1, -34],
    shadowSize: [41, 41]
})

const onMarkerClick = async (report) => {
    const street = await reverseGeocode(report.latitude, report.longitude)
    selectedReport.value = { ...report, street }
}

onMounted(() => {
    const map = L.map('map').setView([-6.1754, 106.8272], 12)
    const corner1 = L.latLng(-6.074000, 106.682000)
    const corner2 = L.latLng(-6.398333, 106.971667)
    const bounds = L.latLngBounds(corner1, corner2)

    map.setMaxBounds(bounds)
    map.setMinZoom(12)
    map.setMaxZoom(18)

    fetchAll().then(reports => {
        reports.forEach(report => {
            const marker = L.marker([report.latitude, report.longitude], { icon: greenIcon }).addTo(map)
            marker.on('click', () => onMarkerClick(report))
        })
    })

    L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
        attribution: '© OpenStreetMap contributors'
    }).addTo(map)
})
</script>

<template>
    <div class="map-container">
        <div id="map"></div>
        <FloodSidebar :report="selectedReport" @close="selectedReport = null" />
    </div>
</template>

<style scoped>
.map-container {
    position: relative;
    width: 100%;
    height: 100vh;
}

#map {
    width: 100%;
    height: 100%;
}
</style>
