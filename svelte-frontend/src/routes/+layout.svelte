<script>
  import Navbar from '$lib/Navbar.svelte';
	import './layout.css';

	let { children } = $props();
</script>

<svelte:head></svelte:head>

<div>
  <Navbar />
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
