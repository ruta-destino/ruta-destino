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

{#if data.regiones.length > 0}
    <table border="1">
        <thead>
            <th>Nombre</th>
            <th>Numero</th>
            <th>Eliminar</th>
        </thead>
        <tbody>
            {#each data.regiones as region (region.id)}
                <tr>
                    <td>{region.nombre}</td>
                    <td>{region.numero}</td>
                    <td>
                        <form
                            action="?/delete"
                            method="post"
                            on:submit={eliminar}
                        >
                            <input type="submit" value="X" />
                            <input type="hidden" value={region.id} name="id" />
                        </form>
                    </td>
                </tr>
            {/each}
        </tbody>
    </table>
{:else}
    <p>No hay regiones</p>
{/if}
