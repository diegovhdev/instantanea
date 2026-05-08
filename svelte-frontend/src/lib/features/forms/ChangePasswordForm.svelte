<script>
  import ErrorMessage from "$lib/components/ErrorMessage.svelte";
  import InputField from "$lib/components/InputField.svelte";
  import LoadingButton from "$lib/components/LoadingButton.svelte";

  let {callback} = $props()

  let passwordForm = $state({
    password: "",
    newPassword: ""
  })

  let loading = $state(false)
  let errorMessage = $state("")

  const clearForm = () => {
    passwordForm.password = "";
    passwordForm.newPassword = "";
    loading = false;
  }

  async function handleSubmit(e) {
    e.preventDefault();
    loading = true;
    try {
      await callback();
      errorMessage = "";
    } catch(error) {
      errorMessage = error.message;
    }
    clearForm();
  }


</script>

<form onsubmit={handleSubmit}>
  <InputField name="password" bind:value={passwordForm.password} type="password">Contreseña: </InputField>
  <InputField name="newPassword" bind:value={passwordForm.newPassword} type="password">Nueva contraseña: </InputField>
  <LoadingButton marginTop="0" {loading}>Cambiar contraseña</LoadingButton>
  <ErrorMessage error={errorMessage} />
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

</style>