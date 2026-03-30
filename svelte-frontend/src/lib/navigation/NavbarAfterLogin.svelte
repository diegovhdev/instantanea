<script>
  import { logout } from "$lib/services/api";
  import { userState } from "$lib/stores/global-state.svelte";

</script>

<nav>
  <h1>Instantanea 📸</h1>
  <div>
    <h1>{userState.username}</h1>
    <a href="/upload">Publicar</a>
    <div>
      <button class="font-semibold" onclick={logout}>Cerrar Sesión</button>
      <img src="/logout-icon.svg" alt="logout icon">
    </div>
  </div>
</nav>

<style>

  nav {
    align-items: center;
    background-color: var(--color-primary);
    color: var(--color-secondary);
    display: flex;
    justify-content: space-between;
    padding: 16px;
    place-self: stretch;
  }

  nav > h1 {
    font-weight: 600;
    font-size: 1.5rem;
  }

  nav > div {
    display: flex;
    justify-content: space-between;
    gap: 1.5rem;
  }

  nav > div > h1 {
    font-size: 1.25rem;
    margin-top: 0.5rem;
    font-weight: 600;
  }

  nav > div > a {
    font-weight: 600;
    background-color: var(--color-secondary);
    border-radius: 0.75rem;
    padding: 0.5rem;
    border-width: 2px;
    color: var(--color-primary);
  }

   nav > div > a:hover {
    background-color: var(--color-secondary);
   }

   nav > div > div {
    background-color: var(--color-primary-hover);
    border-radius: 12px;
    display: flex;
    height: 2.5rem;
    justify-content: space-between;
    padding: 0.5rem;
    width: 10.75rem;
   }

   nav > div > div:hover {
    background-color: var(--color-primary);
   }

   button {
    font-weight: 600;
   }

</style>