import { getFollowingUsers } from "$lib/services/api";
import { userState } from "$lib/stores/global-state.svelte";

export async function load() {
  const users = await getFollowingUsers(userState.id);

  return {
    users,
  };
}
