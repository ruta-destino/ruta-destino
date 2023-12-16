<script lang="ts">
    import type { PageData } from "./$types";
    export let data: PageData;
    export let form;
</script>

<h1>Terminal</h1>
<a href="/terminal">Ver terminales</a>
<p>{form?.error || ""}</p>
<form method="post" action="?/update">
    <select name="id_ciudad">
        <option value="">---</option>
        {#each data.ciudades as ciudad (ciudad.id)}
            {#if form?.id_ciudad === String(ciudad.id)}
                <option value={ciudad.id} selected>{ciudad.nombre}</option>
            {:else if form === null && data.terminal.id_ciudad === ciudad.id}
                <option value={ciudad.id} selected>{ciudad.nombre}</option>
            {:else}
                <option value={ciudad.id}>{ciudad.nombre}</option>
            {/if}
        {/each}
    </select>
    <input
        type="text"
        placeholder="Nombre"
        name="nombre"
        value={form?.nombre ?? data.terminal.nombre}
    />
    <input
        type="text"
        placeholder="Dirección"
        name="direccion"
        value={form?.direccion ?? data.terminal.direccion}
    />
    <input
        type="number"
        name="latitud"
        min="-90"
        max="90"
        step="any"
        placeholder="Latitud"
        value={form?.latitud ?? data.terminal.latitud}
    />
    <input
        type="number"
        name="longitud"
        min="-180"
        max="180"
        step="any"
        placeholder="Longitud"
        value={form?.longitud ?? data.terminal.longitud}
    />
    <input type="hidden" value={data.terminal.id} name="id" />
    <input type="submit" value="Actualizar" />
</form>
