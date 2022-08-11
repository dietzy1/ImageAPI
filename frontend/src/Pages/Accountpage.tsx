export default function Accountpage() {
  return (
    <div className="text-white">
      Accountpage I want to show the user their account page I want a field
      which shows their API key I want a bottum to destroy and generate a new
      API key
    </div>
  );
}

//Need to hit the API key generation endpoints

//Function that generates a new API key

//navbar component
/*  function () {
  return fetch("http://localhost:8000/auth/deletekey", {
    method: "POST",
    credentials: "include",
  }).then((response) => {
    if (!response.ok) {
      throw new Error("Failed to logout");
    }
    setUser(false);
  });
}

function () {
  return fetch("http://localhost:8000/auth/generatekey/", {
    method: "POST",
    credentials: "include",
  }).then((response) => {
    if (!response.ok) {
      throw new Error("Failed to logout");
    }
    setUser(false);
  });
*/
