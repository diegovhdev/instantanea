<script>
  import InputField from "./InputField.svelte";

let {textButton, textTitle, fieldsVals, callback} = $props()

let fields = $state(fieldsVals)

const gridRowTemplate = Array(fields.length).fill("1fr").join("_")

const clearForm = () => {
  for (const field of fields) {
    field.value = ""
  }
}

const handleSubmit = (e) => {
  e.preventDefault()
  const payload = {}
  for (const field of fields) {
    payload[field.name] = field.value
  }
  callback(payload)
  clearForm()
} 

</script>

<form class="grid justify-items-center gap-3 grid-rows-[2fr_{gridRowTemplate}_auto] bg-gray-200 px-10 py-12 border-2 rounded-2xl mb-36" onsubmit={handleSubmit}>
  <h1 class="text-2xl font-semibold">{textTitle}</h1>
  {#each fields as field (field.name)}
  <InputField label={field.label} bind:value={field.value} type={field.type}/>
  {/each}
  <button class="bg-gray-600 p-3 rounded-xl border-2 text-white font-bold mt-2">{textButton}</button>
</form>