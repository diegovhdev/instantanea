<script>
  import ErrorMessage from "$lib/components/ErrorMessage.svelte";
  import InputField from "$lib/components/InputField.svelte";
  import LoadingButton from "$lib/components/LoadingButton.svelte";
  import { userState } from "$lib/stores/global-state.svelte";

  let {callback} = $props()

  let username = $state(userState.username)
  let loading = $state(false)
  let errorMessage = $state("")

  async function handleSubmit(e) {
    e.preventDefault();
    loading = true;
    data = {
      username: username
    }
    try {
      await callback(username)
      errorMessage = ""
    } catch(error) {
      errorMessage = error.message;
    }
    loading = false;
  }



</script>

<form onsubmit={handleSubmit}>
  <InputField name="username" bind:value={username}>Nombre de usuario: </InputField>
  <LoadingButton marginTop="0" {loading}>Cambiar nombre</LoadingButton>
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