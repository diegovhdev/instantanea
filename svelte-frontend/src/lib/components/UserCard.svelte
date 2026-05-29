<script>
  import { followUser, unfollowUser } from "$lib/services/api";
  import { userState } from "$lib/stores/global-state.svelte";
  import { getContext } from "svelte";
  import ProfilePicture from "./ProfilePicture.svelte";
  import { fade } from "svelte/transition";


let {data, gap="0"} = $props()
let disabled = $state(false)

let textFollow = $derived(data.following ? "siguiendo" : "seguir")


const callbackFollow = getContext("callbackFollow")

async function handleFollow() {
  callbackFollow(true, data.userId); // optimistic
  try {
    await followUser(data.userId)
  } catch (error) {
    callbackFollow(false, data.userId); // revert
  }
}

async function handleUnfollow() {
  callbackFollow(false, data.userId); // optimistic
  try {
    await unfollowUser(data.userId)
  } catch (error) {
    callbackFollow(true, data.userId); // revert
  }
}

async function handleClick() {
  if (disabled) return;
  disabled = true
  if (data.following) {
    await handleUnfollow()
  } else {
    await handleFollow()
  }
  disabled = false;
}

</script>


<div class="container" out:fade style:gap>
  <button class="card">
      <ProfilePicture src={data.profilePictureUrl} --height="4rem"/>
      <h3>{data.username}</h3>
  </button>
  {#if userState.id !== data.userId}
    <button class="follow-button rounded-2xl" onclick={handleClick}>{textFollow}</button>
  {:else}
    <div></div>
  {/if}
</div>


<style>
  .container {
    justify-content: space-between;
  }

  h3 {
    font-size: 1.2rem;
    font-weight: 600;
  }

  .card {
    gap: 0.7rem;
  }

  div {
    display: flex;
    align-items: center;
  }

  button {
    display: flex;
    align-items: center;
  }

  .follow-button {
    font-weight: 600;
    font-size: 1.2rem;
    background-color: var(--color-primary);
    color: var(--color-secondary);
    padding: 0.4rem 0.8rem;
    cursor: pointer;
    transition: all 0.3s ease;
  }

  .follow-button:hover {
    background-color: var(--color-primary-hover);
  }


</style>