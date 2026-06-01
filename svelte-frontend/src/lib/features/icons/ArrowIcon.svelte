<script>
  let {size = "1rem"} = $props()
</script>

<svg xmlns="http://www.w3.org/2000/svg" height={size} viewBox="0 -960 960 960"><path d="M647-440H160v-80h487L423-744l57-56 320 320-320 320-57-56 224-224Z"/></svg>

<style>
  svg {
    height: calc(var(--size) * 1);
    background-color: inherit;
    aspect-ratio: 1;
    fill: inherit;
    margin-top: var(--margin-top);
  }
</style>