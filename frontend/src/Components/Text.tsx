import { Link } from "react-router-dom";
import { useState } from "react";
import { createPortal } from "react-dom";
import { useGlobalState } from "../logic/context";
import { Loginfunc } from "../logic/fetch";
import { openLoginpage } from "./Navbar";

function Text() {
  const [isOpen, setIsOpen] = useState(false);
  return (
    <div className="text-white text-center">
      <div className="max-w-[800px]  w-full h-fit mx-auto text-center flex-col justify-center pt-[96px] mt-20 mb-12">
        <p className="text-[#00df9a] font-bold p-2">High quality images</p>
        <h1 className="text-7xl font-bold py-2">Sorted by tags</h1>
        <h1 className="text-5xl font-bold">Fast, flexible and easy to use</h1>
        <h1 className="text-2xl font-bold text-gray-500">
          Gain access to an API of hundreds of sorted, high quality images
        </h1>
        <button
          onClick={() => setIsOpen(true)}
          className="bg-greeny w-[200px] rounded-md font-medium mt-6 mx-auto py-3 text-black hover:bg-white shadow-lg shadow-greeny/50 hover:shadow-greeny/30"
        >
          Get started
        </button>
        {/*     <ul>
          <li className="bg-greeny w-[200px] rounded-md font-medium mt-6 mx-auto py-3 text-black hover:bg-white shadow-lg shadow-greeny/50 hover:shadow-greeny/30">
            <Link to="/login"> Get started</Link>
          </li>
        </ul> */}
        <Loginform open={isOpen} onClose={() => setIsOpen(false)} />
      </div>
    </div>
  );
}

export default Text;

export function Loginform({ open, onClose }: any) {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [state, dispatch] = useGlobalState();

  const onsubmitfunc = async (e: any) => {
    e.preventDefault();
    console.log("Logging in");
    const ok = await Loginfunc(username, password);
    if (ok) {
      dispatch({ user: true });
      console.log(state.user);
    }
  };

  if (!open) return null;
  return createPortal(
    <div className="fixed z-10">
      {/*  <Navbar /> */}
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
                Sign in
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
              <div className=" text-gray-400 py-2">
                <p className="flex items-center">
                  <input className="mr-2" type="checkbox" /> Remember Me
                </p>
              </div>
              <button
                className="w-full my-5 py-2 bg-greeny shadow-lg shadow-greeny/50 hover:shadow-greeny/30 text-white font-semibold rounded-lg"
                type="submit"
                onClick={onsubmitfunc}
              >
                Sign in
              </button>
            </form>
          </div>
        </div>
      </div>
    </div>,
    document.getElementById("portal")!
  );
}
