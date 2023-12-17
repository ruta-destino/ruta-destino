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
    req = await fetch(`${API_URL}/empresa/${id}/recorrido`);
    if (!req.ok) {
        throw error(404, "Not Found");
    };
    const recorridos = await req.json();
    return { terminales, recorridos };
}

const form: { [key: string]: FormDataEntryValue | null } = {};

export const actions = {
    insert: async ({ request, cookies }) => {
        const id = cookies.get("empresa_id");
        if (!id) {
            throw redirect(302, "/login");
        };
        const data = await request.formData();
        const f_id_terminal_origen = data.get("id_terminal_origen");
        const f_id_terminal_destino = data.get("id_terminal_destino");
        const f_hora = data.get("hora");
        const f_minuto = data.get("minuto");
        const f_dias = data.getAll("dia");

        form.id_terminal_origen = f_id_terminal_origen;
        form.id_terminal_destino = f_id_terminal_destino;
        form.hora = f_hora;
        form.minuto = f_minuto;

        // Apagar todos los checkbox
        for (let i = 0; i < 7; i++) {
            const dia = String(i);
            form[dia] = null;
        }

        // Encender los checkbox que se enviaron y crear máscara
        const mascara = [..."0000000"];
        for (let f_dia of f_dias) {
            if (typeof f_dia !== "string") {
                continue;
            };
            const dia = parseInt(f_dia);
            if (isNaN(dia) || dia < 0 || dia > 6) {
                continue;
            }
            mascara[dia] = "1";
            form[String(dia)] = String(dia);
        }
        const dias = mascara.join("");
        if (dias.length !== 7) {
            form.error = "Días inválidos";
            return fail(400, form);
        };

        if (typeof f_id_terminal_origen !== "string" || f_id_terminal_origen === "") {
            form.error = "Ingrese un terminal de origen";
            return fail(400, form);
        };
        const id_terminal_origen = parseInt(f_id_terminal_origen);
        if (isNaN(id_terminal_origen) || id_terminal_origen <= 0) {
            form.error = "Ingrese un terminal de origen válido";
            return fail(400, form);
        };

        if (typeof f_id_terminal_destino !== "string" || f_id_terminal_destino === "") {
            form.error = "Ingrese un terminal de destino";
            return fail(400, form);
        };
        const id_terminal_destino = parseInt(f_id_terminal_destino);
        if (isNaN(id_terminal_destino) || id_terminal_destino <= 0) {
            form.error = "Ingrese un terminal de destino válido";
            return fail(400, form);
        };

        if (id_terminal_origen === id_terminal_destino) {
            form.error = "Terminal de origen y destino deben ser diferentes";
            return fail(400, form);
        };

        if (typeof f_hora !== "string" || f_hora === "") {
            form.error = "Ingrese una hora";
            return fail(400, form);
        };
        const hora = parseInt(f_hora);
        if (isNaN(hora)) {
            form.error = "Ingrese una hora válida";
            return fail(400, form);
        };
        if (hora < 0 || hora > 23) {
            form.error = "Ingrese una hora entre 0 y 23";
            return fail(400, form);
        };

        if (typeof f_minuto !== "string" || f_minuto === "") {
            form.error = "Ingrese un minuto";
            return fail(400, form);
        };
        const minuto = parseInt(f_minuto);
        if (isNaN(minuto)) {
            form.error = "Ingrese un minuto válido";
            return fail(400, form);
        };
        if (minuto < 0 || minuto > 59) {
            form.error = "Ingrese un minuto entre 0 y 59";
            return fail(400, form);
        };

        return fail(400, form);
    }
}
