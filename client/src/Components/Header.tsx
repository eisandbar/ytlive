import React, { ReactElement } from "react";
import { useNavigate } from "react-router-dom";

interface HeaderProps {
  title: string;
  url: string;
  gaming: boolean;
}

export const Header = ({ title, url, gaming }: HeaderProps): ReactElement => {
  const navigate = useNavigate();
  const onClick = (): void => {
    navigate(url, { state: { gaming } });
  };

  return (
    <h1>
      <span className="display-6 fw-bold text-left">{title}</span>
      <span className="header-link" onClick={onClick}>
        See More
      </span>
    </h1>
  );
};
