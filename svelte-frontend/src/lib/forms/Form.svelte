<script>
  import InputField from "./InputField.svelte";
  import Spinner from "$lib/components/Spinner.svelte";

let {textButton, textTitle, fieldsVals, callback, onError, onSuccess} = $props()

let fields = $state(fieldsVals)
let buttonWidth = $state(0)
let buttonHeight = $state(0)
let initSpinner = $state(false)

const clearForm = () => {
  for (const field of fields) {
    field.value = ""
  }
  initSpinner = false
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
} 

</script>

<form onsubmit={handleSubmit}>
  <h1>{textTitle}</h1>
  {#each fields as field (field.name)}
  <InputField label={field.label} bind:value={field.value} type={field.type}/>
  {/each}
  <button bind:clientWidth={buttonWidth} bind:clientHeight={buttonHeight} 
    style:width={buttonWidth !== 0 ? `${buttonWidth}px` : "auto"}
    style:height={buttonHeight !== 0 ? `${buttonHeight}px` : "auto"}>
    {#if initSpinner}
      <Spinner --color-theme="white"/>
    {:else}
    <span>
      {textButton}
    </span>
    {/if}
  </button>
</form>

<style>

  form {
    background-color: #e5e7eb;
    border-radius: 16px;
    border-width: 2px;
    display: flex;
    flex-direction: column;
    gap: 12px;
    justify-items: center;
    padding: 40px 40px 24px;
  }

  h1 {
    font-size: 1.5rem;
    font-weight: 600;
  }

  button {
    background-color: #4a5565;
    border-radius: 12px;
    color: #fff;
    font-weight: 700;
    padding: 10px 12px;
  }

  button:hover {
    background-color: #1e2939;
  }

  h1, button {
    align-self: center;
  }
</style>