import Footer from "../Components/Footer";
import Navbar from "../Components/Navbar";

function Documentation() {
  return (
    <div className="h-screen">
      <Navbar />
      <div className="text-white text-center pt-[96px] mt-20">
        {/*   <Navbar /> */}
        <h1>Welcome to the API documentation.</h1>
        <div className="bg-gray-400 ml-40 mr-40 h-screen rounded-3xl text-xl">
          <div>Primary api route is localhost:8000/api/v0/</div>
          <div>Request a single Image: localhost:8000/api/v0/image/</div>
          <div>Request multiple images: localhost:8000/api/v0/images/</div>

          <span></span>
        </div>
        <Footer />
      </div>
    </div>
  );
}

export default Documentation;

//Add a indholdsfortegnelse so more content is accesible
