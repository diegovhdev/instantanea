import { getPostsFromUser } from "$lib/services/api";

export async function load({ params }) {
  const posts = await getPostsFromUser(params.id);

  return {
    posts,
  };
}
