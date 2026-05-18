<script>
  import { slide } from "svelte/transition";

  import {page} from '$app/state';

  const lastSegment = $derived(page.url.pathname.split('/').pop());

  const paths = [
    {name: "new", text: "Nuevo"},
    {name: "popular", text: "Popular"},
    {name: "favorites", text: "Favoritos"},
  ]

</script>

<nav transition:slide={{axis: "x"}}>
  <div>
    {#each paths as path (path.name)}
      <a href={`/${path.name}`} class={path.name === lastSegment ? "selected" : ""}>{path.text}</a>
    {/each}
  </div>
</nav>

<style>
  nav {
    background-color: color-mix(in srgb, var(--color-primary), white 5%);
    grid-column: 1 / 2;
    grid-row: 2 / 3;
    align-self: stretch;
  }
  
  nav > div {
    display: flex;
    position: sticky;
    margin-top: 24px;
    top: 0px;
    flex-direction: column;
    gap: 1rem;
  }


  a {
    color: var(--color-secondary);
    font-weight: 600;
    font-size: 1.3rem;
    padding: 0.3rem 1rem;
    transition: background-color, 0.3s, ease, color 0.3s ease;
    cursor: pointer;
  }

  a:hover {
    background-color: var(--color-secondary);
    color: var(--color-primary)
  }

  a:first-child {
    margin-top: 40px;
  }

  .selected {
    background-color: var(--color-secondary);
    color: var(--color-primary);
  }
</style>