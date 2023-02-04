import React, { ReactElement, useState } from "react";
import { Card, Col } from "react-bootstrap";
import { Stream } from "../util/stream";

const defaultSrc: string = "/purpleWide2.jpg";

interface CardProps {
  item: Stream;
}

export const StreamCard = ({ item }: CardProps): ReactElement => {
  const [overlayTitle, setOverlayTitle] = useState<String>("");

  const url: string | undefined = item.thumbnail !== null ? item.thumbnail : "";

  return (
    <Col className="stream">
      <Card className="text-start stream">
        <div className="overlay-parent">
          <Card.Img
            className="stream-img"
            variant="top"
            onLoad={({ currentTarget }) => {
              if (currentTarget.naturalHeight < 180) {
                currentTarget.src = defaultSrc;
                if (item.channelTitle !== undefined) {
                  setOverlayTitle(item.channelTitle);
                }
              }
            }}
            src={url}
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
          <Card.Text>{item.channelTitle}</Card.Text>
        </Card.Body>
      </Card>
    </Col>
  );
};
