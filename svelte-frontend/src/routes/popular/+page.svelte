<script>
  import Post from '$lib/features/posts/Post.svelte';
  import { fade } from 'svelte/transition';

  export let data;
</script>

<div in:fade>
  {#each data.posts as data (data.postId)}
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