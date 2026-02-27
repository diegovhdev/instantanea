const prefix = "http://localhost:8080";

export async function login(payload) {
  try {
    const response = await fetch(`${prefix}/login`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(payload),
    });
  } catch (error) {}
}

export async function register(payload) {
  try {
    const response = await fetch(`${prefix}/register`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(payload),
    });
  } catch (error) {}
}
