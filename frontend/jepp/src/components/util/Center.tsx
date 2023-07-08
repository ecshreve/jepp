import React from "react";

interface Props {
  width?: string | number;
  children: React.ReactChild;
}

const Center: React.FC<Props> = ({ width, children }: Props) => (
  <div
    style={{
      display: "flex",
      flexDirection: "column",
      flexBasis: 80,
      alignItems: "center",
      justifyContent: "space-around",
      top: 0,
      bottom: 0,
      left: 0,
      right: 0,
      position: "absolute",
      width: width ? width : "100%",
      margin: "0 auto",
      backgroundColor: "transparent"
    }}
  >
    {children}
  </div>
);

export default Center;