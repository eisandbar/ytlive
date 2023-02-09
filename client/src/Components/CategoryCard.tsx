import React, { ReactElement } from "react";
import { Card, Col } from "react-bootstrap";
import { useNavigate } from "react-router-dom";
import { Category } from "../util/category";
import { isEmpty } from "../util/string";

const defaultSrc: string = "/purple3.jpg";

interface CardProps {
  item: Category;
}

export const CategoryCard = ({ item }: CardProps): ReactElement => {
  const navigate = useNavigate();
  const onClick = (): void => {
    navigate(`/category/${item.category}`, { state: { category: item } });
  };

  return (
    <Col className="category">
      <Card className="text-start category" onClick={onClick}>
        <div className="overlay-parent ">
          <Card.Img
            src={!isEmpty(item.url) ? item.url : defaultSrc}
            className="category-img rounded-0"
          ></Card.Img>

          <Card.ImgOverlay>
            <Card.Text className="overlay-title">
              {!isEmpty(item.url) ? "" : item.category}
            </Card.Text>
          </Card.ImgOverlay>
        </div>
        <Card.Body className="tb">
          <Card.Title as="h6" className="ellipses">
            {item.category}
          </Card.Title>
          <Card.Text className="ellipses">
            {item.concurrentViewers} viewers
          </Card.Text>
        </Card.Body>
      </Card>
    </Col>
  );
};
