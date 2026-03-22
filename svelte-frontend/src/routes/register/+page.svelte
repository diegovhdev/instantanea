<script>
  import { goto } from "$app/navigation";
  import { register } from "$lib/services/api";
  import Form from "$lib/forms/Form.svelte";

  let errorMessage = $state("e")
  let invisible = $derived(errorMessage.trim() === "e")

  function onError(error) {
    errorMessage = error.message
  }

  function onSuccess(payload) {
    goto("/login")
  }

  const fieldsVals = [
    {
      name: "username",
      label: "Usuario",
      value: ""
    },
    {
      name: "email",
      label: "Correo",
      value: "",
      type: "email"
    },
    {
      name: "password",
      label: "Contraseña",
      value: "",
      type: "password"
    }
  ]
</script>

<div>
  <h3 class:invisible >{errorMessage}</h3>
  <Form textTitle="Registro" textButton="Registrarse" {fieldsVals} callback={register} {onError} {onSuccess}></Form>
</div>

<style>
  .invisible {
    opacity: 0;
  }

  h3 {
    color: #fb2c36;
  }
</style>

