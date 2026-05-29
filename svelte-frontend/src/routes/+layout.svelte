<script>
  import { userState } from '$lib/stores/global-state.svelte';
  import Navbar from '$lib/features/navigation/Navbar.svelte';
  import NavbarAfterLogin from "$lib/features/navigation/NavbarAfterLogin.svelte"
	import './layout.css';
  import Navsidebar from '$lib/features/navigation/Navsidebar.svelte';

	let { children } = $props();
  
</script>

<svelte:head></svelte:head>

<div class="layout">
  {#if userState.logged}
    <NavbarAfterLogin />
    <Navsidebar />
  {:else}
    <Navbar />
  {/if}
  <div class="content">
    {@render children()}
  </div>
</div>

<style>
  .layout {
    display: grid;
    grid-template-rows: auto 1fr;
    grid-template-columns: auto 1fr;
    place-items: center; /* Centers horizontally and vertically with one line */
    height: 100vh; 
  }

  .content {
    grid-column: 2 / -1;
    grid-row: 2 / -1;
  }
</style>
