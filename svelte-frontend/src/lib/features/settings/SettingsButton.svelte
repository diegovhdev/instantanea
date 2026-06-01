<script>
  import { ProfileOptions } from "$lib/stores/ui-state.svelte";
  import SettingsIcon from "../icons/SettingsIcon.svelte";


  const handleClick = (e) => {
    e.stopPropagation()
    ProfileOptions.show = !(ProfileOptions.show)
  }
 

</script>

<button 
  onclick={handleClick}
>
  <SettingsIcon size="1.8rem" />
</button>

<style>
  button {
    display: flex;
    padding-left: 0.4rem;
    padding-right: 0.2rem;
    align-items: center;
    justify-content: center;
    cursor: pointer;
    transition: transform 0.4s ease;
  }


  button:hover {
    transform: scale(1.1) rotate(270deg);
  }


</style>