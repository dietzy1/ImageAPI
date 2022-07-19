import { useEffect, useState } from "react";
import Navbar from "../Components/Navbar";
import FrontGallery from "../Components/FrontGallery";
import Gallery from "../Components/Gallery";
import Searchbar from "../Components/Searchbar";
import Text from "../Components/Text";
import Footer from "../Components/Footer";

const path = "http://localhost:8000/api/v0/images/?tags=Happy&quantity=20";

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
      <div className="h-screen bg-[#131419]">
        <div className="container mx-auto">
          <Searchbar searchText={(text: any) => setTerm(text)} />

          {!loading && images.length === 0 && (
            <h1 className="text-5xl text-center mx-auto mt-32">
              Unable to find images :/
            </h1>
          )}

          {loading ? (
            <h1 className="text-6xl text-center mx-auto mt-32">Loading...</h1>
          ) : (
            <div className="columns-3 p-40">
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
