<script lang="ts">
    import type { PageData } from "./$types";
    export let data: PageData;
    export let form;
</script>

<h1>Provincia</h1>
<a href="/provincia">Ver provincias</a>
<p>{form?.error || ""}</p>
<form method="post" action="?/update">
    <select name="id_region" required>
        <option value="">---</option>
        {#each data.regiones as region (region.id)}
            {#if form?.id_region === String(region.id)}
                <option value={region.id} selected>{region.nombre}</option>
            {:else if form === null && data.provincia.id_region === region.id}
                <option value={region.id} selected>{region.nombre}</option>
            {:else}
                <option value={region.id}>{region.nombre}</option>
            {/if}
        {/each}
    </select>
    <input
        type="text"
        placeholder="Nombre"
        name="nombre"
        value={form?.nombre ?? data.provincia.nombre}
        required
    />
    <input type="submit" value="Actualizar" />
</form>
