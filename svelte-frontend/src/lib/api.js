import { redirect } from "@sveltejs/kit";
import { userState } from "./global-state.svelte";
import { goto } from "$app/navigation";

const prefix = "http://localhost:8080";

export async function login(payload) {
  const response = await fetch(`${prefix}/login`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(payload),
  });
  if (!response.ok) {
    let message = await response.text();
    throw new Error(message);
  }
}

export async function register(payload) {
  const response = await fetch(`${prefix}/login`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(payload),
  });
  if (!response.ok) {
    let message = await response.text();
    throw new Error(message);
  }
}

export async function logout() {
  fetch(`${prefix}/logout`, {
    method: "GET",
  });
  userState.logged = false;
  userState.username = "";
  console.log("aa");
  goto("/login");
}
