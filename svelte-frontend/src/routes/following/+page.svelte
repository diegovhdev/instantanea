<script>
  import UserCard from '$lib/components/UserCard.svelte';
  import { setContext } from 'svelte';
  import { fade } from 'svelte/transition';

  
  let {data} = $props()
  console.log(data)

  let users = $state(data.users)

  function removeFollow() {
    users = users.filter(u => u.following)
  }

  function callbackFollow(following, userId) {
    for (let i = 0; i < users.length; i++) {
      if (users[i].userId === userId) {
        users[i].following = following
        if (following === false) {
          removeFollow(userId)
        }
      }
    }
  }

  setContext("callbackFollow", callbackFollow)
  
</script>

<div in:fade>
  {#if users.length === 0 }
    <h3>No estas siguiendo a ningun usuario</h3>
  {:else}
  {#each users as data (data.userId)}
    <UserCard {data} gap="2.5rem"/>
  {/each}
  {/if}
</div>


<style>
  div {
    flex-direction: column;
    align-items: center;
    justify-content: flex-start;
    gap: 5rem;
  }

  h3 {
    font-size: 1.5rem;
  }
</style>