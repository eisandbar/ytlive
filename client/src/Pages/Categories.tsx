import React, { ReactElement } from "react";
import { useLocation } from "react-router-dom";
import { CategoryRow } from "../Components/CategoryRow";
import { NavBar } from "../Components/Nav";

interface LocationState {
  gaming: boolean;
}

export const Categories = (): ReactElement => {
  const location = useLocation();
  const { gaming } = location.state as LocationState;
  return (
    <div className="page">
      <NavBar />
      <div className="content">
        <CategoryRow gaming={gaming} rows={-1} />
      </div>
    </div>
  );
};
