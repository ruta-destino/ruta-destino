import { API_URL } from "$env/static/private";

export async function load() {
    const response = await fetch(`${API_URL}/region`)
    if (!response.ok) {
        return { regiones: [] }
    }
    const regiones = await response.json()
    return {
        regiones: regiones
    }
}
