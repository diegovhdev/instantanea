import { getPosts } from "$lib/services/api";

export async function load() {
  const posts = await getPosts("favorites");

  return {
    posts,
  };
}
