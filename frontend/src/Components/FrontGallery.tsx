import React from "react";
import "../index.css";

const image =
  "http://localhost:8000/fileserver/b40fccdb-8e06-4d55-b4bd-b5e662a46149.jpg";

const image1 =
  "http://localhost:8000/fileserver/912d6c8f-f007-4fa2-9370-e0a7ccb69717.jpg";

//animate-sliding

//12 elements maximum I know works mx-4

//This is not the solution its a single large image with the other images that scrolls
function FrontGallery() {
  return (
    <div className="">
      <hr className="my-6 sm:mx-auto  " />
      <div className="overflow-hidden pb-4 ">
        <div className="flex -mx-4 animate-slide">
          <img className="w-32 mx-2 self-start flex-none" src={image1} alt="" />
          <img className="w-32 mx-2 self-start flex-none" src={image} alt="" />
          <img className="w-32 mx-2 self-start flex-none" src={image1} alt="" />
          <img className="w-32 mx-2 self-start flex-none" src={image} alt="" />
          <img className="w-32 mx-2 self-start flex-none" src={image1} alt="" />
          <img className="w-32 mx-2 self-start flex-none" src={image} alt="" />
          <img className="w-32 mx-2 self-start flex-none" src={image1} alt="" />
          <img className="w-32 mx-2 self-start flex-none" src={image} alt="" />
          <img className="w-32 mx-2 self-start flex-none" src={image1} alt="" />
          <img className="w-32 mx-2 self-start flex-none" src={image} alt="" />
          <img className="w-32 mx-2 self-start flex-none" src={image1} alt="" />
          <img className="w-32 mx-2 self-start flex-none" src={image} alt="" />

          <img className="w-32 mx-2 self-start flex-none" src={image1} alt="" />
          <img className="w-32 mx-2 self-start flex-none" src={image} alt="" />
          <img className="w-32 mx-2 self-start flex-none" src={image1} alt="" />
          <img className="w-32 mx-2 self-start flex-none" src={image} alt="" />
          <img className="w-32 mx-2 self-start flex-none" src={image1} alt="" />
          <img className="w-32 mx-2 self-start flex-none" src={image} alt="" />
          <img className="w-32 mx-2 self-start flex-none" src={image1} alt="" />
          <img className="w-32 mx-2 self-start flex-none" src={image} alt="" />
          <img className="w-32 mx-2 self-start flex-none" src={image1} alt="" />
          <img className="w-32 mx-2 self-start flex-none" src={image} alt="" />
          <img className="w-32 mx-2 self-start flex-none" src={image1} alt="" />
          <img className="w-32 mx-2 self-start flex-none" src={image} alt="" />
        </div>
      </div>
      <div className="overflow-hidden pb-4">
        <div className="flex -mx-4 animate-slide">
          <img className="w-32 mx-2 self-start flex-none" src={image1} alt="" />
          <img className="w-32 mx-2 self-start flex-none" src={image} alt="" />
          <img className="w-32 mx-2 self-start flex-none" src={image1} alt="" />
          <img className="w-32 mx-2 self-start flex-none" src={image} alt="" />
          <img className="w-32 mx-2 self-start flex-none" src={image1} alt="" />
          <img className="w-32 mx-2 self-start flex-none" src={image} alt="" />
          <img className="w-32 mx-2 self-start flex-none" src={image1} alt="" />
          <img className="w-32 mx-2 self-start flex-none" src={image} alt="" />
          <img className="w-32 mx-2 self-start flex-none" src={image1} alt="" />
          <img className="w-32 mx-2 self-start flex-none" src={image} alt="" />
          <img className="w-32 mx-2 self-start flex-none" src={image1} alt="" />
          <img className="w-32 mx-2 self-start flex-none" src={image} alt="" />

          <img className="w-32 mx-2 self-start flex-none" src={image1} alt="" />
          <img className="w-32 mx-2 self-start flex-none" src={image} alt="" />
          <img className="w-32 mx-2 self-start flex-none" src={image1} alt="" />
          <img className="w-32 mx-2 self-start flex-none" src={image} alt="" />
          <img className="w-32 mx-2 self-start flex-none" src={image1} alt="" />
          <img className="w-32 mx-2 self-start flex-none" src={image} alt="" />
          <img className="w-32 mx-2 self-start flex-none" src={image1} alt="" />
          <img className="w-32 mx-2 self-start flex-none" src={image} alt="" />
          <img className="w-32 mx-2 self-start flex-none" src={image1} alt="" />
          <img className="w-32 mx-2 self-start flex-none" src={image} alt="" />
          <img className="w-32 mx-2 self-start flex-none" src={image1} alt="" />
          <img className="w-32 mx-2 self-start flex-none" src={image} alt="" />
        </div>
      </div>
      <hr className="my-6 sm:mx-auto  " />
    </div>
  );
}

export default FrontGallery;
