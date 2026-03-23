<script>
  import InputField from "./InputField.svelte";
  import Spinner from "$lib/components/Spinner.svelte";
  import ButtonLoader from "$lib/components/ButtonLoader.svelte";

let {textButton, textTitle, fieldsVals, callback, onError, onSuccess} = $props()

let fields = $state(fieldsVals)
let buttonWidth = $state(0)
let buttonHeight = $state(0)
let spin = $state(false)

const clearForm = () => {
  for (const field of fields) {
    field.value = ""
  }
  spin = false
}

const handleSubmit = async (e) => {
  e.preventDefault()
  const payload = {}
  for (const field of fields) {
    payload[field.name] = field.value
  }
  spin = true
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
  <ButtonLoader {spin} width="4.2rem" height="2.5rem">{textButton}</ButtonLoader>
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