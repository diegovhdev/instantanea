<script>
  import LoadingButton from "$lib/components/LoadingButton.svelte";
  import ProfilePicture from "$lib/components/ProfilePicture.svelte";
  import { userState } from "$lib/stores/global-state.svelte";
  import { fade } from "svelte/transition";
  import DeleteIcon from "../icons/DeleteIcon.svelte";
  import ErrorMessage from "$lib/components/ErrorMessage.svelte";

  let {callback} = $props()

  const steps = [
    {
      text: "Cambiar imagen ↑"
    },
    {
      text: "Confirmar cambio"
    }
  ]

  let files = $state()
  let imageURL = $derived(files && files[0] ? URL.createObjectURL(files[0]) : userState.profilePictureUrl)
  let index = $derived(files ? 1 : 0)
  let loading = $state(true)
  let errorMessage = $state("")

  const clear = () => {
    files = undefined;
    loading = false;
  }

  async function handleSubmit(e) {
    e.preventDefault()
    loading = true
    const formData = new FormData();
    try {
      await callback(formData);
      errorMessage = "";
    } catch(error) {
      errorMessage = error.message;
    }
    clear()
  }

  

</script>
<scrip>

</scrip>

<form onsubmit={handleSubmit}>
  <button onclick={() => files = undefined} class={imageURL === userState.profilePictureUrl ? "invisible" : ""} type="button">
    <DeleteIcon size="1.8rem"/>
  </button>
  <ProfilePicture  --height="120px" src={imageURL}/>
  {#if index===0}
    <label for="profile-upload" id="profile-label" in:fade={{duration: 250}}>
      {steps[index].text}
    </label>
  {:else}
    <div in:fade={{duration: 250}}>
      <LoadingButton>{steps[index].text}</LoadingButton>
    </div>
    {/if}
  <input type="file" accept="image/png, image/jpeg" name="profile-upload" id="profile-upload" class="inputfile" bind:files required oninput={errorMessage=""}>
  <ErrorMessage error={errorMessage}/>
</form>

<style>
  form {
    display: flex;
    flex-direction: column;
    align-items: center;
    gap: 0.2rem;
  }

  .inputfile {
    height: 0.1px;
    opacity: 0;
    overflow: hidden;
    position: absolute;
    width: 0.1px;
    z-index: -1;
  }

  label {
    font-weight: 600;
    background-color: var(--color-primary);
    color: var(--color-secondary);
    padding: 0.6rem 1rem;
    border-radius: 12px;
    cursor: pointer;
    transition: background-color 0.3s ease;
  }

  label:hover {
    background-color: var(--color-primary-hover);
  }

  button {
    font-weight: 700;
    align-self: flex-end;
    fill: var(--color-primary);
    --background-color: red;
    transition: fill 0.3s ease, background-color 0.3s ease;
    border-radius: 16px;
    cursor: pointer;
  }

  button:hover {
    fill: var(--color-secondary);
    background-color: var(--color-primary);
  }

  .invisible {
    visibility: hidden;
  }
</style>