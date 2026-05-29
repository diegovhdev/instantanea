<script>
  import { unvotePost, votePost } from "$lib/services/api";
  import VoteIcon from "../icons/VoteIcon.svelte";
  let {data} = $props()
  let votes = $state(data.votes)
  let voted = $state(data.voted)
  let disabled = $state(false)

  async function handleVote() {
    votes += 1
    voted = true
    try {
      await votePost(data.postId)
    } catch (error) {
      console.log(error.message)
      votes -= 1
      voted = false
    };
    
  }

  async function handleUnvote() {
    votes -= 1
    voted = false
    try {
      await unvotePost(data.postId)
    } catch (error) {
      console.log(error.message)
      votes += 1
      voted = true
    };
    
  }

  async function handleClick() {
    if (disabled) return;
    disabled = true
    if (voted) {
      await handleUnvote()
    } else {
      await handleVote()
    }
    disabled = false;
  }

</script>

<button class="rounded-xl" class:voted onclick={handleClick}>
  <VoteIcon />
  <p>{votes}</p>
</button>

<style> 

  p {
    font-size: 1.5rem;
    font-weight: 600;
  }

  .voted {
    fill: var(--color-voted);
    color: var(--color-voted);
  }

  button {
    display: flex;
    justify-content: center;
    align-items: center;
    height: var(--height);
    --size: var(--height);
    cursor: pointer;
    fill: var(--color-primary);
    color: var(--color-primary);
    gap: 0.4rem;
    transition: all 0.3s ease;
  }

  button:hover {
    fill: var(--color-voted);
    color: var(--color-voted);
    transform: scale(1.15);
  }

  .voted:hover {
    fill: var(--color-primary-hover);
    color: var(--color-primary-hover);
  }
</style>

