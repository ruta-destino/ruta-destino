<script lang="ts">
    import type { PageData } from "./$types";
    export let data: PageData;
    export let form;
</script>

<h1>Ciudad</h1>
<a href="/ciudad">Ver ciudades</a>
<p>{form?.error || ""}</p>
<form method="post" action="?/update">
    <select name="id_provincia" required>
        <option value="">---</option>
        {#each data.provincias as provincia (provincia.id)}
            {#if form?.id_provincia === String(provincia.id)}
                <option value={provincia.id} selected>
                    {provincia.nombre}
                </option>
            {:else if form === null && data.ciudad.id_provincia === provincia.id}
                <option value={provincia.id} selected>
                    {provincia.nombre}
                </option>
            {:else}
                <option value={provincia.id}>{provincia.nombre}</option>
            {/if}
        {/each}
    </select>
    <input
        type="text"
        placeholder="Nombre"
        name="nombre"
        value={form?.nombre ?? data.ciudad.nombre}
        required
    />
    <input type="submit" value="Actualizar" />
</form>
