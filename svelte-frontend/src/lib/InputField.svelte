<script>
  let {name, value, type="text"} = $props()
</script>

<div>
  <label for={name}>{name}:</label>
  <input {type} id={name} name={name} bind:value class="border-2 rounded-md border-gray-400 hover:border-gray-600 
                                                        active:border-gray-600 focus:outline-none pl-2 bg-white"/>
</div>

<style>
  div {
    display: grid;
    grid-template-columns: 95px 1fr;
  }
</style>