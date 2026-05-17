<script>
  import { slide } from "svelte/transition";

</script>

<nav>
  <h1>Instantanea 📸</h1>
  <div>
    <a href="/login" class="ingresar">Ingresar</a>
    <a href="/register" class="registrar">Registrarse</a>
  </div>
</nav>

<style>
  nav {
    align-items: center;
    background-color: var(--color-primary);
    color: var(--color-secondary);
    display: flex;
    justify-content: space-between;
    padding: 1rem;
    place-self: stretch;
    grid-column: 1 / -1;
    grid-row: 1 / 2;
  }

  h1 {
    font-size: 1.5rem;
    font-weight: 600;
  }

  div {
    display: flex;
    align-items: center;
    gap: 1.25rem;
  }

  a {
    font-size: 1.1rem;
    font-weight: 600;
  }

  .ingresar {
    border-radius: 1rem;
    padding: 0.5rem;
    color: var(--color-primary);
    background-color: var(--color-secondary);
    transition: background-color 0.3s ease;
  }

  .ingresar:hover {
    background-color: var(--color-secondary-hover);
  }

  .registrar {
    border-radius: 1rem;
    padding: 0.5rem;
    transition: background-color 0.3s ease;
  }

  .registrar:hover {
    background-color: var(--color-bg-secondary);
  }
</style>