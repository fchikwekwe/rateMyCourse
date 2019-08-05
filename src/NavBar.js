import React from "react";
import { Link } from "react-router-dom";
import "./NavBar.css";

const NavBar = props => {
  return (
    <div>
      <ul>
        <Link to="/" className="navbar__link">
          <li>Home</li>
        </Link>
        <Link to="/login" className="navbar__link">
          <li>Login</li>
        </Link>
        <Link to="/signup" className="navbar__link">
          <li>Sign Up</li>
        </Link>
      </ul>
    </div>
  );
};

export default NavBar;
