import { redirect } from "@sveltejs/kit";
import { clearState, userState } from "$lib/stores/global-state.svelte";
import { goto } from "$app/navigation";

const prefix = "/api";

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
  clearState();
  goto("/login");
}

export async function post(formData) {
  const response = await fetch(`${prefix}/posts`, {
    method: "POST",
    body: formData,
    credentials: "include",
  });

  if (response.status === 401) {
    clearState();
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

  if (response.status === 401) {
    clearState();
    goto("/login");
    return;
  }

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

  if (response.status === 401) {
    clearState();
    goto("/login");
    return;
  }

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

  if (response.status === 401) {
    clearState();
    goto("/login");
    return;
  }

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

  if (response.status === 401) {
    clearState();
    goto("/login");
    return;
  }

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

  if (response.status === 401) {
    clearState();
    goto("/login");
    return [];
  }

  if (!response.ok) {
    let message = await response.text();
    console.log(message);
    throw new Error(message);
  }

  if (response.status == 204) {
    return [];
  }
  const data = await response.json();
  return data;
}

export async function votePost(postId) {
  const response = await fetch(`${prefix}/posts/${postId}/vote`, {
    method: "POST",
    credentials: "include",
  });

  if (response.status === 401) {
    clearState();
    goto("/login");
    return;
  }

  if (!response.ok) {
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

  if (response.status === 401) {
    clearState();
    goto("/login");
    return;
  }

  if (!response.ok) {
    let message = await response.text();
    console.log(message);
    throw new Error(message);
  }
}

export async function followUser(userId) {
  const response = await fetch(`${prefix}/users/${userId}/follow`, {
    method: "POST",
    credentials: "include",
  });

  if (response.status === 401) {
    clearState();
    goto("/login");
    return;
  }

  if (!response.ok) {
    let message = await response.text();
    console.log(message);
    throw new Error(message);
  }
}

export async function unfollowUser(userId) {
  const response = await fetch(`${prefix}/users/${userId}/follow`, {
    method: "DELETE",
    credentials: "include",
  });

  if (response.status === 401) {
    clearState();
    goto("/login");
    return;
  }

  if (!response.ok) {
    let message = await response.text();
    console.log(message);
    throw new Error(message);
  }
}

export async function getFollowingUsers(userId) {
  const response = await fetch(`${prefix}/users/${userId}/following`, {
    method: "GET",
    credentials: "include",
  });

  if (response.status === 401) {
    clearState();
    goto("/login");
    return;
  }

  if (!response.ok) {
    let message = await response.text();
    console.log(message);
    throw new Error(message);
  }

  if (response.status == 204) {
    return [];
  }
  const data = await response.json();
  return data;
}

export async function deletePost(postId) {
  const response = await fetch(`${prefix}/posts/${postId}`, {
    method: "DELETE",
    credentials: "include",
  });

  if (response.status === 401) {
    clearState();
    goto("/login");
    return;
  }

  if (!response.ok) {
    let message = await response.text();
    console.log(message);
    throw new Error(message);
  }
}

export async function getPostsFromUser(userId) {
  const response = await fetch(`${prefix}/users/${userId}/posts`, {
    method: "GET",
    credentials: "include",
  });

  if (response.status === 401) {
    clearState();
    goto("/login");
    return;
  }

  if (!response.ok) {
    let message = await response.text();
    console.log(message);
    throw new Error(message);
  }

  if (response.status == 204) {
    return [];
  }
  const data = await response.json();
  return data;
}
