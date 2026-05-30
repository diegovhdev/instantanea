<script>
  import { deletePost } from "$lib/services/api";
  import { getContext } from "svelte";
  import TrashIcon from "../icons/TrashIcon.svelte";
  let {data} = $props()
  let disabled = $state(false)

  const callback = getContext("callbackRemove")

  async function handleClick() {
    if (disabled) return;
    disabled = true
    try {
      await deletePost(data.postId)
    } catch(error) {
      console.log(error.message)
      disabled = false
      return
    }
    callback(data.postId)
  }


</script>

<button class="rounded-2xl" onclick={handleClick}>
  <TrashIcon />
</button>

<style>
  button {
    display: flex;
    justify-content: center;
    align-items: center;
    height: var(--height);
    --size: var(--height);
    cursor: pointer;
    fill: var(--color-primary);
    transition: all 0.3s ease;
  }

  button:hover {
    fill: var(--color-error);
    transform: scale(1.15);
  }
</style>