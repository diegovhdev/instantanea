<script>
  import { goto } from "$app/navigation";
  import { login } from "$lib/services/api";
  import Form from "$lib/forms/Form.svelte";
  import { userState } from "$lib/stores/global-state.svelte";

  let errorMessage = $state("e")
  let invisible = $derived(errorMessage.trim() === "e")

  function onError(error) {
    errorMessage = error.message
  }

  function onSuccess(payload) {
    userState.username = payload.username;
    userState.logged = true;
    goto("/")
  }

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

<div>
  <h3 class:invisible >{errorMessage}</h3>
  <Form textTitle="Iniciar Sesión" textButton="Entrar" {fieldsVals} callback={login} {onError} {onSuccess}/>
</div>

<style>
  .invisible {
    opacity: 0;
  }

  h3 {
    color: #fb2c36
  }
</style>