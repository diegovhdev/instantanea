<script>
  import DeleteIcon from "$lib/features/icons/DeleteIcon.svelte";

  let {error = "", fontSize = "1.1rem"} = $props()
</script>


<p class={error === "" ? "invisible" : ""} style:font-size={fontSize}>
  <DeleteIcon --size={fontSize}/>
  {error === "" ? "placeholder" : error}
</p>

<style>

  p {
    color: var(--color-error);
    display: flex;
    align-items: center;
    fill: var(--color-error);
    gap: 0.3rem;
  }

  .invisible {
    visibility: hidden;
  }
</style>