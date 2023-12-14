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

type FormKey = "error" | "nombre" | "id_region";
const form: { [key in FormKey]: FormDataEntryValue | null } = {
    error: null, nombre: null, id_region: null
}

export const actions = {
    insert: async ({ request }) => {
        const data = await request.formData();
        const f_nombre = data.get("nombre");
        let f_id_region = data.get("region");
        form.nombre = f_nombre;
        form.id_region = f_id_region;

        if (typeof f_nombre !== "string" || f_nombre === "") {
            form.error = "Ingrese un nombre";
            return fail(400, form);
        }
        const nombre = f_nombre;

        if (typeof f_id_region !== "string" || f_id_region === "") {
            form.error = "Seleccione una región";
            return fail(400, form);
        }
        const id_region = parseInt(f_id_region);
        if (isNaN(id_region)) {
            form.error = "Seleccione una región válida"
            return fail(400, form);
        }

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
    },
    delete: async ({ request }) => {
        const data = await request.formData();
        let id = data.get("id") || "";

        const req = await fetch(`${API_URL}/provincia/${id}`, {
            method: "DELETE",
            headers: { "Content-Type": "application/json" }
        })

        if (!req.ok) {
            const res = await req.json()
            form.error = res["error"]
            return fail(400, form)
        }
    }
}
