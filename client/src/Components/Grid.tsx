import React, { ReactElement, useEffect, useRef, useState } from "react";
import Col from "react-bootstrap/Col";
import Row from "react-bootstrap/Row";

interface GridProps {
  minWidth: number;
  items: any[];
  Child: React.ElementType;
  rows: number;
}

export const Grid = (props: GridProps): ReactElement => {
  const [cols, setCols] = useState<number>(0);
  const ref = useRef<HTMLDivElement>(null);

  // Gets the width of the div to see how many cards we can fit
  useEffect(() => {
    const setWidth = (): void => {
      if (ref.current != null) {
        setCols(Math.floor(ref.current.offsetWidth / props.minWidth));
      } else {
        setCols(0);
      }
    };

    setWidth();

    const handleResize = (): void => {
      setWidth();
    };

    window.addEventListener("resize", handleResize);
  }, [ref.current]);

  const displayed = props.rows > 0 ? props.rows * cols : undefined;
  const categories = props.items.slice(0, displayed);

  //   Cols start to misbehave when a lot less than specified in row
  //  So we fix that by adding in empty columns
  let hidden: any[] = [];
  if (cols !== 0 && categories.length % cols !== 0) {
    hidden = Array.apply(null, Array(cols - (categories.length % cols)));
  }

  return (
    <div ref={ref} className="category-row ovh">
      {/* We need to specify sm in cases where displayed.length < cols */}
      <Row className="g-2" sm={cols}>
        {categories.map((x, i) => (
          <props.Child key={i} item={x} />
        ))}

        {hidden.map((x, i) => (
          <Col key={i}></Col>
        ))}
      </Row>
    </div>
  );
};
