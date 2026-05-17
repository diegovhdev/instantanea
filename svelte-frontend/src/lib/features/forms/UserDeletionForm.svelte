<script>
  import ErrorMessage from "$lib/components/ErrorMessage.svelte";
  import InputField from "$lib/components/InputField.svelte";
  import LoadingButton from "$lib/components/LoadingButton.svelte";

  let {callback} = $props()

  let deletionForm = $state({
    password: ""
  })

  let loading = $state(false)
  let errorMessage = $state("")

  const clearForm = () => {
    deletionForm.password = ""
    loading = false
  }

  async function handleSubmit(e) {
    e.preventDefault()
    loading = true
    try {
      await callback(deletionForm)
    } catch(error) {
      errorMessage = error.message
    }
    clearForm()
  }

</script>

<form onsubmit={handleSubmit}>
  <h3>Eliminacion de cuenta</h3>
  <InputField name="password" bind:value={deletionForm.password} type="password">Contraseña: </InputField>
  <LoadingButton {loading}>Eliminar Cuenta</LoadingButton>
  <ErrorMessage error={errorMessage}/>
</form>


<style>

form {
  margin-top: 0.8rem;
  place-self: center;
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 1rem;
  padding: 12px 20px;
  padding-top: 0;
  border-radius: 16px;
}

h3 {
  font-size: 1.15rem;
  font-weight: 600;
  color: var(--color-primary-hover)
}
</style>