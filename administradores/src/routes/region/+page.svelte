<script lang="ts">
    import type { PageData } from "./$types";
    export let data: PageData;
    export let form;
</script>

<h1>Región</h1>

<p>{form?.error || ""}</p>
<form method="post" action="?/insert">
    <input
        type="text"
        placeholder="Nombre"
        name="nombre"
        value={form?.nombre || ""}
    />
    <input
        type="number"
        placeholder="Numero"
        name="numero"
        min="0"
        value={form?.numero || ""}
    />
    <input type="submit" value="Insertar" />
</form>

<table border="1">
    <thead>
        <th>Nombre</th>
        <th>Numero</th>
    </thead>
    <tbody>
        {#each data.regiones as region (region.id)}
            <tr>
                <td>{region.nombre}</td>
                <td>{region.numero}</td>
            </tr>
        {/each}
    </tbody>
</table>
