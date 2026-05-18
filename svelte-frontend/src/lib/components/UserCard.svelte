<script>
  import { userState } from "$lib/stores/global-state.svelte";
  import ProfilePicture from "./ProfilePicture.svelte";


let {data} = $props()

</script>


<div class="container">
  <div class="card">
      <ProfilePicture src={data.profilePictureUrl} --height="4rem"/>
      <h3>{data.username}</h3>
  </div>
  {#if userState.id !== data.userId}
    <button>Seguir</button>
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


</style>