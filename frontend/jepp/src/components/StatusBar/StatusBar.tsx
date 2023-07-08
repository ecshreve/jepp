import React from "react";
import { Button } from "react-bootstrap";

import "./StatusBar.css";

type StatusBarProps = {
  gameTitle: string;
  handleClickRestart: () => void;
  handleClickNewGame: () => void;
};

export const StatusBar = (props: StatusBarProps) => {
  return (
    <div className="status-bar">
      <div style={{ paddingTop: "3px" }}>{props.gameTitle}</div>
      <div>
        <Button
          style={{
            marginRight: "5px",
            background: "#031297",
          }}
          onClick={props.handleClickRestart}
        >
          Restart
        </Button>
        <Button
          style={{ background: "#031297" }}
          onClick={props.handleClickNewGame}
        >
          New Game
        </Button>
      </div>
    </div>
  );
};
