<script>
  import { goto } from "$app/navigation";
  import { register } from "$lib/services/api";
  import RegisterForm from "$lib/features/RegisterForm.svelte";

  let errorMessage = $state("e")
  let invisible = $derived(errorMessage.trim() === "e")

  function onError(error) {
    errorMessage = error.message
  }

  function onSuccess(payload) {
    goto("/login")
  }

</script>

<div>
  <h3 class:invisible >{errorMessage}</h3>
  <RegisterForm callback={register} {onSuccess} {onError}/>
</div>

<style>
  .invisible {
    opacity: 0;
  }

  h3 {
    color: var(--color-error);
  }
</style>

