<script>
  import { userState } from "$lib/stores/global-state.svelte";
  import ProfilePicture from "./ProfilePicture.svelte";

  let {src, children} = $props()
  let hover = $state(false)
</script>

<a class="profile" href={`users/${userState.id}`}>
  <ProfilePicture {src} />
  {@render children()}
</a>

<style>

.profile {
  display: flex;
  padding: 0.4rem 0.8rem;
  padding-right: 0;
  align-items: center;
  gap: 0.5rem;
  font-size: 1.1rem;
  font-weight: 600;
  cursor: pointer;
  transform: scale(0.9);
  transition: transform 0.3s ease;
}

.profile:hover {
  text-decoration: underline;
  transform: scale(1);
}

</style>