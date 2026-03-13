<script>
  import { uploadImage } from "./api";
  import Spinner from "./Spinner.svelte"

  let text = $state("")
  let files = $state()
  let imageURL = $derived(files && files[0] ? URL.createObjectURL(files[0]) : "")
  let initSpinner = $state(false)


  function clear() {
    initSpinner = false
    files = undefined
    text = ""
  }

  async function handleSubmit(e) {
    e.preventDefault()
    initSpinner = true
    const formData = new FormData();
    formData.append("image", files[0])
    formData.append("text", text)
    await uploadImage(formData)
    clear()
  }


</script>

<form class="flex flex-col border-2 px-6 py-4 rounded-xl" onsubmit={handleSubmit}>
  <textarea maxlength="255" name="text" id="post" bind:value={text} class="border-none" placeholder="Escribe aca..."></textarea>
  <label for="image-upload" id="image-label" class="font-semibold bg-gray-200 p-2 rounded-xl border-2 hover:bg-gray-300">Subir imagen ↑</label>
  <input type="file" accept="image/png, image/jpeg" name="image-upload" id="image-upload" class="inputfile" bind:files>
  {#if imageURL !== ""}
    <img src={imageURL} alt="">
  {/if}
  <button class="bg-gray-700 text-white font-semibold p-2 rounded-xl hover:bg-gray-900" disabled={initSpinner}>
    {#if initSpinner}
      <Spinner --color-theme="white"/>
    {:else}
      Publicar
    {/if}
  </button>
</form>

<style>

  form {
    place-items: center;
    gap: 10px;
  }

  .inputfile {
    width: 0.1px;
    height: 0.1px;
    opacity: 0;
    overflow: hidden;
    position: absolute;
    z-index: -1;
  }

  img {
    width: 300px;
  }

  textarea {
      all: unset;
      width: 400px;
      aspect-ratio: 3 / 1;
  }

</style>