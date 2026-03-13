import { page } from "$app/state";
import { userState } from "$lib/global-state.svelte";
import { redirect } from "@sveltejs/kit";

export function load() {
  if (!userState.logged) {
    throw redirect(302, "login");
  }
}
