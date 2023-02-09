import React, { ReactElement, useEffect, useState, useRef } from "react";
import { Grid } from "./Grid";
import InfiniteScroll from "react-infinite-scroll-component";
import { Stream } from "../util/stream";
import { host, streamMinWidth } from "../util/const";
import { StreamCard } from "./StreamCard";

interface ScrollProps {
  filters?: string[];
  gaming?: boolean;
  maxResults?: number;
}

export const Scroll = (props: ScrollProps): ReactElement => {
  const [offset, setOffset] = useState(0);
  const [items, setItems] = useState<Stream[]>([]);
  const [hasMore, setHasMore] = useState(true);
  const ref = useRef<HTMLDivElement>(null);
  const maxResults = props.maxResults !== undefined ? props.maxResults : 20;

  const getItems = async (): Promise<void> => {
    try {
      const response = await fetch(genUrl(props, maxResults, offset));
      if (response.ok) {
        response.json().then(
          (res: Stream[]) => {
            setItems(items.concat(res));
            setOffset(offset + res.length);
            if (res.length === 0) {
              setHasMore(false);
            }
          },
          (rej) => {}
        );
      }
    } catch (error) {
      console.log(error);
    }
  };

  // This hook makes sure that enough items are loaded that we can scroll
  useEffect(() => {
    const checkHeight = (): void => {
      if (
        ref.current != null &&
        ref.current.scrollHeight < window.innerHeight
      ) {
        getItems().then(
          (res) => {},
          (rej) => {}
        );
      }
    };

    checkHeight();

    const handleResize = (): void => {
      checkHeight();
    };

    window.addEventListener("resize", handleResize);
  }, [ref.current, items]);

  return (
    <div className="section full-height">
      <h1 className="header-title fw-bold text-left">LiveStreams</h1>
      <div ref={ref}>
        <InfiniteScroll
          dataLength={items.length}
          next={getItems}
          hasMore={hasMore}
          loader={<></>}
        >
          <Grid
            items={items}
            Child={StreamCard}
            minWidth={streamMinWidth}
            rows={-1}
          />
        </InfiniteScroll>
      </div>
    </div>
  );
};

const genUrl = (
  props: ScrollProps,
  maxResults: number,
  offset: number
): string => {
  let url: string = `http://${host}/streams?maxResults=${maxResults}&offset=${offset}`;

  if (props.gaming === true) {
    url += "&gaming=true";
  }

  if (props.filters !== undefined && props.filters?.length > 0) {
    props.filters.forEach((x, i) => {
      url += "&filters=" + encodeURIComponent(x);
    });
  }

  return encodeURI(url);
};
