const STATE_KEY = "app_user_state";
const MAX_AGE_MS = 8 * 60 * 1000; // 8 minutos

export let userState = $state({
  userRole: "user",
  username: "",
  logged: false,
  profilePictureUrl: "",
  id: 0,
});

export function saveState() {
  const payload = {
    data: { ...userState },
    savedAt: Date.now(),
  };
  localStorage.setItem(STATE_KEY, JSON.stringify(payload));
}

export function loadState() {
  try {
    const raw = localStorage.getItem(STATE_KEY);
    if (!raw) return false;

    const { data, savedAt } = JSON.parse(raw);
    const expired = Date.now() - savedAt > MAX_AGE_MS;
    if (expired) {
      localStorage.removeItem(STATE_KEY);
      return false;
    }

    Object.assign(userState, data);
    return true;
  } catch {
    return false;
  }
}

export function clearState() {
  localStorage.removeItem(STATE_KEY);
  Object.assign(userState, {
    userRole: "user",
    username: "",
    logged: false,
    profilePictureUrl: "",
    id: 0,
  });
}
