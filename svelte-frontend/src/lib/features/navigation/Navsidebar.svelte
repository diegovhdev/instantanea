<script>
  import { slide } from "svelte/transition";

  import {page} from '$app/state';
  import ArrowIcon from "../icons/ArrowIcon.svelte";
  import { isMobile } from "$lib/services/util";

  let showPaths = $state(true)

  const paths = [
    {name: "new", text: "Nuevo"},
    {name: "popular", text: "Popular"},
    {name: "favorites", text: "Favoritos"},
    {name: "following", text: "Siguiendo"}
  ]

</script>

{#key showPaths}
<nav transition:slide={{axis: "x"}}>
  <div>
    {#if isMobile()}
      <button class={showPaths ? "showing" : ""} onclick={() => showPaths = !showPaths}>
        <ArrowIcon size="2.5rem"/>
      </button>
    {/if}
    {#if showPaths}
      {#each paths as path (path.name)}
        <a href={`/${path.name}`} class={`/${path.name}` === page.url.pathname ? "selected" : ""}>{path.text}</a>
      {/each}
    {/if}
  </div>
</nav>
{/key}

<style>
  nav {
    background-color: color-mix(in srgb, var(--color-primary), white 5%);
    display: flex;
    flex-direction: column;
    grid-column: 1 / 2;
    grid-row: 2 / 3;
    align-self: stretch;
    align-items: center;
  }
  
  nav > div {
    display: flex;
    position: sticky;
    margin-top: 12px;
    top: 0px;
    flex-direction: column;
    gap: 1rem;
    align-items: stretch;
    fill: var(--color-secondary)
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

  nav > div :first-child {
    margin-top: 1rem;
  }

  nav > div button:first-child {
    align-self: center;
  }
  
  .selected {
    background-color: var(--color-secondary);
    color: var(--color-primary);
  }

  button {
    transition: all 0.4s ease;
    padding-left: 0.5rem;
    padding-right: 1rem;
    cursor: pointer;
    transform: rotate(0deg);
  }
  
  button:hover {
    transform: scale(1.2);
  }
  
  .showing {
    transform: rotate(180deg)
  }
</style>