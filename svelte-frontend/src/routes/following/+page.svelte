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
  {#each users as data (data.userId)}
    <UserCard {data} gap="3rem"/>
  {/each}
</div>


<style>
  div {
    flex-direction: column;
    align-items: center;
    justify-content: flex-start;
    gap: 5rem;
  }
</style>