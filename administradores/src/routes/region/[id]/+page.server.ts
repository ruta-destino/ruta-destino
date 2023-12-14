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

export const actions = {
    update: async ({ request }) => {
        const data = await request.formData();
        const id = data.get("id");
        const nombre = data.get("nombre");
        let numero = data.get("numero")?.valueOf();
        const form = {
            nombre: nombre,
            numero: numero,
            error: ""
        };

        if (nombre === null || nombre === "") {
            form.error = "Ingrese un nombre";
            return fail(400, form);
        }
        if (typeof numero !== "string") {
            form.error = "Ingrese un número";
            return fail(400, form);
        }
        if (numero === "") {
            form.error = "Ingrese un número";
            return fail(400, form);
        }
        numero = parseInt(numero);

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
