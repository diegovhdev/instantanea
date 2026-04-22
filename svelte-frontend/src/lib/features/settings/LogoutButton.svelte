<script>
  import { logout } from "$lib/services/api";
  import LogoutIcon from "../icons/LogoutIcon.svelte";

  let {children, size = "1.8rem"} = $props()

</script>

<button onclick={logout}>
  <LogoutIcon size="1.6rem"/>
  {@render children?.()}
</button>

<style>
  button {
    font-weight: 600;
    color: var(--color-secondary);
    border-radius: 12px;
    gap: 0.3rem;
    display: flex;
    align-items: center;
    justify-content: center;
    padding: 0.5rem;
    cursor: pointer;
    transition: background-color 0.3s ease, color 0.3s ease, fill 0.3s ease;
    fill: var(--color-secondary);
  }

  button:hover {
     background-color: var(--color-secondary);
     color: var(--color-bg-secondary);
     fill: var(--color-bg-secondary)
  }


</style>