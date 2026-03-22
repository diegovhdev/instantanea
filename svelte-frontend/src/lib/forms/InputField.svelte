<script>
  let {label, value = $bindable(), type="text"} = $props()
</script>

<div>
  <label for={label}> {label}:</label>
  <input {type} id={label} name={label} required bind:value >
</div>

<style>
  div {
    display: grid;
    grid-template-columns: 100px 1fr;
  }

  div > label {
    display: flex;
    justify-content: flex-end;
    padding-right: 4px;
  }

  div > input {
    background-color: white;
    border: 2px #99a1af solid;
    border-radius: 6px;
    padding-left: 6px;
  }

  div > input:hover {
    border-color: #4a5565;
  }

  div > input:active {
    border-color: #4a5565;
  }

</style>