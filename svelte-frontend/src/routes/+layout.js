import { page } from "$app/state";
import { userState } from "$lib/global-state.svelte";
import { redirect } from "@sveltejs/kit";

export const ssr = false;
export const prerender = false;

export function load({ url }) {
  if (
    !userState.logged &&
    url.pathname !== "/login" &&
    url.pathname !== "/register"
  ) {
    throw redirect(302, "/login");
  }
}
