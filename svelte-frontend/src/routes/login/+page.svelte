<script>
  import { login } from "$lib/api";
  import Form from "$lib/Form.svelte";


  let fieldsVals = [
    {
      name: "username",
      label: "Usuario",
      value: ""
    },
    {
      name: "password",
      label: "Contraseña",
      value: "",
      type: "password"
    }
  ]

</script>

<Form textTitle="Iniciar Sección" textButton="Entrar" {fieldsVals} callback={login}/>

<style>

</style>