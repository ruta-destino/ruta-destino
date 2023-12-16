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

<h1>Ciudad</h1>

<p>{form?.error || ""}</p>
<form method="post" action="?/insert">
    <select name="id_provincia" required>
        <option value="">---</option>
        {#each data.provincias as provincia (provincia.id)}
            {#if form?.id_provincia === String(provincia.id)}
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
        value={form?.nombre || ""}
        required
    />
    <input type="submit" value="Insertar" />
</form>

{#if data.ciudades.length > 0}
    <table border="1">
        <thead>
            <th>Provincia</th>
            <th>Ciudad</th>
            <th>Actualizar</th>
            <th>Eliminar</th>
        </thead>
        <tbody>
            {#each data.ciudades as ciudad (ciudad.id)}
                <tr>
                    <td>{ciudad.nombre_provincia}</td>
                    <td>{ciudad.nombre}</td>
                    <td><a href={`/ciudad/${ciudad.id}`}>⟳</a></td>
                    <td>
                        <form
                            action="?/delete"
                            method="post"
                            on:submit={eliminar}
                        >
                            <input type="submit" value="X" />
                            <input type="hidden" value={ciudad.id} name="id" />
                        </form>
                    </td>
                </tr>
            {/each}
        </tbody>
    </table>
{:else}
    <p>No hay ciudades</p>
{/if}
