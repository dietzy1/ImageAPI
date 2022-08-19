import { useEffect, useState } from "react";
import Navbar from "../Components/Navbar";
import FrontGallery from "../Components/FrontGallery";
import Gallery from "../Components/Gallery";
import Text from "../Components/Text";
import Footer from "../Components/Footer";
import { refreshSessionfunc } from "../logic/fetch";
import { useGlobalState } from "../logic/context";

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
  const [state, dispatch] = useGlobalState();

  const updateQueryState = (state: any) => {
    console.log("updateQueryState");
    setQuery(state);
    console.log(query);
  };

  useEffect(() => {
    onsubmit();
    if (query === "") {
      getImagesEmpty();
    } else getImages(query);
  }, [query]);

  //Hardcoded images to display something to gallery by default
  const getImagesEmpty = async () => {
    try {
      const res = await fetch(
        `http://localhost:8000/api/v0/images/?tags=Happy&quantity=20&key=${process.env.REACT_APP_API_KEY}`
      );
      setImages(await res.json());
      setLoading(false);
    } catch (error) {
      console.log(error);
    }
    console.log(images);
  };

  //Fetches images from the server
  const getImages = async (query: string) => {
    try {
      const res = await fetch(
        `http://localhost:8000/api/v0/images/?tags=${query}&quantity=25&key=${process.env.REACT_APP_API_KEY}`
      );
      setImages(await res.json());
      setLoading(false);
    } catch (error) {
      console.log(error);
    }
  };

  async function onsubmit() {
    const ok = await refreshSessionfunc();
    if (ok) {
      dispatch({ user: true });
      console.log(state.user);
    }
    if (!ok) {
      dispatch({ user: false });
      console.log(state.user);
    }
  }

  return (
    <div>
      <div className="h-screen">
        <Navbar triggerParentUpdate={updateQueryState} />
        <Text />
        <FrontGallery />

        <div>
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
            <div>
              <div className="text-white text-3xl text-center">
                Use the search bar to find your favorite pepe!
              </div>
              <div className="text-greeny text-center text-xl">
                Found {images.length} pepes for the tag "{query}" sorted
                randomly
              </div>
              <div className="mt-3 justify-center flex items-center">
                <hr className="mt-3 sm:mx-auto w-[38rem]" />
              </div>

              <div className="columns-5 pt-10 px-24">
                {images.map((image) => (
                  <Gallery key={image.uuid} image={image} />
                ))}
              </div>
            </div>
          )}
        </div>
        <Footer />
      </div>
    </div>
  );
}

export default Home;
