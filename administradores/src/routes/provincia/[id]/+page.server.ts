import { API_URL } from "$env/static/private";
import { error, fail, redirect } from "@sveltejs/kit";
import type { Provincia, Region } from "$lib/server/models.js";

export async function load({ params }) {
    const id = params.id;
    let req = await fetch(`${API_URL}/provincia/${id}`);
    if (!req.ok) {
        throw error(404, "Not Found");
    }
    const provincia: Provincia = await req.json();
    req = await fetch(`${API_URL}/region`);
    if (!req.ok) {
        throw error(404, "Not Found");
    }
    const regiones: Region[] = await req.json()
    const data = {
        regiones: regiones,
        provincia: provincia
    }
    return data;
}

type FormKey = "error" | "nombre" | "id_region";
const form: { [key in FormKey]: FormDataEntryValue | null } = {
    nombre: null, id_region: null, error: null
};

export const actions = {
    update: async ({ request }) => {
        const data = await request.formData();
        const f_id = data.get("id");
        const f_nombre = data.get("nombre");
        const f_id_region = data.get("id_region");
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
        if (isNaN(id_region) || id_region <= 0) {
            form.error = "Seleccione una región válida";
            return fail(400, form);
        }

        const req = await fetch(`${API_URL}/provincia/${f_id}`, {
            method: "POST",
            body: JSON.stringify({ nombre, id_region }),
            headers: { "Content-Type": "application/json" }
        });

        if (!req.ok) {
            const res = await req.json();
            form.error = res["error"];
            return fail(400, form);
        }

        throw redirect(302, "/provincia")
    }
}
