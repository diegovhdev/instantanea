<script>
  import { goto } from "$app/navigation";
  import { login } from "$lib/services/api";
  import { userState } from "$lib/stores/global-state.svelte";
  import LoginForm from "$lib/features/forms/LoginForm.svelte";
  import ErrorMessage from "$lib/components/ErrorMessage.svelte";

  let errorMessage = $state("")

  function onError(error) {
    errorMessage = error.message
  }

  function onSuccess(data) {
    console.log(data)
    userState.username = data.username;
    userState.profilePictureUrl = data.profilePictureUrl
    userState.id = data.userId;
    userState.userRole = data.userRole;
    userState.logged = true;
    goto("/new")
  }

</script>

<div>
  <ErrorMessage error={errorMessage} />
  <LoginForm callback={login} {onError} {onSuccess}/>
</div>

<style>

  div {
    display: flex;
    flex-direction: column;
    gap: 1rem;
  }

</style>