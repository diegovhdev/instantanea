<script>
  import DownloadIcon from "../icons/DownloadIcon.svelte";


</script>

<button class="rounded-2xl">
  <DownloadIcon />
</button>

<style>
  button {
    display: flex;
    justify-content: center;
    align-items: center;
    height: var(--height);
    --size: var(--height);
    cursor: pointer;
    fill: var(--color-primary);
    transition: all 0.3s ease;
  }

  button:hover {
    fill: var(--color-primary-hover);
    transform: scale(1.15);
  }
</style>