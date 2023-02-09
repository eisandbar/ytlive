import React, { ReactElement } from "react";
import { useLocation } from "react-router-dom";
import { Scroll } from "../Components/Scroll";
import { NavBar } from "../Components/Nav";

interface LocationState {
  gaming: boolean;
}

export const Streams = (): ReactElement => {
  const location = useLocation();
  const { gaming } = location.state as LocationState;
  return (
    <div className="page">
      <NavBar />
      <div className="content">
        <Scroll filters={[]} gaming={gaming} maxResults={50} />
      </div>
    </div>
  );
};
