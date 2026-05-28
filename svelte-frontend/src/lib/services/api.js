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
  const data = await response.json();
  return data;
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

export async function patchProfilePicture(formData) {
  const response = await fetch(
    `${prefix}/users/${userState.id}/profile-picture`,
    {
      method: "PATCH",
      body: formData,
      credentials: "include",
    },
  );

  if (!response.ok) {
    let message = await response.text();
    throw new Error(message);
  }

  const data = await response.json();

  if (!data.pictureURL) {
    let message = "Ocurrio un error en el servidor";
    throw new Error(message);
  }

  userState.profilePictureUrl = data.pictureURL;
}

export async function patchUsername(data) {
  const response = await fetch(`${prefix}/users/${userState.id}/username`, {
    method: "PATCH",
    headers: {
      "Content-Type": "application/json",
    },
    credentials: "include",
    body: JSON.stringify(data),
  });

  if (!response.ok) {
    let message = await response.text();
    throw new Error(message);
  }

  userState.username = data.username;
}

export async function patchEmail(data) {
  const response = await fetch(`${prefix}/users/${userState.id}/email`, {
    method: "PATCH",
    headers: {
      "Content-Type": "application/json",
    },
    credentials: "include",
    body: JSON.stringify(data),
  });

  if (!response.ok) {
    let message = await response.text();
    throw new Error(message);
  }
}

export async function postPassword(data) {
  const response = await fetch(`${prefix}/users/${userState.id}/password`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    credentials: "include",
    body: JSON.stringify(data),
  });

  if (!response.ok) {
    let message = await response.text();
    throw new Error(message);
  }
}

export async function deleteUser(data) {
  const response = await fetch(`${prefix}/users/${userState.id}`, {
    method: "DELETE",
    headers: {
      "Content-Type": "application/json",
    },
    credentials: "include",
    body: JSON.stringify(data),
  });

  if (!response.ok) {
    let message = await response.text();
    throw new Error(message);
  } else {
    logout();
  }
}

export async function getPosts(orderedBy) {
  const response = await fetch(`${prefix}/posts?ordered-by=${orderedBy}`, {
    method: "GET",
    credentials: "include",
  });

  if (!response.ok) {
    let message = await response.text();
    console.log(message);
    throw new Error(message);
  }
  const data = await response.json();
  return data;
}

export async function votePost(postId) {
  const response = await fetch(`${prefix}/posts/${postId}/vote`, {
    method: "POST",
    credentials: "include",
  });

  if (!response.ok) {
    console.log("aaa");
    let message = await response.text();
    console.log(message);
    throw new Error(message);
  }
}

export async function unvotePost(postId) {
  const response = await fetch(`${prefix}/posts/${postId}/vote`, {
    method: "DELETE",
    credentials: "include",
  });

  if (!response.ok) {
    console.log("aaa");
    let message = await response.text();
    console.log(message);
    throw new Error(message);
  }
}
