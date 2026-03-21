<script>
  import { logout } from "./api";
  import { userState } from "./global-state.svelte";

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
    background-color: #101828;
    color: #fff;
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
    background-color: #e5e7eb;
    border-radius: 0.75rem;
    padding: 0.5rem;
    border-width: 2px;
    color: #000;
  }

   nav > div > a:hover {
    background-color: #d1d5dc;
   }

   nav > div > div {
    background-color: #364153;
    border-radius: 12px;
    display: flex;
    height: 2.5rem;
    justify-content: space-between;
    padding: 0.5rem;
    width: 10.75rem;
   }

   nav > div > div:hover {
    background-color: #101828;
   }

   button {
    font-weight: 600;
   }

</style>