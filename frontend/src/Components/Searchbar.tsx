import { useState } from "react";

const Searchbar = ({ searchText }: any) => {
  const [text, setText] = useState("");

  const onSubmit = (e: any) => {
    e.preventDefault();
    searchText(text);
    console.log(text);
  };

  return (
    <div className="max-w-lg rounded overflow-hidden my-10 mx-auto justify-center ">
      <h1 className="text-white mx-auto text-center text-5xl font-bold py-6">
        Search by tag!
      </h1>

      <form onSubmit={onSubmit} className="w-full">
        <div className="flex items-center border-b-2  py-2 bg-white rounded-xl ml-4">
          <svg
            aria-hidden="true"
            className="w-6 h-6 text-gray-500 dark:text-gray-400"
            fill="none"
            stroke="currentColor"
            viewBox="0 0 24 24"
            xmlns="http://www.w3.org/2000/svg"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"
            ></path>
          </svg>
          <input
            onChange={(e) => setText(e.target.value)}
            className="appearance-none bg-transparent border-none w-full text-gray-500 mr-2 py-1 px-2 leading-tight focus:outline-none rounded-md"
            type="text"
            placeholder="Search Image Tag"
          />

          <button
            className="flex-shrink-0 focus:outline-none focus:ring-greeny font-medium  bg-greeny hover:bg-teal-700 mr-4 hover:border-teal-700  text-white py-1 px-2 rounded-lg text-md"
            type="submit"
          >
            Search
          </button>
        </div>
      </form>
    </div>
  );
};
export default Searchbar;
