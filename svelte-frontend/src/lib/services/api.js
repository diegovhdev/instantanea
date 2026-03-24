import { redirect } from "@sveltejs/kit";
import { userState } from "$lib/stores/global-state.svelte";
import { goto } from "$app/navigation";

const prefix = "http://localhost:8080";

export async function login(payload) {
  const response = await fetch(`${prefix}/auth/login`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    credentials: "include",
    body: JSON.stringify(payload),
  });
  if (!response.ok) {
    let message = await response.text();
    throw new Error(message);
  }
}

export async function register(payload) {
  const response = await fetch(`${prefix}/auth/register`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    credentials: "include",
    body: JSON.stringify(payload),
  });
  if (!response.ok) {
    let message = await response.text();
    throw new Error(message);
  }
}

export async function logout() {
  fetch(`${prefix}/auth/logout`, {
    method: "GET",
    credentials: "include",
  });
  userState.logged = false;
  userState.username = "";
  goto("/login");
}

export async function post(formData) {
  const response = await fetch(`${prefix}/posts`, {
    method: "POST",
    body: formData,
    credentials: "include",
  });

  if (response.status === 401) {
    userState.logged = false;
    goto("/login");
  }
}

export async function feed(payload) {
  const response = await fetch(`${prefix}/feed`, {
    method: "GET",
    body: payload,
    credentials: "include",
  });

  if (response.status === 401) {
    userState.logged = false;
    goto("/login");
  }
}
