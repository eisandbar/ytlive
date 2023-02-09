import React, { ReactElement } from "react";
import { useNavigate } from "react-router-dom";
import Nav from "react-bootstrap/Nav";
import Navbar from "react-bootstrap/Navbar";

interface NavProps {
  show?: boolean;
}

export const NavBar = ({ show = true }: NavProps): ReactElement => {
  const navigate = useNavigate();

  const onClickCategories = (): void => {
    navigate("/categories/all", { state: { gaming: false } });
  };

  const onClickStreams = (): void => {
    navigate("/streams/all", { state: { gaming: false } });
  };

  return (
    <Navbar sticky="top">
      <Navbar.Brand href="/">
        <h2 className="fw-bold YT">YT</h2>
        <h2 className="Live">Live</h2>
      </Navbar.Brand>
      <Navbar.Toggle aria-controls="basic-navbar-nav" />
      <Navbar.Collapse id="basic-navbar-nav">
        <Nav>
          <Nav.Link onClick={onClickCategories}>Categories</Nav.Link>
          <Nav.Link onClick={onClickStreams}>Streams</Nav.Link>
        </Nav>
      </Navbar.Collapse>
    </Navbar>
  );
};
