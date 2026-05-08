<script>
  import { goto } from "$app/navigation";
  import { login } from "$lib/services/api";
  import { userState } from "$lib/stores/global-state.svelte";
  import LoginForm from "$lib/features/forms/LoginForm.svelte";

  let errorMessage = $state("e")
  let invisible = $derived(errorMessage.trim() === "e")

  function onError(error) {
    errorMessage = error.message
  }

  function onSuccess(data) {
    console.log(data)
    userState.username = data.username;
    userState.profilePictureUrl = data.profileImageUrl
    userState.id = data.id;
    userState.logged = true;
    goto("/")
  }

</script>

<div>
  <h3 class:invisible >{errorMessage}</h3>
  <LoginForm callback={login} {onError} {onSuccess}/>
</div>

<style>

  .invisible {
    opacity: 0;
  }

  h3 {
    color: var(--color-error)
  }
</style>