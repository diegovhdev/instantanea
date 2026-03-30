<script>
  import LoadingButton from "$lib/components/LoadingButton.svelte";
  import InputField from "$lib/components/InputField.svelte";
  import { fade } from "svelte/transition";

  let {callback, onError, onSuccess} = $props()

  let fields = $state({
    username: "",
    password: ""
  })

  let loading = $state(false)

  const clearForm = () => {
    fields.username = ""
    fields.password = ""
    loading = false
  }

  const handleSubmit = async (e) => {
    e.preventDefault()
    const payload = {
      username: fields.username,
      password: fields.password
    }
    loading = true
    try {
      await callback(payload)
      onSuccess(payload)
    } catch(error) {
      onError(error)
    }
    clearForm()
  }

</script>

<form class="shadow-xl" onsubmit={handleSubmit} in:fade>
  <h1>Iniciar Sesión</h1>
  <InputField name="username" bind:value={fields.username}>Usuario:</InputField>
  <InputField name="password" bind:value={fields.password} type="password">Contraseña:</InputField>
  <LoadingButton {loading}>Entrar</LoadingButton>
</form>

<style>
  form {
    background-color: var(--color-secondary);
    border-radius: 16px;
    display: flex;
    flex-direction: column;
    gap: 1rem;
    justify-items: center;
    padding: 40px 40px 24px;
  }

  h1 {
    font-size: 1.5rem;
    font-weight: 600;
    color: var(--color-primary);
  }



  h1 {
    align-self: center;
  }
</style>