import { API_URL } from "$env/static/private";
import { fail } from "@sveltejs/kit";
import type { Provincia, Region } from "$lib/server/models.js";

export async function load() {
    const data: { regiones: Region[], provincias: Provincia[] } = {
        regiones: [],
        provincias: []
    }
    let req = await fetch(`${API_URL}/region`);
    if (!req.ok) {
        return data;
    }
    let res = await req.json();
    data.regiones = res;
    req = await fetch(`${API_URL}/provincia`);
    if (!req.ok) {
        return data;
    }
    res = await req.json();
    data.provincias = res;
    return data;
}

export const actions = {
    insert: async ({ request }) => {
        const data = await request.formData();
        const nombre = data.get("nombre");
        let id_region = data.get("region")?.valueOf();
        const form = {
            nombre: nombre,
            id_region: id_region,
            error: ""
        };

        if (nombre === null || nombre === "") {
            form.error = "Ingrese un nombre";
            return fail(400, form);
        }
        if (typeof id_region !== "string") {
            form.error = "Seleccione una región";
            return fail(400, form);
        }
        if (id_region === "") {
            form.error = "Seleccione una región";
            return fail(400, form);
        }
        id_region = parseInt(id_region);

        const req = await fetch(`${API_URL}/provincia`, {
            method: "POST",
            body: JSON.stringify({ nombre, id_region }),
            headers: { "Content-Type": "application/json" }
        });

        if (!req.ok) {
            const res = await req.json();
            form.error = res["error"];
            return fail(400, form);
        }
    }
}
