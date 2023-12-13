<script lang="ts">
    import type { PageData } from "./$types";
    export let data: PageData;
</script>

<h1>Región</h1>
<table border="1">
    <thead>
        <th>Nombre</th>
        <th>Numero</th>
    </thead>
    <tbody>
        {#each data.regiones as region (region.id)}
            <tr>
                <td>{region.nombre}</td>
                <td>{region.numero}</td>
            </tr>
        {/each}
    </tbody>
</table>
