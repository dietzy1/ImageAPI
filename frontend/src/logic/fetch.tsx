export async function Loginfunc(username: string, password: string) {
  const formData = new FormData();
  formData.set("username", username);
  formData.set("password", password);

  const res = await fetch("http://localhost:8000/auth/signin/", {
    method: "POST",
    body: formData,
    credentials: "include",
  });
  if (!res.ok) {
    console.log("res not ok login");
    return Promise.resolve(false);
  }
  console.log("res ok login");
  return Promise.resolve(true);
}

export async function Signupfunc(username: string, password: string) {
  const formData = new FormData();
  formData.set("username", username);
  formData.set("password", password);

  const res = await fetch("http://localhost:8000/auth/signup/", {
    method: "POST",
    body: formData,
    credentials: "include",
  });
  if (!res.ok) {
    return Promise.resolve(false);
  }
  return Promise.resolve(true);
}

export async function refreshSessionfunc() {
  const res = await fetch("http://localhost:8000/auth/refresh/", {
    method: "POST",
    credentials: "include",
  });
  if (!res.ok) {
    return Promise.resolve(false);
  }
  return Promise.resolve(true);
}

export async function Logoutfunc() {
  const res = await fetch("http://localhost:8000/auth/signout/", {
    method: "POST",
    credentials: "include",
  });
  if (!res.ok) {
    return Promise.resolve(false);
  }
  return Promise.resolve(true);
}
