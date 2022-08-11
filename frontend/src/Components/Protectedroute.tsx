import React from "react";
import { Navigate } from "react-router-dom";
import { UseAuth } from "../logic/Context";

const Protectedroute = ({ children }: any) => {
  const auth = UseAuth(); //Context hook

  if (!auth().user) {
    return <Navigate to="/" />;
  }
  return children;
};

export default Protectedroute;

//Ok so how it works is there is a cookie set with the name of session_token which is a keyvaluepair where the value is a uuid
//So the cookie needs to be sent along to the API and be verified for each protected route
//If the cookie is valid then the user is logged in
//
