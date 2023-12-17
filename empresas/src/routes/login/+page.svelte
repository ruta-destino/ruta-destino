<script>
    export let form;
</script>

<h1>Inicio de sesión</h1>

<p>{form?.error ?? ""}</p>
<form action="?/login" method="post">
    <input type="number" placeholder="Id" min="1" required name="id" />
    <input type="submit" value="Iniciar sesión" />
</form>
