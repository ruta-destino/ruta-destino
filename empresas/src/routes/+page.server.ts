import { redirect } from "@sveltejs/kit";

export const load = async ({ cookies }) => {
    const id = cookies.get("empresa_id");
    const nombre = cookies.get("empresa_nombre");

    if (!id || !nombre) {
        throw redirect(302, "/login");
    };

    return { id, nombre };
}
