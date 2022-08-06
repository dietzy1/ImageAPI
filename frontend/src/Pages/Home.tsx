import { useEffect, useState } from "react";
import Navbar from "../Components/Navbar";
import FrontGallery from "../Components/FrontGallery";
import Gallery from "../Components/Gallery";
import Text from "../Components/Text";
import Footer from "../Components/Footer";
import "../index.css";

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

  const getImagesEmpty = async () => {
    try {
      const res = await fetch(
        `http://localhost:8000/api/v0/images/?tags=Happy&quantity=20`
      );
      setImages(await res.json());
      setLoading(false);
    } catch (error) {
      console.log(error);
    }
  };

  const getImages = async () => {
    try {
      const res = await fetch(
        `http://localhost:8000/api/v0/images/?tags=${term}&quantity=25`
      );
      setImages(await res.json());
      setLoading(false);
    } catch (error) {
      console.log(error);
    }
  };

  useEffect(() => {
    if (term === "") {
      getImagesEmpty();
    } else getImages();
  }, []);

  return (
    <div>
      <div className="h-screen">
        <Navbar />
        <Text />
        <FrontGallery />

        <div className="container mx-auto">
          {!loading && images.length === 0 && (
            <div>
              <h1 className="text-5xl text-center mx-auto mt-32">
                Unable to find images :/
              </h1>
            </div>
          )}
          {loading ? (
            <h1 className="text-6xl text-center mx-auto mt-16 text-white">
              Unable to find images :/
              <img className="w-64" src={image1} alt="" />
            </h1>
          ) : (
            <div className="columns-5 p-20">
              {images.map((image) => (
                <Gallery key={image.uuid} image={image} />
              ))}
            </div>
          )}
        </div>
        <Footer />
      </div>
    </div>
  );
}

export default Home;
