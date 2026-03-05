<script>
  import InputField from "./InputField.svelte";
  import Spinner from "./Spinner.svelte";

let {textButton, textTitle, fieldsVals, callback, onError, onSuccess} = $props()

let fields = $state(fieldsVals)
let initSpinner = $state(false)

const gridRowTemplate = Array(fields.length).fill("1fr").join("_")

const clearForm = () => {
  for (const field of fields) {
    field.value = ""
  }
}

const handleSubmit = async (e) => {
  e.preventDefault()
  const payload = {}
  for (const field of fields) {
    payload[field.name] = field.value
  }
  initSpinner = true
  try {
    await callback(payload)
    onSuccess(payload)
  } catch(error) {
    onError(error)
  }
  clearForm()
  initSpinner = false
} 

</script>

<form class="grid justify-items-center gap-3 grid-rows-[2fr_{gridRowTemplate}_auto] bg-gray-200 px-10 py-12 border-2 rounded-2xl mb-36" onsubmit={handleSubmit}>
  <h1 class="text-2xl font-semibold">{textTitle}</h1>
  {#each fields as field (field.name)}
  <InputField label={field.label} bind:value={field.value} type={field.type}/>
  {/each}
  <button class="bg-gray-600 w-auto h-12 rounded-xl border-2 text-white font-bold mt-2 hover:bg-gray-800 px-3">
    {#if initSpinner}
      <Spinner --color-theme="white"/>
    {:else}
      {textButton}
    {/if}
  </button>
</form>

<style>
  span {
    opacity: 0;
  }
</style>