import { useState } from "react";

const Searchbar = ({ triggerParentUpdate }: any) => {
  const [query, setQuery] = useState("");
  const onSubmit = (e: any) => {
    e.preventDefault();
    triggerParentUpdate(query);
  };

  return (
    <div className="rounded w-[600px] mx-auto  ">
      <form onSubmit={onSubmit}>
        <div className="flex items-center border-b-2  py-2 bg-white rounded-xl">
          <svg
            aria-hidden="true"
            className="w-6 h-6 ml-3 text-gray-500 dark:text-gray-400 pointer-events-none"
            fill="none"
            stroke="currentColor"
            viewBox="0 0 24 24"
            xmlns="http://www.w3.org/2000/svg"
          >
            <path
              strokeLinecap="round"
              strokeLinejoin="round"
              strokeWidth="2"
              d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"
            ></path>
          </svg>
          <input
            onChange={(e) => setQuery(e.target.value)}
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

export function searchTags() {}
