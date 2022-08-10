import { Link } from "react-router-dom";
import { useState } from "react";
import { Loginform } from "./Loginform";
import { Signupform } from "./Signupform";
import Searchbar from "./Searchbar";
import { UseAuth } from "./Context";

/* const handleLogin () => {}
const handleLogout () => setUser(null); */

/* function Navbar({stateChanger:any}) { */

const Navbar = ({ triggerParentUpdate }: any) => {
  const [isOpen, setIsOpen] = useState(false);
  const [isOpen1, setIsOpen1] = useState(false);
  const auth = UseAuth(); //Context hook

  return (
    <div>
      <nav className="text-white flex justify-center items-center h-16 w-full px-[6rem] fixed top-0 backdrop-blur-lg backdrop-grayscale backdrop-opacity-[99] backdrop-sepia bg-blacky/30 z-[999] maw-w-[1400px] ">
        <svg
          xmlns="http://www.w3.org/2000/svg"
          className="h-[5rem] w-[5rem]"
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
        <h1 className=" text-5xl font-bold text-greeny pr-8">Pepe-Gallery</h1>
        <div className="flex items-center justify-center pl-3">
          <Searchbar triggerParentUpdate={triggerParentUpdate} />
        </div>
        <ul className="flex">
          <li className="p-6 hover:text-greeny">
            <Link to="/"> Home </Link>
          </li>
          <li className="p-6 hover:text-greeny pr-12">
            <Link to="/documentation"> Documentation </Link>
          </li>
        </ul>
        <div className="flex items-center space-x-3">
          <button
            className="p-6 py-3 px-3 hover:text-greeny"
            onClick={() => setIsOpen(true)}
          >
            Login
          </button>
          <button
            className="p-6 py-3 px-3 bg-greeny rounded-xl hover:text-greeny hover:bg-white shadow-lg shadow-greeny/50 hover:shadow-greeny/30"
            onClick={() => setIsOpen1(true)}
          >
            {auth().user ? "Logged in" : "Sign up"}
          </button>
        </div>
      </nav>
      <div>
        <Loginform open={isOpen} onClose={() => setIsOpen(false)} />
        <Signupform open={isOpen1} onClose={() => setIsOpen1(false)} />
      </div>
    </div>
  );
};

export default Navbar;
