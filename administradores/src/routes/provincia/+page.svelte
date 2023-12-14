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
            <th>Actualizar</th>
            <th>Eliminar</th>
        </thead>
        <tbody>
            {#each data.provincias as provincia (provincia.id)}
                <tr>
                    <td>{provincia.nombre_region}</td>
                    <td>{provincia.nombre}</td>
                    <td><a href={`/provincia/${provincia.id}`}>⟳</a></td>
                    <td>
                        <form
                            action="?/delete"
                            method="post"
                            on:submit={eliminar}
                        >
                            <input type="submit" value="X" />
                            <input
                                type="hidden"
                                value={provincia.id}
                                name="id"
                            />
                        </form>
                    </td>
                </tr>
            {/each}
        </tbody>
    </table>
{:else}
    <p>No hay provincias</p>
{/if}
