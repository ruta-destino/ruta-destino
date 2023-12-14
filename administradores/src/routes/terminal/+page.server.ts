import { API_URL } from "$env/static/private";
import { error, fail } from "@sveltejs/kit";
import type { Ciudad, Terminal } from "$lib/server/models.js";

export async function load() {
    let req = await fetch(`${API_URL}/ciudad`);
    if (!req.ok) {
        throw error(404, "Not Found");
    };
    const ciudades: Ciudad[] = await req.json();
    req = await fetch(`${API_URL}/terminal`);
    if (!req.ok) {
        throw error(404, "Not Found");
    };
    const terminales: Terminal[] = await req.json();
    return {
        ciudades,
        terminales
    };
}

type FormKey = "error" | "nombre" | "direccion" | "latitud" | "longitud" | "id_ciudad";
const form: { [key in FormKey]: FormDataEntryValue | null } = {
    nombre: null, direccion: null, latitud: null, longitud: null, id_ciudad: null, error: null
};

export const actions = {
    insert: async ({ request }) => {
        const data = await request.formData();
        const f_nombre = data.get("nombre");
        const f_direccion = data.get("direccion");
        const f_latitud = data.get("latitud");
        const f_longitud = data.get("longitud");
        let f_id_ciudad = data.get("id_ciudad");
        form.nombre = f_nombre;
        form.direccion = f_direccion;
        form.latitud = f_latitud;
        form.longitud = f_longitud;
        form.id_ciudad = f_id_ciudad;

        if (typeof f_nombre !== "string" || f_nombre === "") {
            form.error = "Ingrese un nombre";
            return fail(400, form);
        };
        const nombre = f_nombre;

        if (typeof f_direccion !== "string" || f_direccion === "") {
            form.error = "Ingrese una direccion";
            return fail(400, form);
        };
        const direccion = f_direccion;

        if (typeof f_latitud !== "string" || f_latitud === "") {
            form.error = "Ingrese una latitud";
            return fail(400, form);
        };
        const latitud = parseFloat(f_latitud);
        if (isNaN(latitud)) {
            form.error = "Ingrese una latitud válida";
            return fail(400, form);
        }
        if (latitud > 90 || latitud < -90) {
            form.error = "Ingrese una latitud entre -90 y 90";
            return fail(400, form);
        }

        if (typeof f_longitud !== "string" || f_longitud === "") {
            form.error = "Ingrese una longitud";
            return fail(400, form);
        };
        const longitud = parseFloat(f_longitud);
        if (isNaN(longitud)) {
            form.error = "Ingrese una longitud válida";
            return fail(400, form);
        }
        if (longitud > 180 || longitud < -180) {
            form.error = "Ingrese una longitud entre -180 y 180";
            return fail(400, form);
        }

        if (typeof f_id_ciudad !== "string" || f_id_ciudad === "") {
            form.error = "Seleccione una ciudad";
            return fail(400, form);
        };
        const id_ciudad = parseInt(f_id_ciudad);
        if (isNaN(id_ciudad)) {
            form.error = "Seleccione una ciudad válida";
            return fail(400, form);
        };

        const req = await fetch(`${API_URL}/terminal`, {
            method: "POST",
            body: JSON.stringify({ nombre, direccion, latitud, longitud, id_ciudad }),
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

        const req = await fetch(`${API_URL}/terminal/${id}`, {
            method: "DELETE",
            headers: { "Content-Type": "application/json" }
        });

        if (!req.ok) {
            const res = await req.json();
            form.error = res["error"];
            return fail(400, form);
        }
    }
}
