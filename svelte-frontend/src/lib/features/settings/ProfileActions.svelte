<script>
  import LoadingButton from "$lib/components/LoadingButton.svelte";
  import SettingsIcon from "$lib/features/icons/SettingsIcon.svelte";
  import { fade } from "svelte/transition";
  import LogoutButton from "./LogoutButton.svelte";
  import GearIcon from "../icons/GearIcon.svelte";


</script>

<div transition:fade={{duration: 200}}>
  <a href="/settings">
    <GearIcon size="1.6rem" />
    Ajustes
  </a>
  <LogoutButton size="1rem">Cerrar Sesión</LogoutButton>
</div>

<style>
  div {
    display: flex;
    flex-direction: column;
    background-color: var(--color-bg-secondary);
    align-items: left;
    position: absolute;
    padding: 1rem;
    gap: 0.5rem; 
    border-radius: 1rem;
    right: 0;
    top: 100%;
  }

  a {
    border-radius: 12px;
    font-weight: 500;
    padding: 0.5rem;
    gap: 0.3rem;
    display: flex;
    color: var(--color-secondary);
    fill: var(--color-bg-secondary);
    transition: background-color 0.4s ease, color 0.4s ease, fill 0.4s ease;
  }

  a:hover {
    background-color: var(--color-secondary);
    color: var(--color-bg-secondary);
    fill: var(--color-secondary);
  }
</style>