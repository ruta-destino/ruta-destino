import { API_URL } from "$env/static/private";
import { fail } from "@sveltejs/kit";
import type { Region } from "$lib/server/models.js";

export async function load() {
    const response = await fetch(`${API_URL}/region`);
    if (!response.ok) {
        return { regiones: [] };
    }
    const regiones: Region[] = await response.json();
    return { regiones: regiones };
}

type FormKey = "error" | "nombre" | "numero";
const form: { [key in FormKey]: FormDataEntryValue | null } = {
    error: null, nombre: null, numero: null
}

export const actions = {
    insert: async ({ request }) => {
        const data = await request.formData();
        const f_nombre = data.get("nombre");
        let f_numero = data.get("numero");
        form.nombre = f_nombre;
        form.numero = f_numero;

        if (typeof f_nombre !== "string" || f_nombre === "") {
            form.error = "Ingrese un nombre";
            return fail(400, form);
        }
        const nombre = f_nombre;

        if (typeof f_numero !== "string" || f_nombre === "") {
            form.error = "Ingrese un número";
            return fail(400, form);
        }
        const numero = parseInt(f_numero);
        if (isNaN(numero)) {
            form.error = "Ingrese un número válido";
            return fail(400, form);
        }

        const req = await fetch(`${API_URL}/region`, {
            method: "POST",
            body: JSON.stringify({ nombre, numero }),
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

        const req = await fetch(`${API_URL}/region/${id}`, {
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
