import React, { ReactElement, useState } from "react";
import Card from "react-bootstrap/Card";
import Col from "react-bootstrap/Col";
import { Stream } from "../util/stream";

const defaultSrc: string = "/purpleWide2.jpg";

interface CardProps {
  item: Stream;
}

export const StreamCard = ({ item }: CardProps): ReactElement => {
  const [overlayTitle, setOverlayTitle] = useState<String>("");

  const thumbnail: string | undefined =
    item.thumbnail !== null ? item.thumbnail : "";
  const url: string = item.id !== undefined ? item.id : "";

  return (
    <Col className="stream">
      <a className="no-style" href={`https://www.youtube.com/watch?v=${url}`}>
        <Card className="text-start stream">
          <div className="overlay-parent">
            <Card.Img
              className="stream-img rounded-0"
              variant="top"
              onLoad={({ currentTarget }) => {
                if (currentTarget.naturalHeight < 180) {
                  currentTarget.src = defaultSrc;
                  if (item.channelTitle !== undefined) {
                    setOverlayTitle(item.channelTitle);
                  }
                }
              }}
              src={thumbnail}
              alt={defaultSrc}
            />
            <Card.ImgOverlay>
              <Card.Text className="overlay-title">{overlayTitle}</Card.Text>
              <Card.Text className="overlay-text text-light">
                Views: {item.concurrentViewers}
              </Card.Text>
            </Card.ImgOverlay>
          </div>

          <Card.Body className="tb">
            <Card.Title as="h6" className="ellipses">
              {item.title}
            </Card.Title>
            <Card.Text className="ellipses">{item.channelTitle}</Card.Text>
          </Card.Body>
        </Card>
      </a>
    </Col>
  );
};
