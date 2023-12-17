<script lang="ts">
    import type { PageData } from "./$types";

    export let data: PageData;
    export let form;
</script>

<h1>Terminal</h1>
<p>{form?.error ?? ""}</p>
<form action="?/link" method="post">
    <select name="id_terminal" required>
        <option value="">---</option>
        {#each data.terminales as terminal (terminal.id)}
            {#if form?.id_terminal === String(terminal.id)}
                <option value={terminal.id} selected>{terminal.nombre}</option>
            {:else}
                <option value={terminal.id}>{terminal.nombre}</option>
            {/if}
        {/each}
    </select>
    <input type="submit" value="Vincular" />
</form>

{#if data.empresa_terminales.length > 0}
    <table border="1">
        <thead>
            <th>Terminal</th>
            <th>Desvincular</th>
        </thead>
        <tbody>
            {#each data.empresa_terminales as terminal (terminal.id)}
                <tr>
                    <td>{terminal.nombre}</td>
                    <td>
                        <form action="?/unlink" method="post">
                            <input
                                type="hidden"
                                value={terminal.id}
                                name="id"
                            />
                            <input type="submit" value="X" />
                        </form>
                    </td>
                </tr>
            {/each}
        </tbody>
    </table>
{:else}
    <p>No hay terminales</p>
{/if}
