export function useNominatim() {
    const reverseGeocode = async (lat, lon) => {
        try {
            const response = await fetch(
                `https://nominatim.openstreetmap.org/reverse?lat=${lat}&lon=${lon}&format=json`
            )
            const data = await response.json()
            return data.address?.road ?? null
        } catch (error) {
            console.error('Error reverse geocoding:', error)
            return null
        }
    }

    return { reverseGeocode }
}
