import { Link } from "react-router-dom";

function Appbar() {
  return (
    <div className="text-white flex justify-between items-center h-16 w-full mx-auto px-4 fixed top-0 z-50  bg-[#131419] ">
      {/*max w-1400px/}
      {/* 1240 */}
      <h1 className=" text-4xl font-bold text-[#00df9a] p-16">Image-Gallery</h1>
      <ul className="flex">
        {/*   <li className="p-8"> </li> */}
        <li className="p-6 hover:text-[#00df9a]">
          <Link to="/"> Home </Link>
        </li>
        <li className="p-6 hover:text-[#00df9a]">
          <Link to="/about"> About </Link>
        </li>
        <li className="p-6 hover:text-[#00df9a] pr-24">
          <Link to="/documentation"> Documentation </Link>
        </li>
      </ul>
    </div>
  );
}

export default Appbar;
