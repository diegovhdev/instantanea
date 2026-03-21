<script>
</script>

<div>
</div>

<style>
  div {
    border-radius: 50%;
    border: var(--border-width, 5px) solid var(--color-theme, gray);
    border-bottom-color: transparent;
    height: 100%;
    aspect-ratio: 1;
    animation: spin 1s linear infinite;
  }

  div {
    place-self: center;
  }

  @keyframes spin {
    to {
      transform: rotate(360deg);
    }
  }

  span {
    opacity: 0;
  }
</style>