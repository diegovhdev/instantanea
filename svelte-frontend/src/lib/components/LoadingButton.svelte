<script>
  import Spinner from "./Spinner.svelte";

  let {
    children,
    loading = false,
  } = $props();
  
  let height = $state() 

</script>

<button
  disabled={loading}
  aria-busy={loading}
  class="btn"
  bind:clientHeight={height}
>
  <!-- contenido -->
  <span class="content" class:hidden={loading}>
    {@render children()}
  </span>

  <!-- spinner -->
  <span class="overlay" class:show={loading}>
    <Spinner height={height * 0.7} --color-theme="#fff"/>
  </span>
</button>

<style>
  .btn {
    position: relative;
    display: inline-flex;
    align-items: center;
    align-self: center;
    justify-content: center;

    padding: 0.6rem 1rem;
    border-radius: 12px;
    border: none;

    background-color: var(--btn-primary-bg);
    color: var(--btn-primary-text);
    font-weight: 600;

    cursor: pointer;
    transition: opacity 0.2s ease;
  }
  .btn:hover {
    background-color: var(--color-primary-hover);
  }

  .btn:disabled {
    opacity: 0.7;
    cursor: not-allowed;
  }

  /* contenido */
  .content {
    display: inline-block;
    transition: opacity 0.15s ease;
  }

  .hidden {
    visibility: hidden;
  }

  /* overlay */
  .overlay {
    position: absolute;
    inset: 0;

    display: flex;
    align-items: center;
    justify-content: center;

    opacity: 0;
    pointer-events: none;
    transition: opacity 0.15s ease;
  }

  .overlay.show {
    opacity: 1;
  }
</style>