import Footer from "../Components/Footer";
import Navbar from "../Components/Navbar";
import Sidebar from "../Components/Sidebar";

function Documentation() {
  //mt is 44max-w-[650px]
  return (
    <div>
      <Navbar />
      <Sidebar />

      <div className="flex flex-col justify-center mt-14">
        <div className="max-w-[600px] h-[650px] w-full mx-auto bg-gray-900 p-8 px-8 rounded-lg">
          <p className="flex justify-end text-white text-bold text-xl">x</p>
          <h2 className="text-4xl font-bold text-white text-center">
            Documentation
          </h2>

          <div className="flex flex-col text-gray-400 py-2">
            User Name
            <span>Feel free</span>
            <span>filler text</span>
            <span>filler text</span>
            <span>filler text</span>
            <span>filler text</span>
            <div className="rounded-lg bg-gray-700 mt-2 p-2 focus:border-greeny focus:bg-gray-800 focus:outline-none">
              "output the fcking username here"
            </div>
          </div>
          <div className="flex flex-col text-gray-400 py-2">
            API-key
            <span>Feel free</span>
            <span>filler text</span>
            <span>filler text</span>
            <span>filler text</span>
            <span>filler text</span>
          </div>
        </div>
      </div>
      <Footer />
    </div>
  );
}

export default Documentation;

//Add a indholdsfortegnelse so more content is accesible
