import ReactDOM from "react-dom";
import { useState } from "react";

export function Signupform({ open, onClose }: any) {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  if (!open) return null;

  return ReactDOM.createPortal(
    <div
      className="top-0 bottom-0 right-0 left-0 fixed z-[1] backdrop-blur-lg shadow-3xl p-60"
      onClick={onClose}
    >
      <div
        className=" z-[1]"
        onClick={(e) => {
          e.stopPropagation();
        }}
      >
        <div className="flex flex-col justify-center ">
          <form className="max-w-[450px] w-full mx-auto bg-gray-900 p-8 px-8 rounded-lg">
            <p
              onClick={onClose}
              className="flex justify-end text-white text-bold text-xl"
            >
              x
            </p>
            <h2 className="text-4xl font-bold text-white text-center">
              Sign up
            </h2>
            <div className="flex flex-col text-gray-400 py-2">
              <label className="text-start" htmlFor="username">
                User Name
              </label>
              <input
                className="rounded-lg bg-gray-700 mt-2 p-2 focus:border-greeny focus:bg-gray-800 focus:outline-none"
                type="text"
                onChange={(e) => setUsername(e.target.value)}
              />
            </div>
            <div className="flex flex-col text-gray-400 py-2">
              <label className="text-start" htmlFor="password">
                Password
              </label>
              <input
                className="rounded-lg bg-gray-700 mt-2 p-2 focus:border-greeny focus:bg-gray-800 focus:outline-none"
                type="password"
                onChange={(e) => setPassword(e.target.value)}
              />
            </div>
            {/*   <div className=" text-gray-400 py-2">
              <p className="flex items-center">
                <input className="mr-2" type="checkbox" /> Remember Me
              </p>
            </div> */}
            <button
              className="w-full my-5 py-2 bg-greeny shadow-lg shadow-greeny/50 hover:shadow-greeny/30 text-white font-semibold rounded-lg"
              type="submit"
              onClick={() => {
                registerUser(username, password);
              }}
            >
              Sign in
            </button>
          </form>
        </div>
      </div>
    </div>,
    document.getElementById("portal")!
  );
}

export async function registerUser(username: string, password: string) {
  const formData = new FormData();
  formData.set("username", username);
  formData.set("password", password);
  return await fetch("http://localhost:8000/auth/signup/", {
    method: "POST",
    body: formData,
    credentials: "include",
  });
}
