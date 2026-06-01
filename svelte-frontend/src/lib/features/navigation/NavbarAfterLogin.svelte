<script>
  import UserProfile from "$lib/components/UserProfile.svelte";
  import SettingsIcon from "$lib/features/icons/SettingsIcon.svelte";
  import { logout } from "$lib/services/api";
  import { userState } from "$lib/stores/global-state.svelte";
  import { ProfileOptions } from "$lib/stores/ui-state.svelte";
  import { fly, slide } from "svelte/transition";
  import LogoutButton from "../settings/LogoutButton.svelte";
  import ProfileActions from "../settings/ProfileActions.svelte";
  import SettingsButton from "../settings/SettingsButton.svelte";

  let showProfileActions = $state()

  const ProfileOptionsPopup = (e) => {
    e.stopPropagation()
    ProfileOptions.show = true;
  }

</script>

<nav transition:slide>
  <h1>Instantanea <span>📸</span></h1>
  <div>
    <a href="/upload">Publicar</a>
    <UserProfile --height="3.4rem" src={userState.profilePictureUrl}>{userState.username}</UserProfile>
    <SettingsButton />
    {#if ProfileOptions.show}
      <ProfileActions />
    {/if}
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
    grid-column: 1 / -1;
    grid-row: 1 / 2;
  }

  nav > h1 {
    font-size: clamp(1.4rem, 2.5vw, 2rem);
    display: flex;
    font-weight: 600;
    flex-wrap: wrap;
    gap: 1rem;
    justify-content: center;
  }

  nav > div {
    display: flex;
    justify-content: space-around;
    align-items: center;
    font-size: 1.1rem;
    position: relative;
  }

  nav > div > a {
    font-size: 1.1rem;
    font-weight: 600;
    background-color: var(--color-secondary);
    border-radius: 0.75rem;
    padding: 0.5rem;
    border-width: 2px;
    color: var(--color-primary);
    transition: background-color 0.3s ease;
  }

   nav > div > a:hover {
    background-color: var(--color-secondary-hover);
   }

</style>