import { API_URL } from "$env/static/private";
import { error, fail } from "@sveltejs/kit";
import type { Ciudad, Provincia } from "$lib/server/models.js";

export async function load() {
    let req = await fetch(`${API_URL}/provincia`);
    if (!req.ok) {
        throw error(404, "Not Found");
    }
    const provincias: Provincia[] = await req.json();
    req = await fetch(`${API_URL}/ciudad`);
    if (!req.ok) {
        throw error(404, "Not Found");
    }
    const ciudades: Ciudad[] = await req.json();
    return {
        provincias,
        ciudades
    };
}

const form: { [key: string]: FormDataEntryValue | null } = {
    nombre: null,
    id_provincia: null,
    error: null
};

export const actions = {
    insert: async ({ request }) => {
        const data = await request.formData();
        const f_nombre = data.get("nombre");
        let f_id_provincia = data.get("id_provincia");
        form.nombre = f_nombre;
        form.id_provincia = f_id_provincia;

        if (typeof f_nombre !== "string" || f_nombre === "") {
            form.error = "Ingrese un nombre";
            return fail(400, form);
        };
        const nombre = f_nombre;

        if (typeof f_id_provincia !== "string" || f_id_provincia === "") {
            form.error = "Seleccione una provincia";
            return fail(400, form);
        };
        const id_provincia = parseInt(f_id_provincia);
        if (isNaN(id_provincia)) {
            form.error = "Seleccione una provincia válida";
            return fail(400, form);
        };

        const req = await fetch(`${API_URL}/ciudad`, {
            method: "POST",
            body: JSON.stringify({ nombre, id_provincia }),
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

        const req = await fetch(`${API_URL}/ciudad/${id}`, {
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
