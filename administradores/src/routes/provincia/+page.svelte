<script lang="ts">
    import type { PageData } from "./$types";
    export let data: PageData;
    export let form;
</script>

<h1>Provincia</h1>

<p>{form?.error || ""}</p>
<form method="post" action="?/insert">
    <select name="region">
        <option value="">---</option>
        {#each data.regiones as region (region.id)}
            <option value={region.id}>{region.nombre}</option>
        {/each}
    </select>
    <input
        type="text"
        placeholder="Nombre"
        name="nombre"
        value={form?.nombre || ""}
    />
    <input type="submit" value="Insertar" />
</form>

{#if data.provincias.length > 0}
    <table border="1">
        <thead>
            <th>Region</th>
            <th>Provincia</th>
        </thead>
        <tbody>
            {#each data.provincias as provincia (provincia.id)}
                <tr>
                    <td>{provincia.id_region}</td>
                    <td>{provincia.nombre}</td>
                </tr>
            {/each}
        </tbody>
    </table>
{:else}
    <p>No hay provincias</p>
{/if}
