import { API_URL } from "$env/static/private";
import { error, fail, redirect } from "@sveltejs/kit";
import type { Ciudad, Provincia } from "$lib/server/models.js";

export async function load({ params }) {
    const id = params.id;
    let req = await fetch(`${API_URL}/ciudad/${id}`);
    if (!req.ok) {
        throw error(404, "Not Found");
    }
    const ciudad: Ciudad = await req.json();
    req = await fetch(`${API_URL}/provincia`);
    if (!req.ok) {
        throw error(404, "Not Found");
    }
    const provincias: Provincia[] = await req.json()
    const data = {
        provincias,
        ciudad
    }
    return data;
}

type FormKey = "error" | "nombre" | "id_provincia";
const form: { [key in FormKey]: FormDataEntryValue | null } = {
    nombre: null, id_provincia: null, error: null
};

export const actions = {
    update: async ({ request }) => {
        const data = await request.formData();
        const f_id = data.get("id");
        const f_nombre = data.get("nombre");
        const f_id_provincia = data.get("id_provincia");
        form.nombre = f_nombre;
        form.id_provincia = f_id_provincia;

        if (typeof f_nombre !== "string" || f_nombre === "") {
            form.error = "Ingrese un nombre";
            return fail(400, form);
        }
        const nombre = f_nombre;

        if (typeof f_id_provincia !== "string" || f_id_provincia === "") {
            form.error = "Seleccione una región";
            return fail(400, form);
        }
        const id_provincia = parseInt(f_id_provincia);
        if (isNaN(id_provincia)) {
            form.error = "Seleccione una provincia válida";
            return fail(400, form);
        }

        const req = await fetch(`${API_URL}/ciudad/${f_id}`, {
            method: "POST",
            body: JSON.stringify({ nombre, id_provincia }),
            headers: { "Content-Type": "application/json" }
        });

        if (!req.ok) {
            const res = await req.json();
            form.error = res["error"];
            return fail(400, form);
        }

        throw redirect(302, "/ciudad")
    }
}
