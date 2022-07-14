import React from "react";
import { useEffect, useState } from "react";
import { text } from "stream/consumers";
import Appbar from "../Components/Appbar";
import FrontGallery from "../Components/FrontGallery";
import Gallery from "../Components/Gallery";
import Searchbar from "../Components/Searchbar";
import Text from "../Components/Text";

const path =
  "http://localhost:8000/image/?uuid=4e14eaf7-8fdc-4022-b3f5-aa197bf447f4";

export type imageType = {
  name: string;
  uuid: string;
  tags: Array<string>;
  created: string;
  filepath: string;
};

function Home() {
  const [images, setImages] = useState<imageType>({} as imageType);
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
  }, [term]);

  return (
    <div>
      <Appbar />
      <Text />
      <FrontGallery />
      <Searchbar searchText={(text: any) => setTerm(text)} />
      {!loading && images === null && (
        <h1 className="text-white"> No images found</h1>
      )}

      {/*   {loading ? (
        <h1 className="text-6xl text-center mx-auto mt-32">Loading...</h1>
      ) : (
        <div className="grid grid-cols-3 gap-4">
            <Searchbar key={images.uuid} image={images.filepath} />
          }
        </div>
      )} */}

      <Gallery />
    </div>
  );
}

export default Home;
