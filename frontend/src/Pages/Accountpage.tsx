import { useState } from "react";
import { createPortal } from "react-dom";

import Footer from "../Components/Footer";
import Navbar from "../Components/Navbar";
import { useGlobalState } from "../logic/context";

/* if (!open) return null;
return createPortal(
  <div
    className="top-0 bottom-0 right-0 left-0 fixed z-[1] backdrop-blur-lg shadow-3xl p-60"
    onClick={onClose}
  >
    <div
      className="flex flex-col justify-center  "
      onClick={(e) => {
        e.stopPropagation();
      }}
    >
      <form className="max-w-[450px] w-full mx-auto bg-gray-900 p-8  rounded-lg">
        <p
          onClick={onClose}
          className="flex justify-end text-white text-bold text-xl"
        >
          x
        </p> */
//I need to apply different z values
export function Accountpage({ open, onClose }: any) {
  const [apiKey, setAPIKey] = useState("");
  const [toggle, setToggle] = useState(false);
  const [state, dispatch] = useGlobalState();

  const triggerToggle = () => {
    setToggle(!toggle);
    console.log(toggle);
    FetchAPIKey({ setAPIKey });
  };

  if (!open) return null;
  return createPortal(
    <div
      className="top-0 bottom-0 right-0 left-0 fixed z-[1] backdrop-blur-lg shadow-3xl p-60"
      onClick={onClose}
    >
      <div
        className="flex flex-col justify-center  "
        onClick={(e) => {
          e.stopPropagation();
        }}
      >
        <div className="max-w-[650px] h-[650px] w-full mx-auto bg-gray-900 p-8 px-8 rounded-lg">
          <p
            className="flex justify-end text-white text-bold text-xl"
            onClick={onClose}
          >
            x
          </p>
          <h2 className="text-4xl font-bold text-white text-center">
            account page
          </h2>

          <div className="flex flex-col text-gray-400 py-2">
            User Name
            <div className="rounded-lg bg-gray-700 mt-2 p-2 focus:border-greeny focus:bg-gray-800 focus:outline-none">
              "output the fcking username here"
            </div>
          </div>
          <div className="flex flex-col text-gray-400 py-2">
            <button
              className="my-5 py-2 bg-greeny shadow-lg shadow-greeny/50 hover:shadow-greeny/30 text-white font-semibold rounded-lg"
              /* onClick={() => FetchAPIKey({ setAPIKey })} */
              onClick={triggerToggle}
            >
              Show API key
            </button>
            API-key
            <div className="rounded-lg bg-gray-700 mt-2 p-2 focus:border-greeny focus:bg-gray-800 focus:outline-none">
              {toggle ? apiKey : "****************"}
            </div>
            <span>Feel free</span>
            <span>filler text</span>
            <span>filler text</span>
            <span>filler text</span>
            <span>filler text</span>
            <button
              className="my-5 py-2 bg-greeny shadow-lg shadow-greeny/50 hover:shadow-greeny/30 text-white font-semibold rounded-lg"
              onClick={() => GenerateAPIKey({ setAPIKey })}
            >
              Generate API-key
            </button>
          </div>
        </div>
      </div>
    </div>,
    document.getElementById("portal")!
  );
}

//Function that fetches a new api key
async function FetchAPIKey({ setAPIKey }: any) {
  const res = await fetch("http://localhost:8000/auth/showkey/", {
    method: "GET",
    credentials: "include",
  });
  setAPIKey(await res.json());
}

async function GenerateAPIKey({ setAPIKey }: any) {
  const res = await fetch("http://localhost:8000/auth/generatekey/", {
    method: "POST",
    credentials: "include",
  });
  setAPIKey(res.json());
}
