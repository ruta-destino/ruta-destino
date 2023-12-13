import { API_URL } from "$env/static/private";
import { fail } from "@sveltejs/kit";

type Region = {
    id: number;
    nombre: string;
    numero: number;
};

export async function load() {
    const response = await fetch(`${API_URL}/region`);
    if (!response.ok) {
        return { regiones: [] };
    }
    const regiones: Region[] = await response.json();
    return { regiones: regiones };
}

export const actions = {
    insert: async ({ request }) => {
        const data = await request.formData();
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
            console.log((await req.json()))
            return fail(400, {})
        }
    }
}
