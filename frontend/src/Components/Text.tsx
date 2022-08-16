import { Link } from "react-router-dom";

function Text() {
  return (
    <div className="text-white text-center">
      <div className="max-w-[800px]  w-full h-fit mx-auto text-center flex-col justify-center pt-[96px] mt-20 mb-16">
        <p className="text-[#00df9a] font-bold p-2">High quality images</p>
        <h1 className="text-7xl font-bold py-2">Sorted by tags</h1>
        <h1 className="text-5xl font-bold">Fast, flexible and easy to use</h1>
        <h1 className="text-2xl font-bold text-gray-500">
          Gain access to an API of hundreds of sorted, high quality images
        </h1>
        <ul>
          <li className="bg-greeny w-[200px] rounded-md font-medium my-6 mx-auto py-3 text-black hover:bg-white shadow-lg shadow-greeny/50 hover:shadow-greeny/30">
            <Link to="/login"> Get started</Link>
          </li>
        </ul>
      </div>
    </div>
  );
}

export default Text;
