<script>
  let {src, children, callback} = $props()
  let hover = $state(false)
</script>

<button class="profile" onclick={callback}>
  <div class="avatar">
    <img {src} alt="profile">
  </div>
  {@render children()}
</button>

<style>

.profile {
  display: flex;
  padding: 0.4rem 0.8rem;
  padding-right: 0;
  align-items: center;
  gap: 0.5rem;
  font-size: 1.1rem;
  font-weight: 600;
  cursor: pointer;
  transform: scale(0.9);
  transition: transform 0.3s ease;
}

.profile:hover {
  text-decoration: underline;
  transform: scale(1);
}

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