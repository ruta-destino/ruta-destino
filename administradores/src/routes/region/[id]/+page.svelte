<script lang="ts">
    import type { PageData } from "./$types";
    export let data: PageData;
    export let form;
</script>

<h1>Región</h1>
<a href="/region">Ver regiones</a>
<p>{form?.error || ""}</p>
<form action="?/update" method="post">
    <input
        type="text"
        value={form?.nombre ?? data.region.nombre}
        name="nombre"
        placeholder="Nombre"
        required
    />
    <input
        type="number"
        value={form?.numero ?? data.region.numero}
        name="numero"
        min="1"
        placeholder="Numero"
        required
    />
    <input type="submit" value="Actualizar" />
</form>
