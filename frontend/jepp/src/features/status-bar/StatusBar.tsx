import React from "react";
import { Button } from "../jepp-button/Button";

import "./StatusBar.css";

type StatusBarProps = {
  gameTitle: string;
  handleClickRestart: () => void;
  handleClickNewGame: () => void;
};

export const StatusBar = (props: StatusBarProps) => {
  return (
    <div className="status-bar">
      <div style={{ paddingTop: "3px" }}><strong>{props.gameTitle}</strong></div>
      <div>
        <Button
          style={{
            marginRight: "5px",
          }}
          onClick={props.handleClickRestart}
          label="Restart"
          size="sm"
          className="jepp-button jepp-button--sm jepp-button--custom"
        >
        </Button>
        <Button
          onClick={props.handleClickNewGame}
          label="New Game"
          size="sm"
          className="jepp-button jepp-button--sm jepp-button--custom"
        >
        </Button>
      </div>
    </div>
  );
};
