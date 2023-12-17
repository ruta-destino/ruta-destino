<script lang="ts">
    import type { PageData } from "./$types";

    export let data: PageData;
</script>

<h1>{data?.nombre}</h1>
<p>
    Por medio de esta interfaz podrá controlar todos sus datos, para así poder
    compartirlos con el resto de los usuarios que deseen consultar la
    información aquí presente.
</p>
