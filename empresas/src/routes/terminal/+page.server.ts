import { env } from "$env/dynamic/private";
import type { Terminal } from "$lib/server/models";
import { error, fail, redirect } from "@sveltejs/kit";
const API_URL = env.API_URL;

export const load = async ({ cookies }) => {
    const id = cookies.get("empresa_id");
    if (!id) {
        throw redirect(302, "/login");
    };
    let req = await fetch(`${API_URL}/terminal`);
    if (!req.ok) {
        throw error(404, "Not Found");
    };
    const terminales: Terminal[] = await req.json();
    req = await fetch(`${API_URL}/empresa/${id}/terminal`);
    if (!req.ok) {
        throw error(404, "Not Found");
    };
    const empresa_terminales: Terminal[] = await req.json();
    return { terminales, empresa_terminales };
}

type FormKey = "error" | "id_terminal";
const form: { [key in FormKey]: FormDataEntryValue | null } = {
    error: null, id_terminal: null
};

export const actions = {
    link: async ({ request, cookies }) => {
        const id = cookies.get("empresa_id");
        if (!id) {
            throw redirect(302, "/login");
        };
        const data = await request.formData();
        const f_id_terminal = data.get("id_terminal");
        form.id_terminal = f_id_terminal;

        if (typeof f_id_terminal !== "string" || f_id_terminal === "") {
            form.error = "Ingrese un terminal";
            return fail(400, form);
        };
        const id_terminal = parseInt(f_id_terminal);
        if (isNaN(id_terminal) || id_terminal <= 0) {
            form.error = "Ingrese un terminal válido";
            return fail(400, form);
        };

        const req = await fetch(`${API_URL}/empresa/${id}/terminal`, {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify({ id_terminal })
        });
        if (!req.ok) {
            const res = await req.json();
            form.error = res["error"];
            return fail(400, form);
        };
    },
    unlink: async ({ cookies, request }) => {
        const id = cookies.get("empresa_id");
        if (!id) {
            throw redirect(302, "/login");
        };
        const data = await request.formData();
        const f_id_terminal = data.get("id");

        if (typeof f_id_terminal !== "string" || f_id_terminal === "") {
            form.error = "Seleccione un terminal";
            return fail(400, form);
        }
        const id_terminal = parseInt(f_id_terminal);
        if (isNaN(id_terminal) || id_terminal <= 0) {
            form.error = "Seleccione un terminal válido";
            return fail(400, form);
        }

        const req = await fetch(`${API_URL}/empresa/${id}/terminal`, {
            method: "DELETE",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify({ id_terminal })
        });
        if (!req.ok) {
            const res = await req.json();
            form.error = res["error"];
            return fail(400, form);
        }
    }
}
