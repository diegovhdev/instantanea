<script>
  import Spinner from "./Spinner.svelte";
  let {children, spin, width="auto", height="auto"} = $props();
  console.log("Hola")
</script>

<button style:width style:height style:padding={spin ? "0.3rem" : "0.5rem"}>
{#if spin}
  <Spinner --color-theme="#fff" --border-width="5px"/>
{:else}
  {@render children()}
{/if}

</button>

<style> 
  button {
    display: flex;
    align-items: center;
    background-color: #364153;
    border-radius: 12px;
    color: #fff;
    font-size: 1.1rem;
    font-weight: 600;
    justify-content: center;
    place-self: center;
  }

  button:hover {
    background-color: #101828;
  }
</style>