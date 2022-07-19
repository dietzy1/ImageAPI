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
    </div>
  );
}

{
  /* <div className="overflow-hidden h-[500px]">
<div className="flex -mx-4 animate-slide">
 */
}

{
  /* <img className="w-64 mx-4 self-start flex-none" src={image1} alt="" />
<img className="w-64 mx-4 self-start flex-none" src={image} alt="" />
<img className="w-64 mx-4 self-start flex-none" src={image1} alt="" />
<img className="w-64 mx-4 self-start flex-none" src={image} alt="" />
<img className="w-64 mx-4 self-start flex-none" src={image1} alt="" />
<img className="w-64 mx-4 self-start flex-none" src={image} alt="" /> */
}

/* 
slider =
  "height-[250px] mx-auto position-relative w-full display-grid place-items:center ";

slidetrack = "flex w-[4500px]";

slide = " h-[200px] w-[250px] flex center p-[15px] "

img ="w-[100%]" */

export default FrontGallery;

function tempimage() {
  return (
    <section className="overflow-hidden animate-sliding w-full">
      {/*  <div className="container px-5 py-2 mx-auto lg:pt-24 lg:px-32"> */}
      <div className="flex flex-wrap -m-1 md:-m-2">
        <div className="flex flex-wrap w-1/2">
          <div className="w-1/2 p-1 md:p-2">
            <img
              src={image}
              alt="gallery"
              className="block object-cover object-center
                  w-full h-full rounded-lg"
            ></img>
          </div>
          <div className="w-1/2 p-1 md:p-2">
            <img
              src={image1}
              alt="gallery"
              className="block object-cover object-center w-full h-full rounded-lg"
            ></img>
          </div>
          <div className="w-full p-1 md:p-2">
            <img
              src={image1}
              alt="gallery"
              className="block object-cover object-center w-full h-full rounded-lg"
            ></img>
          </div>
        </div>
        <div className="flex flex-wrap w-1/2">
          <div className="w-full p-1 md:p-2">
            <img
              src={image}
              alt="gallery"
              className="block object-cover object-center w-full h-full rounded-lg"
            ></img>
          </div>
          <div className="w-1/2 p-1 md:p-2">
            <img
              src={image1}
              alt="gallery"
              className="block object-cover object-center w-full h-full rounded-lg"
            ></img>
          </div>
          <div className="w-1/2 p-1 md:p-2">
            <img
              src={image}
              alt="gallery"
              className="block object-cover object-center w-full h-full rounded-lg"
            ></img>
          </div>
        </div>
      </div>
      {/*      </div> */}
    </section>
  );
}
