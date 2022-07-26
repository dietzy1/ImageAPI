import { Link } from "react-router-dom";
import { useState } from "react";

/* const handleLogin () => {}
const handleLogout () => setUser(null); */

function Appbar() {
  const [user, setUser] = useState(null); //Need to declare the type of user here
  return (
    <div className="text-white flex justify-between items-center h-16 w-full mx-auto px-4 fixed top-0 z-50 ">
      {/*max w-1400px/}
      {/* 1240 */}
      <h1 className=" text-4xl font-bold text-greeny p-16">Image-Gallery</h1>
      <ul className="flex">
        {/*   <li className="p-8"> </li> */}
        <li className="p-6 hover:text-greeny">
          <Link to="/"> Home </Link>
        </li>
        <li className="p-6 hover:text-greeny">
          <Link to="/about"> Github </Link>
        </li>
        <li className="p-6 hover:text-greeny pr-24">
          <Link to="/documentation"> Documentation </Link>
        </li>

        <li className="p-6 hover:text-greeny">
          <Link to="/login"> Login </Link>
        </li>
        <button className=" text-white hover:text-greeny whitespace-nowrap mx-auto">
          {user ? "Logged in" : "Sign up"}
        </button>
      </ul>
    </div>
  );
}

export default Appbar;
