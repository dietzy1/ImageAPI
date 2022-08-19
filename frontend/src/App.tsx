import Home from "./Pages/Home";

import React from "react";

//import { ReactDOM } from "react";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";

//These components are not loaded until you go to their path
const Documentation = React.lazy(() => import("./Pages/Documentation"));

const Loading = () => <p>Loading ...</p>;

function App() {
  return (
    <div>
      <React.Suspense fallback={<Loading />}>
        <Router>
          <Routes>
            <Route path="/" element={<Home />} />
            <Route path="/documentation" element={<Documentation />} />
          </Routes>
        </Router>
      </React.Suspense>
    </div>
  );
}

export default App;
