import { env } from '$env/dynamic/private';
import type { Empresa } from '$lib/server/models.js';
import { fail, redirect } from '@sveltejs/kit';
const API_URL = env.API_URL;

type FormKey = "error";
const form: { [key in FormKey]: FormDataEntryValue | null } = {
    error: null
}

export const actions = {
    login: async ({ request, cookies }) => {
        const data = await request.formData();
        const f_id = data.get("id");

        if (typeof f_id !== "string" || f_id === "") {
            form.error = "Ingrese un id";
            return fail(400, form);
        };
        const id = parseInt(f_id);
        if (isNaN(id) || id <= 0) {
            form.error = "Ingrese un id válido";
            return fail(400, form);
        };

        const req = await fetch(`${API_URL}/empresa/${id}`)
        if (!req.ok) {
            form.error = "Esa empresa no existe";
            return fail(400, form);
        };

        const empresa: Empresa = await req.json();
        cookies.set("empresa_id", String(empresa.id), { path: "/" });
        cookies.set("empresa_nombre", String(empresa.nombre), { path: "/" });

        throw redirect(302, "/");
    }
};
