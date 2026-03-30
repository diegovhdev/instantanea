<script>
  let {height} = $props()
</script>

<div class="spinner" style:height={`${height}px`}>
</div>

<style>

  .spinner {
    aspect-ratio: 1;
    border: var(--border-width, 5px) solid var(--color-secondary, black);
    border-bottom-color: transparent;
    border-radius: 50%;
    animation: rotation 1s linear infinite;
  }

    @keyframes rotation {
    0% {
        transform: rotate(0deg);
    }
    100% {
        transform: rotate(360deg);
    }
    } 

</style>