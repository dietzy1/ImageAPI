import React from "react";

function Text() {
  return (
    <div className="text-white text-center">
      <div className="max-w-[800px] mt-[-96px] w-full h-screen  mx-auto text-center flex flex-col justify-center">
        <p className="text-[#00df9a] font-bold p-2">High quality images</p>
        <h1 className="text-7xl font-bold py-2">Sorted by tags</h1>
        <h1 className="text-5xl font-bold">Fast, flexible and easy to use</h1>
        <h1 className="text-2xl font-bold text-gray-500">
          Gain access to an API of hundreds of sorted, high quality images
        </h1>
        <button className="bg-[#00df9a] w-[200px] rounded-md font-medium my-6 mx-auto py-3 text-black hover:bg-white">
          Get Started
        </button>
      </div>
    </div>
  );
}

export default Text;
