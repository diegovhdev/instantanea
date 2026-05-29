<script>
  import Post from '$lib/features/posts/Post.svelte';
  import { setContext } from 'svelte';
  import { fade } from 'svelte/transition';

  let {data} = $props()
  let posts = $state(data.posts)

  function callbackFollow(following, userId) {
    for (let i = 0; i < posts.length; i++) {
      if (posts[i].userId === userId) {
        posts[i].following = following
      }
    }
  }

  setContext("callbackFollow", callbackFollow)
  
</script>

<div in:fade>
  {#each posts as data (data.postId)}
  <Post {data} />
  {/each}
</div>


<style>
  div {
    margin-top: 3rem;
    display: flex;
    flex-direction: column;
    align-items: center;
    gap: 5rem;
  }
</style>