<script lang="ts">
    import type { PageData } from "./$types";

    export let data: PageData;
    export let form;

    const dias = [
        { numero: 1, nombre: "Lunes", id: "lunes" },
        { numero: 2, nombre: "Martes", id: "martes" },
        { numero: 3, nombre: "Miércoles", id: "miercoles" },
        { numero: 4, nombre: "Jueves", id: "jueves" },
        { numero: 5, nombre: "Viernes", id: "viernes" },
        { numero: 6, nombre: "Sábado", id: "sabado" },
        { numero: 0, nombre: "Domingo", id: "domingo" },
    ];
</script>

<h1>Recorrido</h1>
<a href="/recorrido">Ver recorridos</a>
<p>{form?.error ?? ""}</p>
<form action="?/update" method="post">
    <select name="id_terminal_origen" required>
        <option value="">---</option>
        {#each data.terminales as terminal (terminal.id)}
            {#if form?.id_terminal_origen === String(terminal.id)}
                <option value={terminal.id} selected>{terminal.nombre}</option>
            {:else if form === null && data.recorrido.id_terminal_origen === terminal.id}
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
            {:else if form === null && data.recorrido.id_terminal_destino === terminal.id}
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
        value={form?.hora ?? data.recorrido.hora}
    />
    <input
        type="number"
        name="minuto"
        required
        min="0"
        max="59"
        placeholder="Minuto"
        value={form?.minuto ?? data.recorrido.minuto}
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
        {:else if form === null && data.recorrido[dia.numero] === "1"}
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
    <input type="submit" value="Actualizar" />
</form>
