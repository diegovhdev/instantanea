<script>
  let {src} = $props()
</script>

<div class="avatar">
    <img {src} alt="profile">
</div>

<style>

.avatar {
  aspect-ratio: 1;
  border-radius: 50%;
  height: calc(var(--height) * 1.12);
  overflow: hidden;
}

.avatar img {
  width: 100%;
  height: 100%;
  image-rendering: auto;
  object-fit: cover;
  object-position: center;
}

</style>