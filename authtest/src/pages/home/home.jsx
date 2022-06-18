import React, { useState } from "react";
import "./home.css";

import axios from "axios";

const Homepage = () => {
  const [id, setId] = useState();
  const [password, setPassword] = useState();
  const [searchValue, setSearchValue] = useState();
  const [searchResult, setSearchResult] = useState();
  const [deleteValue, setDeleteValue] = useState();

  const onChange = (event) => {
    const {
      target: { name, value },
    } = event;
    if (name === "id") {
      setId(value);
    } else if (name === "password") {
      setPassword(value);
    }
  };

  const handleSearchValue = (event) => {
    const {
      target: { value },
    } = event;
    setSearchValue(value);
  };

  const handleDeleteValue = (event) => {
    const {
      target: { value },
    } = event;
    setDeleteValue(value);
  };

  const onSubmit = async (event) => {
    event.preventDefault();
    const user = { name: id, password: password };
    await axios.post("http://localhost:1323/users", user).then((response) => {
      console.log(response);
    });
  };

  const searchUserById = async (event) => {
    event.preventDefault();
    await axios
      .get(`http://localhost:1323/users/${searchValue}`)
      .then((response) => {
        console.log(response);
      });
  };

  const deleteUserById = async (event) => {
    event.preventDefault();
    await axios
      .delete(`http://localhost:1323/users/${deleteValue}`)
      .then((response) => {
        console.log(response);
      });
  };

  return (
    <div className="homepage">
      <div className="homepage__auth-container">
        <form className="homepage__auth" onSubmit={onSubmit}>
          <h1 className="homepage__auth-title">Create User Test</h1>
          <input type="text" placeholder="ID" name="id" onChange={onChange} />
          <input
            type="password"
            placeholder="PASSWORD"
            name="password"
            onChange={onChange}
          />
          <input type="submit" />
        </form>
      </div>
      <div className="homepage__auth-container">
        <form className="homepage__auth" onSubmit={searchUserById}>
          <h1 className="homepage__auth-title">Search User Test</h1>
          <input
            type="text"
            placeholder="ID"
            name="id"
            onChange={handleSearchValue}
          />
          <input type="submit" />
        </form>
      </div>
      <div className="homepage__auth-container">
        <form className="homepage__auth" onSubmit={deleteUserById}>
          <h1 className="homepage__auth-title">Delete User Test</h1>
          <input
            type="text"
            placeholder="ID"
            name="id"
            onChange={handleDeleteValue}
          />
          <input type="submit" />
        </form>
      </div>
    </div>
  );
};

export default Homepage;
