import { Loginfunc } from "../logic/fetch";
import { useState } from "react";
import { Link } from "react-router-dom";
import Footer from "../Components/Footer";
import Navbar from "../Components/Navbar";
import { useGlobalState } from "../logic/context";

//List of shit that needs to be correct
//Must be of this format
//onClick={onsubmitfunc}

//And the onsubmit must be of const onsubmitfunc = async (e: any) => {

function Loginpage() {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [state, dispatch] = useGlobalState();

  const onsubmitfunc = async (e: any) => {
    e.preventDefault();
    console.log("Logging in");
    const ok = await Loginfunc(username, password);
    if (ok) {
      dispatch({ user: true });
      console.log(state.user);
    }
  };

  //The navbar component is picking up the onsubmit function
  return (
    <div>
      <Navbar />
      <div className="flex flex-col justify-center mt-44">
        <form className="max-w-[450px] w-full mx-auto bg-gray-900 p-8 px-8 rounded-lg">
          <li className="flex justify-end text-white text-bold text-xl">
            <Link to="/"> X </Link>
          </li>
          <h2 className="text-4xl font-bold text-white text-center">Login</h2>
          <div className="flex flex-col text-gray-400 py-2">
            <label className="text-start" htmlFor="username">
              User Name
            </label>
            <input
              className="rounded-lg bg-gray-700 mt-2 p-2 focus:border-greeny focus:bg-gray-800 focus:outline-none"
              type="text"
              onChange={(e) => setUsername(e.target.value)}
            />
          </div>
          <div className="flex flex-col text-gray-400 py-2">
            <label className="text-start" htmlFor="password">
              Password
            </label>
            <input
              className="rounded-lg bg-gray-700 mt-2 p-2 focus:border-greeny focus:bg-gray-800 focus:outline-none"
              type="password"
              onChange={(e) => setPassword(e.target.value)}
            />
          </div>
          <div className=" text-gray-400 py-2">
            <p className="flex items-center">
              <input className="mr-2" type="checkbox" /> Remember Me
            </p>
          </div>
          <button
            className="w-full my-5 py-2 bg-greeny shadow-lg shadow-greeny/50 hover:shadow-greeny/30 text-white font-semibold rounded-lg"
            type="submit"
            onClick={onsubmitfunc}
          >
            Login
          </button>
        </form>
      </div>
      <Footer />
    </div>
  );
}

export default Loginpage;

/* 
function Login(username: string, password: string) {
  console.log("login called client side");
  const formData = new FormData();
  formData.set("username", username);
  formData.set("password", password);
  globalUser.user = true;
  fetch("http://localhost:8000/auth/signin/", {
    method: "POST",
    body: formData,
    credentials: "include",
    //Important the request must come from localhost:3000 and not localhost:3000? -Quite litterally with ? aswell
  }).then((response) => {
    console.log(response);
    if (!response.ok) {
      console.log("failed to login");
      return null;
    }
    console.log("set user to true");
    globalUser.user = true;
  });
} */
