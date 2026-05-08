<script>
  import ErrorMessage from "$lib/components/ErrorMessage.svelte";
  import InputField from "$lib/components/InputField.svelte";
  import LoadingButton from "$lib/components/LoadingButton.svelte";

  let {callback} = $props()

  let emailForm = $state({
    email: "",
    newEmail: ""
  })

  let loading = $state(false)
  let errorMessage = $state("")

  const clearForm = () => {
    emailForm.email = "";
    emailForm.newEmail = "";
    loading = false;
  }

  async function handleSubmit(e) {
    e.preventDefault();
    loading = true;
    try {
      await callback(emailForm);
      errorMessage = ""
    } catch(error) {
      errorMessage = error.message;
    }
    clearForm()
  }

</script>

<form onsubmit={handleSubmit}>
  <InputField name="email" bind:value={emailForm.email}>Correo: </InputField>
  <InputField name="newEmail" bind:value={emailForm.newEmail}>Nuevo correo: </InputField>
  <LoadingButton marginTop="0" {loading}>Cambiar correo</LoadingButton>
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

</style>