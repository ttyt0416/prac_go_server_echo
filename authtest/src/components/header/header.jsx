import React from "react";
import "./header.css";

import { Link } from "react-router-dom";

const Header = () => {
  return (
    <header className="header">
      <nav>
        <ul>
          <li>
            <Link to="/">User Test</Link>
          </li>
          <li>
            <Link to="/auth">Auth Test</Link>
          </li>
        </ul>
      </nav>
    </header>
  );
};

export default Header;
