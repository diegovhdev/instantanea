<script>
  import { slide } from "svelte/transition";


</script>

<nav transition:slide={{axis: "x"}}>
  <div>
    <a href="/new">Nuevo</a>
    <a href="/popular">Popular</a>
    <a href="/favorites">Favoritos</a>
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
    flex-direction: column;
    position: sticky;
    top: 200px;
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
</style>