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

<h1>Empresa</h1>

<p>{form?.error || ""}</p>
<form method="post" action="?/insert">
    <input
        type="text"
        placeholder="Nombre"
        name="nombre"
        value={form?.nombre || ""}
    />
    <input type="submit" value="Insertar" />
</form>

{#if data.empresas.length > 0}
    <table border="1">
        <thead>
            <th>Nombre</th>
            <th>Actualizar</th>
            <th>Eliminar</th>
        </thead>
        <tbody>
            {#each data.empresas as empresa (empresa.id)}
                <tr>
                    <td>{empresa.nombre}</td>
                    <td><a href={`empresa/${empresa.id}`}>⟳</a></td>
                    <td>
                        <form
                            action="?/delete"
                            method="post"
                            on:submit={eliminar}
                        >
                            <input type="submit" value="X" />
                            <input type="hidden" value={empresa.id} name="id" />
                        </form>
                    </td>
                </tr>
            {/each}
        </tbody>
    </table>
{:else}
    <p>No hay empresas</p>
{/if}
