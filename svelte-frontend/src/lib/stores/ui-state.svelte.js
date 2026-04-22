export let ProfileOptions = $state({
  show: false,
});

const closeProfileMenu = () => {
  ProfileOptions.show = false;
};

document.body.addEventListener("click", closeProfileMenu);
