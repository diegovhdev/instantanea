<script>
  import { logout } from "./api";
  import { userState } from "./global-state.svelte";

</script>

<nav class="flex justify-between p-4 bg-gray-900 text-white">
  <h1 class="font-semibold text-2xl">Instantanea 📸</h1>
  <div class="flex justify-between gap-6">
    <h1 class="text-xl mt-2 font-semibold">{userState.username}</h1>
    <div class="h-10 w-43 flex justify-between bg-gray-700 rounded-xl p-2 hover:bg-gray-900">
      <button class="font-semibold" onclick={logout}>Cerrar Sesión</button>
      <img src="/logout-icon.svg" alt="logout icon">
    </div>
  </div>
</nav>

<style>

  nav {
    place-self: stretch;
    align-items: center;
  }

  div {
    align-items: middle;
  }

</style>