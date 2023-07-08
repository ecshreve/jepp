import React from "react";
import "./App.css";
import { StatusBar } from "./features/status-bar/StatusBar";

import {useAppSelector} from "./app/hooks";

import Config from "./features/config/Config";

function App() {
  const config = useAppSelector((state) => state.config);
  // const config = useAppSelector((state) => state.config);
  return (
    <div className="App">
      {!config.gameActive ? (
        <Config />
      ) : (
        <StatusBar
          gameTitle={"hello"}
          handleClickNewGame={() => {}}
          handleClickRestart={() => {}}
        />
      )}
    </div>
  );
}

export default App;
