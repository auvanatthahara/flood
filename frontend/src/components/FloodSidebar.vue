<script setup>
defineProps({
    report: Object,
})
</script>

<template>
    <div class="sidebar" v-if="report">
        <h2>Flood Report</h2>
        <div class="field">
            <span class="label">Street</span>
            <span>{{ report.street ?? 'N/A' }}</span>
        </div>
        <div class="field">
            <span class="label">Flood Depth</span>
            <span>{{ report.flood_depth != null ? report.flood_depth + ' cm' : 'N/A' }}</span>
        </div>
        <div class="field">
            <span class="label">Description</span>
            <span>{{ report.raw_text ?? 'No description' }}</span>
        </div>
    </div>
</template>

<style scoped>
.sidebar {
    position: absolute;
    top: 0;
    right: 0;
    width: 280px;
    height: 100vh;
    background: white;
    z-index: 1000;
    padding: 16px;
    box-shadow: -2px 0 8px rgba(0, 0, 0, 0.15);
    box-sizing: border-box;
}

h2 {
    margin: 0 0 16px;
    font-size: 1rem;
}

.field {
    display: flex;
    flex-direction: column;
    gap: 4px;
    margin-bottom: 12px;
}

.label {
    font-size: 0.75rem;
    color: #888;
    text-transform: uppercase;
    letter-spacing: 0.05em;
}
</style>
