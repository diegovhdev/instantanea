export function elapsedtime(dateStr) {
  const date = new Date(dateStr);
  const diff = (date - Date.now()) / 1000;

  const formatter = new Intl.RelativeTimeFormat("es", {
    numeric: "auto",
  });

  const units = [
    ["year", 31536000],
    ["month", 2592000],
    ["day", 86400],
    ["hour", 3600],
    ["minute", 60],
    ["second", 1],
  ];

  for (const [unit, seconds] of units) {
    const value = Math.round(diff / seconds);

    if (Math.abs(value) >= 1) {
      return formatter.format(value, unit);
    }
  }

  return "ahora";
}
