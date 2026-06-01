export async function downloadImage(url) {
  const filename = url.split("/").at(-1);

  const res = await fetch(url);
  const blob = await res.blob();

  const blobUrl = URL.createObjectURL(blob);

  const link = document.createElement("a");
  link.href = blobUrl;
  link.download = filename;

  document.body.appendChild(link);
  link.click();
  document.body.removeChild(link);

  URL.revokeObjectURL(blobUrl);
}
