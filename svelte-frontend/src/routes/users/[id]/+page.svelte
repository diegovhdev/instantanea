<script>
  import Post from '$lib/features/posts/Post.svelte';
  import { setContext } from 'svelte';
  import { fade } from 'svelte/transition';

  let {data} = $props();
  let posts = $state(data.posts)

  function callbackFollow(following, userId) {
    for (let i = 0; i < posts.length; i++) {
      if (posts[i].userId === userId) {
        posts[i].following = following
      }
    }
  }

  function removePost(postId) {
    posts = posts.filter(p => p.postId !== postId)
  }

  setContext("callbackRemove", removePost)
  setContext("callbackFollow", callbackFollow)
</script>

<div in:fade>
  {#if posts.length === 0 }
    <h3>Este usuario no ha publicado ninguna imagen</h3>
  {:else}
  {#each posts as data (data.postId)}
    <Post {data} />
  {/each}
  {/if}
</div>


<style>
  div {
    margin-top: 3rem;
    display: flex;
    flex-direction: column;
    align-items: center;
    gap: 5rem;
  }

  h3 {
    font-size: 1.5rem;
  }
</style>