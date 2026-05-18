<script>
  import VoteIcon from "../icons/VoteIcon.svelte";
  let {data} = $props()
  let votes = $state(data.votes)
</script>

<button class="rounded-xl">
  <VoteIcon />
  <p>{votes}</p>
</button>

<style> 

  p {
    font-size: 1.5rem;
    font-weight: 600;
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
    fill: var(--color-primary-hover);
    color: var(--color-primary-hover);
    transform: scale(1.15);
  }
</style>

