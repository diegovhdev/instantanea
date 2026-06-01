
<script>
  import UserCard from "$lib/components/UserCard.svelte";
  import { fade } from "svelte/transition";
  import DownloadButton from "./DownloadButton.svelte";
  import VoteButton from "./VoteButton.svelte";
  import RemoveButton from "./RemoveButton.svelte";
  import { userState } from "$lib/stores/global-state.svelte";
  import { elapsedtime } from "$lib/services/elapsed-time";

  let {data} = $props()

</script>

<div class="container shadow-2xl rounded-2xl" in:fade out:fade>
  <UserCard {data} />
  <p>{data.text}</p>
  <p class="date">{elapsedtime(data.createdAt)}</p>
  <div class="post-image">
    <img src={data.url} alt="post">
  </div>
  <div class="footer">
    <div>
      <VoteButton --height="2.5rem" {data}/>
      <DownloadButton --height="2.5rem" {data}/>
    </div>
    {#if userState.id === data.userId || userState.userRole === "mod"}
      <RemoveButton --height="2.5rem" {data}/>
    {/if}
  </div>
</div>


<style>
  .container {
    display: flex;
    flex-direction: column;
    padding: 20px;
    gap: 1rem;
  }

  div > .footer {
    display: flex;
    justify-content: space-between;
  }

  .footer > div {
    display: flex;
    gap: 1rem;

  }

  .post-image {
    max-width: 600px;
    width: 100%;           
    align-self: center;
  }

  .post-image img {
    width: 100%;
    height: auto;          /* mantiene proporciones siempre */
    max-height: 700px;     /* límite visual */
    object-fit: contain;   /* si llega al límite, no recorta */
    display: block;
  }

  p {
    white-space: pre-wrap;
  }

  .date {
    font-size: 0.9rem;
    color: var(--color-primary-hover);
    font-weight: 500;
  }
</style>