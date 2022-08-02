import { Link } from "react-router-dom";
import { useState } from "react";
import Searchbar from "./Searchbar";
import LoginModal from "./Loginform";

/* const handleLogin () => {}
const handleLogout () => setUser(null); */

function Navbar() {
  const [user, setUser] = useState(null); //Need to declare the type of user here
  const [isOpen, setIsOpen] = useState(false);

  return (
    <div>
      <nav className="text-white flex justify-between items-center h-16 w-full mx-auto px-4 fixed top-0 z-1000 ">
        {/*max w-1400px/}
      {/* 1240 */}
        <h1 className=" text-4xl font-bold text-greeny p-16">Pepe-Gallery</h1>
        <ul className="flex">
          {/*   <li className="p-8"> </li> */}
          <li className="p-4 hover:text-greeny">
            <Link to="/"> Home </Link>
          </li>
          <li className="p-4 hover:text-greeny pr-12">
            <Link to="/documentation"> Documentation </Link>
          </li>
          {/*ultag*/}
          <div className="flex items-center space-x-3">
            <button
              className="py-3 px-3 hover:text-greeny"
              onClick={() => setIsOpen(true)}
            >
              Login
            </button>
            <a
              href=""
              className="py-3 px-3 bg-greeny rounded-xl hover:text-greeny hover:bg-white shadow-lg shadow-greeny/50 hover:shadow-greeny/30"
            >
              {user ? "Logged in" : "Sign up"}
            </a>
          </div>
        </ul>
      </nav>
      <div className="">
        <LoginModal open={isOpen} onClose={() => setIsOpen(false)} />
      </div>
    </div>
  );
}

export default Navbar;
