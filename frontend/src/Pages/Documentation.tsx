import Footer from "../Components/Footer";
import Navbar from "../Components/Navbar";
import Sidebar from "../Components/Sidebar";

function Documentation() {
  //mt is 44max-w-[650px]
  return (
    <div>
      <Navbar />
      <Sidebar />

      <div className="flex flex-col justify-center pt-24 ">
        <div className="max-w-[850px] w-full mx-auto p-8 px-8 ">
          <h2 className="text-4xl font-bold text-white text-center mb-8">
            API-Documentation
          </h2>

          <h2 className="text-4xl font-bold text-white underline">
            Getting started
          </h2>
          <div className="relative border-l border-gray-200 dark:border-gray-700 pl-4 mb-10 mt-2">
            <div className="flex flex-col text-gray-400 py-2">
              <div className="absolute w-3 h-3 rounded-full mt-1.5 -left-1.5 border dark:border-gray-900 dark:bg-gray-700" />
              <span className="mb-1 text-sm font-normal leading-none text-gray-500">
                Access
              </span>

              <span className="mb-4">
                To access the API, please create an developer account first.
              </span>

              <span className="mb-1 text-sm font-normal leading-none text-gray-500">
                Access
              </span>
              <div className="absolute w-3 h-3 rounded-full mt-1.5 -left-1.5 border dark:border-gray-900 dark:bg-gray-700" />

              <span className="mb-4">
                <div className="absolute w-3 h-3 rounded-full mt-1.5 -left-1.5 border dark:border-gray-900 dark:bg-gray-700" />
                Once you have been successfully registered as a user, go ahead
                and click the account button at the top right which leads to the
                account page.
              </span>

              <span>
                <div className="absolute w-3 h-3 rounded-full mt-1.5 -left-1.5 border dark:border-gray-900 dark:bg-gray-700" />
                At the account page you are able to see the custom generated
                API-key for your account. This key is required to access the
                API. Incase your key has been compromised, you can generate a
                new one at any time, however only one key can be active at a
                time.
              </span>
              <span className="mb-4">
                The API is currently in beta and is subject to change.
              </span>
            </div>
            <h2 className="text-4xl font-bold text-white underline">
              Data model
            </h2>
            <div className="flex flex-col text-gray-400 py-2">
              <div className="absolute w-3 h-3 rounded-full mt-1.5 -left-1.5 border dark:border-gray-900 dark:bg-gray-700" />
              <span>{model}</span>
              <span>filler text</span>
              <span>filler text</span>
              <span>filler text</span>
            </div>
            <h2 className="text-4xl font-bold text-white underline">API-Key</h2>
            <div className="flex flex-col text-gray-400 py-2">
              <span>Feel free</span>
              <span>filler text</span>
              <span>filler text</span>
              <span>filler text</span>
              <span>filler text</span>
            </div>
            <h2 className="text-4xl font-bold text-white underline">
              Endpoints
            </h2>
            <div className="flex flex-col text-gray-400 py-2">
              <span>localhost:8000/api/v0/</span>
              <span>localhost:8000/auth/</span>
              <span>localhost:8000/api/v0/image/</span>
              <span>{endpoint}</span>
            </div>
            <div className="mt-24" />
          </div>
        </div>
      </div>
      <Footer />
    </div>
  );
}

export default Documentation;

const endpoint = "localhost:8000/api/v0/images/{tags}/{quantity}";
const endpoint1 = "localhost:8000/api/v0/images/{}/{quantity}";
const model = `export type imageType = {
  name: string;
  uuid: string;
  tags: Array<string>;
  created: string;
  filepath: string;
};`;
