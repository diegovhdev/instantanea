import { getPosts } from "$lib/services/api";

export async function load() {
  const data = await getPosts("id");
  console.log(data);

  return {
    data,
  };
}
