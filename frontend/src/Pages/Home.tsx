import { useEffect, useState } from "react";
import Navbar from "../Components/Navbar";
import FrontGallery from "../Components/FrontGallery";
import Gallery from "../Components/Gallery";
import Searchbar from "../Components/Searchbar";
import Text from "../Components/Text";
import Footer from "../Components/Footer";
import "../index.css";

const path = "http://localhost:8000/api/v0/images/?tags=Happy&quantity=20";

const image1 =
  "http://localhost:8000/fileserver/912d6c8f-f007-4fa2-9370-e0a7ccb69717.jpg";

export type imageType = {
  name: string;
  uuid: string;
  tags: Array<string>;
  created: string;
  filepath: string;
};

function Home() {
  const [images, setImages] = useState<imageType[]>({} as imageType[]);
  const [term, setTerm] = useState("");
  const [loading, setLoading] = useState(true);

  const getImages = async () => {
    const r = await fetch(path);
    const images = await r.json();
    setImages(images);
    setLoading(false);
    console.log(images);
  };

  useEffect(() => {
    getImages();
  }, []);

  return (
    <div>
      <div className="h-screen">
        <Navbar />
        <Text />
        <FrontGallery />
      </div>

      <div className="">
        <div className="relative border-l border-white h-full border-r ml-14 mr-14 ">
          <hr className="my-6 sm:mx-auto  " />
          <div className="container mx-auto">
            {/*vertical line element*/}

            <Searchbar searchText={(text: any) => setTerm(text)} />

            {!loading && images.length === 0 && (
              <div>
                <h1 className="text-5xl text-center mx-auto mt-32">
                  Unable to find images :/
                </h1>
                <img
                  className="w-64 mx-2 self-start flex-none"
                  src={image1}
                  alt=""
                />
              </div>
            )}

            {loading ? (
              <h1 className="text-6xl text-center mx-auto mt-32 text-white">
                Unable to find images :/
                <img className="w-64" src={image1} alt="" />
              </h1>
            ) : (
              <div className="columns-3 p-40">
                {images.map((image) => (
                  <Gallery key={image.uuid} image={image} />
                ))}
              </div>
            )}
          </div>
          <hr className="my-6 sm:mx-auto  " />
          {/*vertical line element*/}
        </div>
      </div>

      <Footer />
    </div>
  );
}

export default Home;
