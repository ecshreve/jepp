import React from "react";
import {
  Dropdown, DropdownProps, ButtonGroup, Button
} from "react-bootstrap";
import "./Dropdown.css";

interface Props extends DropdownProps {
  options: string[];
  selection: string;
  setter: (s: string) => void;
}

const JeppDropdown = (props: Props) => {
  const items = props.options.map((o) => {
    return <Dropdown.Item onSelect={() => props.setter(o)}>{o}</Dropdown.Item>;
  });

  return (
    <Dropdown as={ButtonGroup} className="dropdown">
      <Button
        className="dropdown-val"
        variant="success"
      >
        {props.selection}
      </Button>
      <Dropdown.Toggle
        split
        variant="info"
        id="dropdown-basic"
      ></Dropdown.Toggle>
      <Dropdown.Menu>{items}</Dropdown.Menu>
    </Dropdown>
  )
};

export default JeppDropdown;
