import { useEffect, useState } from "react";
import Navbar from "../Components/Navbar";
import FrontGallery from "../Components/FrontGallery";
import Gallery from "../Components/Gallery";
import Text from "../Components/Text";
import Footer from "../Components/Footer";
import { UseAuth } from "../Components/Context";
import "../index.css";

/* const image1 =
  "http://localhost:8000/fileserver/912d6c8f-f007-4fa2-9370-e0a7ccb69717.jpg";
 */
export type imageType = {
  name: string;
  uuid: string;
  tags: Array<string>;
  created: string;
  filepath: string;
};

function Home() {
  const [images, setImages] = useState<imageType[]>({} as imageType[]);
  const [loading, setLoading] = useState(true);
  const [query, setQuery] = useState("");
  const auth = UseAuth(); //Context hook

  const updateQueryState = (state: any) => {
    setQuery(state);
    console.log(query);
  };

  //Hardcoded images to display something to gallery by default
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

  //Fetches images from the server
  const getImages = async (query: string) => {
    try {
      const res = await fetch(
        `http://localhost:8000/api/v0/images/?tags=${query}&quantity=25`
      );
      setImages(await res.json());
      setLoading(false);
    } catch (error) {
      console.log(error);
    }
  };

  useEffect(() => {
    if (query === "") {
      getImagesEmpty();
    } else getImages(query);
  }, [query]);

  return (
    <div>
      <div className="h-screen">
        <Navbar user={auth().user} triggerParentUpdate={updateQueryState} />

        <Text />
        <FrontGallery />

        <div className="">
          {"container mx-auto"}
          {!loading && images.length === 0 && (
            <div>
              <h1 className="text-5xl text-center mx-auto mt-32">
                Unable to find images 😞
              </h1>
            </div>
          )}
          {loading ? (
            <div className="pb-28">
              <h1 className="text-5xl text-center mx-auto text-white flex justify-center p-3 font-bold">
                Unable to load images 😞
              </h1>
              <div className="flex justify-center">
                <img
                  className="rounded-3xl w-64"
                  src="/static/pepes/test1.jpg"
                  alt=""
                />
              </div>
            </div>
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
