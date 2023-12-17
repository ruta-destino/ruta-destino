import { redirect } from "@sveltejs/kit";

export const load = ({ cookies }) => {
    cookies.delete("empresa_id", { path: "/" });
    cookies.delete("empresa_nombre", { path: "/" });
    throw redirect(302, "/login");
}
