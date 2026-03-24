<script>
  import { post } from "$lib/services/api";
  import Spinner from "$lib/components/Spinner.svelte"
  import ButtonLoader from "$lib/components/ButtonLoader.svelte";

  let text = $state("")
  let files = $state()
  let imageURL = $derived(files && files[0] ? URL.createObjectURL(files[0]) : "")

  let spinnerWidth = $state(0)
  let spinnerHeight = $state()
  let spin = $state(false)


  function clear() {
    spin = false
    files = undefined
    text = ""
  }

  async function handleSubmit(e) {
    e.preventDefault()
    spin = true
    const formData = new FormData();
    formData.append("image", files[0])
    formData.append("text", text)
    await post(formData)
    clear()
  }


</script>

<form onsubmit={handleSubmit}>
  <textarea maxlength="255" name="text" id="post" bind:value={text} placeholder="Escribe aca..."></textarea>
  <label for="image-upload" id="image-label">Subir imagen ↑</label>
  <input type="file" accept="image/png, image/jpeg" name="image-upload" id="image-upload" class="inputfile" bind:files required>
  {#if imageURL !== ""}
    <img src={imageURL} alt="">
  {/if}
  <ButtonLoader {spin} width="6rem" height="2.5rem">Publicar</ButtonLoader>
</form>

<style>

  form {
    border-radius: 12px;
    border-width: 2px;
    display: flex;
    flex-direction: column;
    gap: 10px;
    padding: 16px 24px;
    place-items: center;
  }

  label {
    background-color: #e5e7eb;
    border-radius: 8px;
    border-width: 2px;
    font-weight: 600;
    padding: 8px;
  }

  label:hover {
    background-color: #d1d5dc;
  }

  .inputfile {
    height: 0.1px;
    opacity: 0;
    overflow: hidden;
    position: absolute;
    width: 0.1px;
    z-index: -1;
  }

  img {
    width: 300px;
  }

  textarea {
      all: unset;
      aspect-ratio: 3 / 1;
      width: 400px;
  }

</style>