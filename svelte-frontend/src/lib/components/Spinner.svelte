<script>
</script>

<div class="loader">
</div>

<style>


  .loader {
    height: 100%;
    aspect-ratio: 1;
    border: var(--border-width, 5px) solid var(--color-theme, gray);
    border-bottom-color: transparent;
    border-radius: 50%;
    animation: rotation 1s linear infinite;
    }

    @keyframes rotation {
    0% {
        transform: rotate(0deg);
    }
    100% {
        transform: rotate(360deg);
    }
    } 

</style>