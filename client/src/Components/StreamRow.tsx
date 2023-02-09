import React, { ReactElement, useEffect, useState } from "react";
import { Category } from "../util/category";
import { host, streamMinWidth } from "../util/const";
import { Grid } from "./Grid";
import { Header } from "./Header";
import { StreamCard } from "./StreamCard";

interface RowProps {
  gaming: boolean;
  maxResults?: number;
}

export const StreamRow = (props: RowProps): ReactElement => {
  const [streams, setStreams] = useState<Category[]>([]);
  const maxResults = props.maxResults !== undefined ? props.maxResults : 30;

  useEffect(() => {
    getStreams().then(
      (res) => {},
      (rej) => {}
    );
  }, []);

  const getStreams = async (): Promise<void> => {
    try {
      const response = await fetch(
        `${host}/streams?maxResults=${maxResults}&gaming=${props.gaming.toString()}`
      );
      if (response.ok) {
        response.json().then(
          (res) => {
            setStreams(res);
          },
          (rej) => {}
        );
      }
    } catch (error) {
      console.log(error);
    }
  };

  return (
    <div className="section ">
      {props.gaming ? (
        <Header
          title="Top Gaming Streams"
          url={"/streams/gaming"}
          gaming={true}
        />
      ) : (
        <Header title="Top Streams" url={"/streams/all"} gaming={false} />
      )}
      <Grid
        minWidth={streamMinWidth}
        items={streams}
        Child={StreamCard}
        rows={2}
      />
      <div className="solid" />
    </div>
  );
};
