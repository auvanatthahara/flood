const API_URL = import.meta.env.VITE_API_URL

export function useFloodReports() {
    const fetchAll = async () => {
        try {
            const response = await fetch(`${API_URL}/events`)
            return await response.json()
        } catch (error) {
            console.error('Error fetching flood reports:', error)
            return []
        }
    }

    const fetchByRegion = async (regionCode) => {
        try {
            const response = await fetch(`${API_URL}/events/region/${regionCode}`)
            return await response.json()
        } catch (error) {
            console.error('Error fetching flood reports by region:', error)
            return []
        }
    }

    return { fetchAll, fetchByRegion }
}
