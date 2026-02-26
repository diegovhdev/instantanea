<script>
  import InputField from "$lib/InputField.svelte";


  let userRequest = $state({
    username: "",
    password: ""
  })

</script>

<form class="grid justify-items-center gap-1 grid-rows-[2fr_1fr_1fr] bg-gray-200 px-8 py-6 border-2 rounded-2xl">
  <h1 class="text-2xl font-semibold">Iniciar sección</h1>
  <InputField value={userRequest.username} name="usuario" />
  <InputField value={userRequest.password} name="contraseña" type="password"/>
</form>

<style>
</style>