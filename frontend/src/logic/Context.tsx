import { useContext, useState } from "react";
import { createContext } from "react";

const authContext = createContext(UseProvideAuth); //maybe needs to be some other default value

//Wrapper that exposes the authContext to the app
export function ProvideAuth({ children }: any) {
  /*  const auth = UseProvideAuth(); */
  return (
    <authContext.Provider value={UseProvideAuth}>
      {children}
    </authContext.Provider>
  );
}

//Hooks to access the authContext
export const UseAuth = () => {
  return useContext(authContext);
};

function UseProvideAuth() {
  const [user, setUser] = useState(true); //false if not logged in -- //True if logged in
  const [errors, setErrors] = useState([]);

  //Loginform component
  function login(username: string, password: string) {
    const formData = new FormData();
    formData.set("username", username);
    formData.set("password", password);

    return fetch("http://localhost:8000/auth/signin/", {
      method: "POST",
      body: formData,
      credentials: "include",
      //Important the request must come from localhost:3000 and not localhost:3000? -Quite litterally with ? aswell
    }).then((response) => {
      if (!response.ok) {
        throw new Error("Failed to login");
      }
      setUser(true);
    });
  }

  //navbar component
  function logout() {
    return fetch("http://localhost:8000/auth/", {
      method: "POST",
      credentials: "include",
    }).then((response) => {
      if (!response.ok) {
        throw new Error("Failed to logout");
      }
      setUser(false);
    });
  }

  //Registerform component
  function signup(username: string, password: string) {
    const formData = new FormData();
    formData.set("username", username);
    formData.set("password", password);

    return fetch("http://localhost:8000/auth/signin/", {
      method: "POST",
      body: formData,
      credentials: "include",
      //Important the request must come from localhost:3000 and not localhost:3000? -Quite litterally with ? aswell
    }).then((response) => {
      if (!response.ok) {
        throw new Error("Failed to login");
      }
      setUser(true);
    });
  }
  //Home component
  function refreshSession() {
    return fetch("http://localhost:8000/auth/refresh/", {
      method: "POST",
      credentials: "include",
    }).then((response) => {
      if (!response.ok) {
        setUser(false);
        throw new Error("Failed to refresh session");
      }
      setUser(true);
    });
  }

  return {
    user,
    setUser,
    errors,
    setErrors,

    login,
    logout,
    signup,
    refreshSession,
  };
}

/* function UseProvideAuth1() {
  const [user, setUser] = useState(null);
  const [isLoading, setIsLoading] = useState(true);
  const [errors, setErrors] = useState([]);

  function login() {}
  function logout() {}
  function signup() {}
  function refreshSession() {}
  return {
    user,
    isLoading,
    errors,
    login,
    logout,
    signup,
    refreshSession,
  };
} */
