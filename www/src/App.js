import React, { Component } from 'react';
import './App.css';

class App extends Component {
  render() {
    return (
      <div className="App">
        <h1>Logging testing application</h1>
        <div>Links</div>
        <ul>
          <li><a className="link" href="/users">Users</a></li>
          <li><a className="link" href="/todos">Todos</a></li>
          <li><a className="link" href="/admin">Admin</a></li>
        </ul>
      </div>
    );
  }
}

export default App;
