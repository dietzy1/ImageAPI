import { useState, useEffect } from "react";

function Footer() {
  const [healthcheck, setHealthcheck] = useState(false);

  const healthcheckfunc = async () => {
    const res = await fetch("http://localhost:8000/healthcheck");
    if (!res.ok) {
      setHealthcheck(true);
      console.log(healthcheck);
    }
  };

  useEffect(() => {
    healthcheckfunc();
  }, []);

  return (
    <footer className="shadow px-6 fixed bottom-0 w-screen bg-darky pt-3 mt-20 ">
      <div>
        <div className="text-1xl font-bold whitespace-nowrap text-white flex items-center justify-between pl-14 pr-14">
          <div className=" flex items-center">
            <svg
              xmlns="http://www.w3.org/2000/svg"
              className="h-10 w-10"
              fill="none"
              viewBox="0 0 24 24"
              stroke="currentColor"
              strokeWidth="2"
            >
              <path
                strokeLinecap="round"
                strokeLinejoin="round"
                d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z"
              />
            </svg>
            Pepe-API
          </div>
          <div className="flex flex-row relative">
            API-Status
            {healthcheck ? (
              <div className="rounded-full flex p-1 w-1 h-1 bg-red-800 bot-16 left-16" />
            ) : (
              <div className="rounded-full flex p-1 w-1 h-1 bg-greeny bot-16 left-16" />
            )}
          </div>
        </div>
      </div>
      <hr className="my-3 sm:mx-auto  " />
      <span className="block text-sm text-gray-400 sm:text-center ">
        © 2022 Pepe-API™
      </span>
    </footer>
  );
}

export default Footer;
