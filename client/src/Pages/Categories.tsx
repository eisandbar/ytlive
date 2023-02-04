import React, { ReactElement } from "react";
import { useLocation } from "react-router-dom";
import { CategoryRow } from "../Components/CategoryRow";

interface LocationState {
  gaming: boolean;
}

export const Categories = (): ReactElement => {
  const location = useLocation();
  const { gaming } = location.state as LocationState;
  return (
    <div className="page">
      <div className="content">
        <CategoryRow gaming={gaming} rows={-1} />
      </div>
    </div>
  );
};
