<script lang="ts">
    import type { PageData } from "./$types";

    export let data: PageData;
    export let form;

    const dias = [
        { numero: "1", nombre: "Lunes", id: "lunes" },
        { numero: "2", nombre: "Martes", id: "martes" },
        { numero: "3", nombre: "Miércoles", id: "miercoles" },
        { numero: "4", nombre: "Jueves", id: "jueves" },
        { numero: "5", nombre: "Viernes", id: "viernes" },
        { numero: "6", nombre: "Sábado", id: "sabado" },
        { numero: "0", nombre: "Domingo", id: "domingo" },
    ];
</script>

<h1>Recorrido</h1>
<p>{form?.error ?? ""}</p>
<form action="?/insert" method="post">
    <select name="id_terminal_origen" required>
        <option value="">---</option>
        {#each data.terminales as terminal (terminal.id)}
            {#if form?.id_terminal_origen === String(terminal.id)}
                <option value={terminal.id} selected>{terminal.nombre}</option>
            {:else}
                <option value={terminal.id}>{terminal.nombre}</option>
            {/if}
        {/each}
    </select>
    <select name="id_terminal_destino" required>
        <option value="">---</option>
        {#each data.terminales as terminal (terminal.id)}
            {#if form?.id_terminal_destino === String(terminal.id)}
                <option value={terminal.id} selected>{terminal.nombre}</option>
            {:else}
                <option value={terminal.id}>{terminal.nombre}</option>
            {/if}
        {/each}
    </select>
    <input
        type="number"
        name="hora"
        required
        min="0"
        max="23"
        placeholder="Hora"
        value={form?.hora ?? ""}
    />
    <input
        type="number"
        name="minuto"
        required
        min="0"
        max="59"
        placeholder="Minuto"
        value={form?.minuto ?? ""}
    />
    {#each dias as dia}
        <br />
        <label for={dia.id}>{dia.nombre}</label>
        {#if form?.[dia.numero]}
            <input
                type="checkbox"
                checked
                id={dia.id}
                value={dia.numero}
                name="dia"
            />
        {:else}
            <input type="checkbox" id={dia.id} value={dia.numero} name="dia" />
        {/if}
    {/each}
    <input type="submit" value="Insertar" />
</form>

{#if data.recorridos.length > 0}
    <table border="1">
        <thead>
            <th>Terminal Origen</th>
            <th>Terminal Destino</th>
            <th>Hora</th>
            <th>Minuto</th>
            <th>Días</th>
        </thead>
        <tbody>
            {#each data.recorridos as recorrido (recorrido.id)}
                <tr>
                    <td>{recorrido.nombre_terminal_origen}</td>
                    <td>{recorrido.nombre_terminal_destino}</td>
                    <td>{recorrido.hora}</td>
                    <td>{recorrido.minuto}</td>
                    <td>{recorrido.dias}</td>
                </tr>
            {/each}
        </tbody>
    </table>
{:else}
    <p>No hay recorridos</p>
{/if}
