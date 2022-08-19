import "../index.css";

/* const image =
  "http://localhost:8000/fileserver/b40fccdb-8e06-4d55-b4bd-b5e662a46149.jpg";

const image1 =
  "http://localhost:8000/fileserver/912d6c8f-f007-4fa2-9370-e0a7ccb69717.jpg";
 */
//animate-sliding

//12 elements maximum I know works mx-4

//This is not the solution its a single large image with the other images that scrolls
function FrontGallery() {
  return (
    <div>
      <hr className="my-6 sm:mx-auto  " />
      <div className="overflow-hidden ">
        <div className="flex -mx-4 animate-slide">
          <img
            className="w-32 mx-2 self-start flex-none"
            src="/static/pepes/1.jpg"
            alt=""
          />
          <img
            className="w-32 mx-2 self-start flex-none"
            src="/static/pepes/2.jpg"
            alt=""
          />
          <img
            className="w-32 mx-2 self-start flex-none"
            src="/static/pepes/3.jpg"
            alt=""
          />
          <img
            className="w-32 mx-2 self-start flex-none"
            src="/static/pepes/4.jpg"
            alt=""
          />
          <img
            className="w-32 mx-2 self-start flex-none"
            src="/static/pepes/5.jpg"
            alt=""
          />
          <img
            className="w-32 mx-2 self-start flex-none"
            src="/static/pepes/6.jpg"
            alt=""
          />
          <img
            className="w-32 mx-2 self-start flex-none"
            src="/static/pepes/7.jpg"
            alt=""
          />
          <img
            className="w-32 mx-2 self-start flex-none"
            src="/static/pepes/8.jpg"
            alt=""
          />
          <img
            className="w-32 mx-2 self-start flex-none"
            src="/static/pepes/9.jpg"
            alt=""
          />
          <img
            className="w-32 mx-2 self-start flex-none"
            src="/static/pepes/10.jpg"
            alt=""
          />
          <img
            className="w-32 mx-2 self-start flex-none"
            src="/static/pepes/11.jpg"
            alt=""
          />
          <img
            className="w-32 mx-2 self-start flex-none"
            src="/static/pepes/12.jpg"
            alt=""
          />

          <img
            className="w-32 mx-2 self-start flex-none"
            src="/static/pepes/1.jpg"
            alt=""
          />
          <img
            className="w-32 mx-2 self-start flex-none"
            src="/static/pepes/2.jpg"
            alt=""
          />
          <img
            className="w-32 mx-2 self-start flex-none"
            src="/static/pepes/3.jpg"
            alt=""
          />
          <img
            className="w-32 mx-2 self-start flex-none"
            src="/static/pepes/4.jpg"
            alt=""
          />
          <img
            className="w-32 mx-2 self-start flex-none"
            src="/static/pepes/5.jpg"
            alt=""
          />
          <img
            className="w-32 mx-2 self-start flex-none"
            src="/static/pepes/6.jpg"
            alt=""
          />
          <img
            className="w-32 mx-2 self-start flex-none"
            src="/static/pepes/7.jpg"
            alt=""
          />
          <img
            className="w-32 mx-2 self-start flex-none"
            src="/static/pepes/8.jpg"
            alt=""
          />
          <img
            className="w-32 mx-2 self-start flex-none"
            src="/static/pepes/9.jpg"
            alt=""
          />
          <img
            className="w-32 mx-2 self-start flex-none"
            src="/static/pepes/10.jpg"
            alt=""
          />
          <img
            className="w-32 mx-2 self-start flex-none"
            src="/static/pepes/11.jpg"
            alt=""
          />
          <img
            className="w-32 mx-2 self-start flex-none"
            src="/static/pepes/12.jpg"
            alt=""
          />
        </div>
      </div>
      <hr className="my-6 sm:mx-auto  " />
    </div>
  );
}

export default FrontGallery;
