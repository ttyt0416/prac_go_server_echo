import React, {useState} from 'react';
import './App.css';

import axios from 'axios'

const App = () => {
  const [id, setId] = useState('');
  const [password, setPassword] = useState('');

  const onChange = (event) => {
    const {
      target: {name, value}
    } = event;
    if (name === 'id') {
      setId(value)
    } else if (name === 'password') {
      setPassword(value)
    }
  }

  const onSubmit = async (event) => {
    event.preventDefault();
    const user = {name: id, password: password}
    axios.post(
      'http://localhost:1323/users',
      user
    )
  }

  return (
    <div className='homepage'>
      <form className='homepage__auth' onSubmit={onSubmit}>
        <input type='name' placeholder='ID' name='email' onChange={onChange} />
        <input type='password' placeholder='PASSWORD' name='password' onChange={onChange} />
        <input type='submit' />
      </form>
    </div>
  )
}

export default App;
