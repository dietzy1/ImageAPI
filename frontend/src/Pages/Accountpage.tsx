import { useState } from "react";

import Footer from "../Components/Footer";
import Navbar from "../Components/Navbar";
import { useGlobalState } from "../logic/context";

export default function Accountpage() {
  const [apiKey, setAPIKey] = useState("");
  const [toggle, setToggle] = useState(false);
  const [state, dispatch] = useGlobalState();

  const triggerToggle = () => {
    setToggle(!toggle);
    console.log(toggle);
    FetchAPIKey({ setAPIKey });
  };

  /*  if (!state.user) {
    window.location.href = "/";
  } */
  return (
    <div>
      <Navbar />
      <div className="flex flex-col justify-center mt-44">
        <div className="max-w-[650px] h-[650px] w-full mx-auto bg-gray-900 p-8 px-8 rounded-lg">
          <p className="flex justify-end text-white text-bold text-xl">x</p>
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
      <Footer />
    </div>
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
