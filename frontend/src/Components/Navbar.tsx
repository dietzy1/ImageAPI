import { Link } from "react-router-dom";

function Appbar() {
  return (
    <div className="text-white flex justify-between items-center h-24 max-w-[1400px] mx-auto px-4">
      {/* 1240 */}
      <h1 className="w-full text-5xl font-bold text-[#00df9a]">
        Image-Gallery
      </h1>
      <ul className="flex">
        {/*   <li className="p-8"> </li> */}
        <li className="p-6 hover:text-[#00df9a]">
          <Link to="/"> Home </Link>
        </li>
        <li className="p-6 hover:text-[#00df9a]">
          <Link to="/about"> About </Link>
        </li>
        <li className="p-6 hover:text-[#00df9a]">
          <Link to="/documentation"> Documentation </Link>
        </li>
      </ul>
    </div>
  );
}

export default Appbar;
