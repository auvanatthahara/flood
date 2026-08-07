<script setup>
defineProps({
    report: Object,
})

defineEmits(['close'])

function formatDate(isoString) {
    if (!isoString) return 'N/A'
    return new Date(isoString).toLocaleString('id-ID', {
        day: 'numeric',
        month: 'short',
        year: 'numeric',
        hour: '2-digit',
        minute: '2-digit',
    })
}

function depthClass(depth) {
    if (depth == null) return ''
    if (depth >= 150) return 'depth-critical'
    if (depth >= 70) return 'depth-high'
    if (depth >= 30) return 'depth-medium'
    return 'depth-low'
}
</script>

<template>
    <div class="sidebar" v-if="report">
        <div class="sidebar-header">
            <h2>Flood Report</h2>
            <button class="close-btn" @click="$emit('close')" aria-label="Close">✕</button>
        </div>

        <div class="field">
            <span class="label">Street</span>
            <span>{{ report.street ?? 'N/A' }}</span>
        </div>
        <div class="field">
            <span class="label">Flood Depth</span>
            <span class="depth" :class="depthClass(report.flood_depth)">
                {{ report.flood_depth != null ? report.flood_depth + ' cm' : 'N/A' }}
            </span>
        </div>
        <div class="field">
            <span class="label">Reported</span>
            <span>{{ formatDate(report.created_at) }}</span>
        </div>
        <div class="field">
            <span class="label">Description</span>
            <span class="description">{{ report.raw_text ?? 'No description' }}</span>
        </div>

        <div class="sidebar-footer">
            Data sourced from <a href="https://petabencana.id" target="_blank">PetaBencana.id</a>
        </div>
    </div>
</template>

<style scoped>
.sidebar {
    position: absolute;
    top: 0;
    right: 0;
    width: 300px;
    height: 100vh;
    background: linear-gradient(180deg, #0f2942 0%, #1a4b78 100%);
    color: #fff;
    z-index: 1000;
    padding: 20px;
    box-shadow: -2px 0 16px rgba(0, 0, 0, 0.3);
    box-sizing: border-box;
    overflow-y: auto;
    display: flex;
    flex-direction: column;
    gap: 16px;
}

.sidebar-header {
    display: flex;
    align-items: center;
    justify-content: space-between;
}

h2 {
    margin: 0;
    font-size: 0.95rem;
    font-weight: 600;
    color: rgba(255, 255, 255, 0.6);
    text-transform: uppercase;
    letter-spacing: 0.06em;
}

.close-btn {
    background: none;
    border: none;
    font-size: 1rem;
    color: rgba(255, 255, 255, 0.45);
    cursor: pointer;
    padding: 4px 6px;
    border-radius: 4px;
    line-height: 1;
    transition: color 0.1s, background 0.1s;
}

.close-btn:hover {
    color: #fff;
    background: rgba(255, 255, 255, 0.1);
}

.field {
    display: flex;
    flex-direction: column;
    gap: 4px;
    padding-bottom: 16px;
    border-bottom: 1px solid rgba(255, 255, 255, 0.1);
}

.field:last-of-type {
    border-bottom: none;
    padding-bottom: 0;
}

.label {
    font-size: 0.7rem;
    font-weight: 600;
    color: rgba(255, 255, 255, 0.45);
    text-transform: uppercase;
    letter-spacing: 0.08em;
}

.depth {
    font-size: 1.1rem;
    font-weight: 600;
}

.depth-low      { color: #4ade80; }
.depth-medium   { color: #fbbf24; }
.depth-high     { color: #f87171; }
.depth-critical { color: #fca5a5; }

.description {
    font-size: 0.9rem;
    color: rgba(255, 255, 255, 0.75);
    line-height: 1.5;
}

.sidebar-footer {
    margin-top: auto;
    padding-top: 16px;
    font-size: 0.75rem;
    color: rgba(255, 255, 255, 0.35);
    border-top: 1px solid rgba(255, 255, 255, 0.1);
}

.sidebar-footer a {
    color: rgba(255, 255, 255, 0.5);
    text-decoration: none;
}

.sidebar-footer a:hover {
    color: rgba(255, 255, 255, 0.85);
}
</style>
