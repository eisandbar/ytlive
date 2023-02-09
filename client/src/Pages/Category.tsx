import React, { ReactElement } from "react";
import { useLocation } from "react-router-dom";
import { Scroll } from "../Components/Scroll";
import { Category } from "../util/category";
import { NavBar } from "../Components/Nav";

interface LocationState {
  category: Category;
}

export const CategoryPage = (): ReactElement => {
  const location = useLocation();
  const { category } = location.state as LocationState;
  return (
    <div className="page">
      <NavBar />
      <div className="content">
        <Scroll filters={[category.category]} />
      </div>
    </div>
  );
};
