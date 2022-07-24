function Footer() {
  return (
    <footer className="p-4 rounded-lg shadow px-6 py-8">
      <div className="flex items-center mb-4">
        <div className="self-center text-1xl font-bold whitespace-nowrap text-white ml-14">
          <div className="flex items-center ">
            <svg
              xmlns="http://www.w3.org/2000/svg"
              className="h-10 w-10"
              fill="none"
              viewBox="0 0 24 24"
              stroke="currentColor"
              stroke-width="2"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z"
              />
            </svg>
            Pepe-API
          </div>
        </div>
      </div>
      <hr className="my-6 sm:mx-auto  " />
      <span className="block text-sm text-gray-400 sm:text-center ">
        © 2022 Pepe-API™
      </span>
    </footer>
  );
}

export default Footer;
