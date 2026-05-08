<script>
  let {size = "1rem"} = $props()
</script>


<svg xmlns="http://www.w3.org/2000/svg" height={size} viewBox="0 -960 960 960"><path d="M256-192.35 192.35-256l224-224-224-224L256-767.65l224 224 224-224L767.65-704l-224 224 224 224L704-192.35l-224-224-224 224Z"/></svg>

<style>
  svg {
    height: calc(var(--size) * 1);
    background-color: var(--color-background);
    aspect-ratio: 1;
    fill: inherit;
  }
</style>