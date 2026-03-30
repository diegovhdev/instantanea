<script>
  let {name, value = $bindable(), type="text", children} = $props()
</script>

<div>
  <label for={`${name}_id`}>
    {@render children?.()}
  </label>
  <input {type} id={`${name}_id`} name={name} required bind:value >
</div>

<style>
  div {
    display: flex;
    flex-direction: column;
    grid-template-columns: 100px 1fr;
    color: var(--color-primary-hover)
  }

  div > label {
    display: flex;
    font-weight: 500;
    justify-content: flex-start;
    padding-right: 4px;
  }

  div > input {
    background-color: var(--color-secondary);
    border: 2px var(--color-border) solid;
    border-radius: 6px;
    padding-left: 6px;
  }

  div > input:hover {
    border-color: var(--color-border-hover);
  }

  div > input:active {
    border-color: var(--color-border-hover);
  }
  
  div > input:focus {
    outline: none;
    border-color: var(--color-border-focus);
  }

</style>