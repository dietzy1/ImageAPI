import React from "react";
import { Navigate } from "react-router-dom";

/* const Protectedroute = ({ children }: any) => {
  const { user: } = UserAuth();

  if (!user) {
    return <Navigate to="/" />;
  }
  return children;
};

export default Protectedroute;



// Send the cookie along with the request
export async function verifyUser(user: userType) {
  return fetch("http://localhost:8000/login", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      Accept: "application/json",
    },
    body: JSON.stringify(user),
  }).then((res) => res.json());
}
 */

//Ok so how it works is there is a cookie set with the name of session_token which is a keyvaluepair where the value is a uuid
//So the cookie needs to be sent along to the API and be verified for each protected route
//If the cookie is valid then the user is logged in
//
