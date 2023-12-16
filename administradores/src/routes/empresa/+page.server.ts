import { API_URL } from "$env/static/private";
import { error, fail } from "@sveltejs/kit";
import type { Empresa } from "$lib/server/models.js";

export async function load() {
    const response = await fetch(`${API_URL}/empresa`);
    if (!response.ok) {
        throw error(404, "Not Found");
    };
    const empresas: Empresa[] = await response.json();
    return { empresas };
}

type FormKey = "error" | "nombre";
const form: { [key in FormKey]: FormDataEntryValue | null } = {
    error: null, nombre: null
}

export const actions = {
    insert: async ({ request }) => {
        const data = await request.formData();
        const f_nombre = data.get("nombre");
        form.nombre = f_nombre;

        if (typeof f_nombre !== "string" || f_nombre === "") {
            form.error = "Ingrese un nombre";
            return fail(400, form);
        }
        const nombre = f_nombre;

        const req = await fetch(`${API_URL}/empresa`, {
            method: "POST",
            body: JSON.stringify({ nombre }),
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

        const req = await fetch(`${API_URL}/empresa/${id}`, {
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
