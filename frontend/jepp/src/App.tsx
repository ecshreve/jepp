import React from 'react';
import './App.css';
import { StatusBar } from './components/StatusBar/StatusBar';

function App() {
  return (
    <div className="App">
      <header className="App-header">app header</header>
      <StatusBar gameTitle={"hello"} handleClickNewGame={()=>{}} handleClickRestart={()=>{}}/>
    </div>
  );
}

export default App;
