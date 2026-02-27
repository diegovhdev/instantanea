<script>
  import { page } from "$app/state";

  page
</script>

<nav class="flex justify-between p-4 bg-gray-900 text-white">
  <h1 class="font-semibold text-2xl">Instantanea 📸</h1>
  <div class="flex gap-6">
    <a href="/login" class="font-semibold">Iniciar Sección</a>
    <a href="/register" class="font-semibold">Registrarse</a>
  </div>
</nav>

<style>
  nav {
    place-self: stretch;
    align-items: center;
  }
</style>