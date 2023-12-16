<script lang="ts">
    import type { PageData } from "./$types.js";
    export let data: PageData;
    export let form;
</script>

<h1>Empresa</h1>

<p>{form?.error || ""}</p>
<form method="post" action="?/update">
    <input
        type="text"
        placeholder="Nombre"
        name="nombre"
        value={form?.nombre ?? data.empresa.nombre}
        required
    />
    <input type="submit" value="Actualizar" />
</form>
