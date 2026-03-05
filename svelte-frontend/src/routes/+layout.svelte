<script>
  import { userState } from '$lib/global-state.svelte';
  import Navbar from '$lib/Navbar.svelte';
  import NavbarAfterLogin from "$lib/NavbarAfterLogin.svelte"
	import './layout.css';

	let { children } = $props();
</script>

<svelte:head></svelte:head>

<div>
  {#if userState.logged}
    <NavbarAfterLogin />
  {:else}
    <Navbar />
  {/if}
  {@render children()}
</div>

<style>
  div {
    display: grid;
    grid-template-rows: auto 1fr;
    place-items: center; /* Centers horizontally and vertically with one line */
    height: 100vh; 
  }
</style>
