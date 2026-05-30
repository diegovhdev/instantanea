<script>
  import { goto } from "$app/navigation";
  import { login } from "$lib/services/api";
  import { saveState, userState } from "$lib/stores/global-state.svelte";
  import LoginForm from "$lib/features/forms/LoginForm.svelte";
  import ErrorMessage from "$lib/components/ErrorMessage.svelte";

  let errorMessage = $state("")
  let invisible = $state(false)

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
    invisible = true
    saveState();
    goto("/new")
  }

</script>

<div class={invisible ? "invisible" : ""}>
  <ErrorMessage error={errorMessage} />
  <LoginForm callback={login} {onError} {onSuccess}/>
</div>

<style>

  div {
    display: flex;
    flex-direction: column;
    gap: 1rem;
  }

  .invisible {
    visibility: hidden;
  }

</style>