import { imageType } from "../Pages/Home";

//If term is empty then the API should return random images

function Gallery({ image }: { image: imageType }) {
  //const tags = image.tags.split(",");

  return (
    <div className="shadow-lg text-white object-top ">
      <div className="">
        <img src={image.filepath} alt="" className="w-full mb-12 rounded-lg" />
        <div />
      </div>
    </div>
  );
}

export default Gallery;
