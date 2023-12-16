<script lang="ts">
    import type { PageData } from "./$types";
    export let data: PageData;
    export let form;

    function eliminar(event: SubmitEvent) {
        const continuar = confirm(`¿Desea continuar?`);
        if (!continuar) {
            event.preventDefault();
        }
    }
</script>

<h1>Terminal</h1>

<p>{form?.error || ""}</p>
<form method="post" action="?/insert">
    <select name="id_ciudad">
        <option value="">---</option>
        {#each data.ciudades as ciudad (ciudad.id)}
            {#if form?.id_ciudad === String(ciudad.id)}
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
        value={form?.nombre || ""}
    />
    <input
        type="text"
        placeholder="Dirección"
        name="direccion"
        value={form?.direccion || ""}
    />
    <input
        type="number"
        name="latitud"
        min="-90"
        max="90"
        step="any"
        placeholder="Latitud"
        value={form?.latitud || ""}
    />
    <input
        type="number"
        name="longitud"
        min="-180"
        max="180"
        step="any"
        placeholder="Longitud"
        value={form?.longitud || ""}
    />
    <input type="submit" value="Insertar" />
</form>

{#if data.terminales.length > 0}
    <table border="1">
        <thead>
            <th>Ciudad</th>
            <th>Terminal</th>
            <th>Direccion</th>
            <th>Latitud</th>
            <th>Longitud</th>
            <th>Actualizar</th>
            <th>Eliminar</th>
        </thead>
        <tbody>
            {#each data.terminales as terminal (terminal.id)}
                <tr>
                    <td>{terminal.nombre_ciudad}</td>
                    <td>{terminal.nombre}</td>
                    <td>{terminal.direccion}</td>
                    <td>{terminal.latitud}</td>
                    <td>{terminal.longitud}</td>
                    <td><a href={`/terminal/${terminal.id}`}>⟳</a></td>
                    <td>
                        <form
                            action="?/delete"
                            method="post"
                            on:submit={eliminar}
                        >
                            <input type="submit" value="X" />
                            <input
                                type="hidden"
                                value={terminal.id}
                                name="id"
                            />
                        </form>
                    </td>
                </tr>
            {/each}
        </tbody>
    </table>
{:else}
    <p>No hay terminales</p>
{/if}
