import { imageType } from "../Pages/Home";

//If term is empty then the API should return random images

function Gallery({ image }: { image: imageType }) {
  //const tags = image.tags.split(",");

  return (
    <div className="shadow-lg text-white object-top relative">
      <img src={image.filepath} alt="" className="w-full mb-12 rounded-lg" />
      <div className="absolute inset-0 bg-cover bg-center z-0">
        <div className="opacity-0 hover:opacity-100 duration-300 absolute inset-0 z-10 flex justify-center items-center text-3xl text-greeny font-semibold hover:backdrop-blur-sm filter backdrop-grayscale">
          <ul className="flex flex-col justify-center">
            <div className="">Name: {image.name}</div>
            <div>Tags: {image.tags}</div>
          </ul>
        </div>
      </div>
    </div>
  );
}

export default Gallery;
