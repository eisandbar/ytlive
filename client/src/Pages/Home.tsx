import React, { ReactElement } from "react";
import { CategoryRow } from "../Components/CategoryRow";
import { StreamRow } from "../Components/StreamRow";
import { NavBar } from "../Components/Nav";

export const Home = (): ReactElement => {
  return (
    <div className="page">
      <NavBar />
      <div className="content">
        <CategoryRow gaming={false} />
        <CategoryRow gaming={true} />
        <StreamRow gaming={false} />
        <StreamRow gaming={true} />
      </div>
    </div>
  );
};
