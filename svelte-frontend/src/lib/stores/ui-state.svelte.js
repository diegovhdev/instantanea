export let ProfileOptions = $state({
  show: false,
});

export let NavbarInfo = $state({});

const closeProfileMenu = () => {
  ProfileOptions.show = false;
};

document.body.addEventListener("click", closeProfileMenu);
