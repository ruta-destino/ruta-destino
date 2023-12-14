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

const form: { [key: string]: FormDataEntryValue | null } = {
    nombre: null,
    id_region: null,
    error: null
};

export const actions = {
    update: async ({ request }) => {
        const data = await request.formData();
        const fId = data.get("id");
        const fNombre = data.get("nombre");
        const fIdRegion = data.get("id_region");

        form.nombre = fNombre;
        form.id_region = fIdRegion;

        if (fNombre === null || fNombre === "") {
            form.error = "Ingrese un nombre";
            return fail(400, form);
        }
        if (typeof fIdRegion !== "string") {
            form.error = "Seleccione una región";
            return fail(400, form);
        }
        if (fIdRegion === "") {
            form.error = "Seleccione una región";
            return fail(400, form);
        }
        const id_region = parseInt(fIdRegion);

        const req = await fetch(`${API_URL}/provincia/${fId}`, {
            method: "POST",
            body: JSON.stringify({ nombre: fNombre, id_region: id_region }),
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
