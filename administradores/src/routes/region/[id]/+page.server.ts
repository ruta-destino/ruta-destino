import { API_URL } from "$env/static/private";
import { error, fail, redirect } from "@sveltejs/kit";
import type { Region } from "$lib/server/models.js";

export async function load({ params }) {
    const id = params.id;
    const req = await fetch(`${API_URL}/region/${id}`);
    if (!req.ok) {
        throw error(404, 'Not Found');
    }
    const res: Region = await req.json();
    return { region: res };
}

type FormKey = "error" | "nombre" | "numero";
const form: { [key in FormKey]: FormDataEntryValue | null } = {
    error: null, nombre: null, numero: null
}

export const actions = {
    update: async ({ request }) => {
        const data = await request.formData();
        const id = data.get("id");
        const f_nombre = data.get("nombre");
        let f_numero = data.get("numero");
        form.nombre = f_nombre;
        form.numero = f_numero;

        if (typeof f_nombre !== "string" || f_nombre === "") {
            form.error = "Ingrese un nombre";
            return fail(400, form);
        }
        const nombre = f_nombre;

        if (typeof f_numero !== "string" || f_numero === "") {
            form.error = "Ingrese un número";
            return fail(400, form);
        }
        const numero = parseInt(f_numero);
        if (isNaN(numero)) {
            form.error = "Ingrese un número válido";
            return fail(400, form);
        }

        const req = await fetch(`${API_URL}/region/${id}`, {
            method: "POST",
            body: JSON.stringify({ nombre, numero }),
            headers: { "Content-Type": "application/json" }
        });

        if (!req.ok) {
            const res = await req.json();
            form.error = res["error"];
            return fail(400, form);
        }

        throw redirect(302, "/region")
    }
}
