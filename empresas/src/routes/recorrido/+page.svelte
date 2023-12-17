<script lang="ts">
    import type { PageData } from "./$types";

    export let data: PageData;
    export let form;
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
    <br />
    <label for="lunes">Lunes</label>
    <input type="checkbox" name="dia" id="lunes" value="1" />
    <br />
    <label for="lunes">Martes</label>
    <input type="checkbox" name="dia" id="martes" value="2" />
    <br />
    <label for="lunes">Miércoles</label>
    <input type="checkbox" name="dia" id="miercoles" value="3" />
    <br />
    <label for="lunes">Jueves</label>
    <input type="checkbox" name="dia" id="jueves" value="4" />
    <br />
    <label for="lunes">Viernes</label>
    <input type="checkbox" name="dia" id="viernes" value="5" />
    <br />
    <label for="lunes">Sábado</label>
    <input type="checkbox" name="dia" id="sabado" value="6" />
    <br />
    <label for="lunes">Domingo</label>
    <input type="checkbox" name="dia" id="domingo" value="0" />
    <input type="submit" value="Insertar" />
</form>

{#if data.recorridos.length > 0}
    <p>Aquí van los recorridos</p>
{:else}
    <p>No hay recorridos</p>
{/if}
