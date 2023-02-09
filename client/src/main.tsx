import React, { ReactElement } from "react";

import "bootstrap/dist/css/bootstrap.min.css";
import ReactDOM from "react-dom/client";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import { Home } from "./Pages/Home";
import "./style.css";
import { CategoryPage } from "./Pages/Category";
import { Streams } from "./Pages/Streams";
import { Categories } from "./Pages/Categories";

const Routing = (): ReactElement => {
  return (
    <Router>
      <Routes>
        <Route path="/category/*" element={<CategoryPage />} />
        <Route path="/streams/*" element={<Streams />} />
        <Route path="/categories/*" element={<Categories />} />
        <Route path="/" element={<Home />} />
      </Routes>
    </Router>
  );
};

ReactDOM.createRoot(document.getElementById("root") as HTMLElement).render(
  <React.StrictMode>
    <Routing />
  </React.StrictMode>
);
