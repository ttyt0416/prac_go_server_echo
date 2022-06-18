import React from "react";
import { Switch, Route } from "react-router";
import "./App.css";

import Header from "./components/header/header";

import Homepage from "./pages/home/home";

const App = () => {
  return (
    <div className="app">
      <Header />
      <Switch>
        <Route exact path="/" component={Homepage} />
      </Switch>
    </div>
  );
};

export default App;
