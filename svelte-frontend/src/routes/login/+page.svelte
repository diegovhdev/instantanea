<script>

  let userRequest = $state({
    username: "",
    password: ""
  })

</script>

<form class="flex-col">
  <div>
    <label for="username">username: </label>
    <input type="text" name="username" id="username" bind:value={userRequest.username}>
  </div>
</form>