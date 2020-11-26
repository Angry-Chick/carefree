import React from "react";

export interface dragPanelProps {
  save?: (loc: locationProps) => void;
  location?: locationProps
  children: JSX.Element
}

export interface locationProps {
  x: number;
  y: number;
  width: number;
  height: number;
}

export function DragPanel(props: dragPanelProps) {
  let loc: locationProps = {
    width: 300,
    height: 300,
    x: 0,
    y: 0,
  }
  if (!!props.location) {
    loc.width = props.location.width
    loc.height = props.location.height
    loc.x = props.location.x
    loc.y = props.location.y

  }
  const [moveTrack, setMoveTrack] = React.useState({
    x: loc.x,
    y: loc.y,
  });
  const [scaleTrack, setScaleTrack] = React.useState({
    x: loc.width,
    y: loc.height,
  });
  const [lastPoint, setLastPoint] = React.useState({ x: 0, y: 0 });
  const [move, setMove] = React.useState(false);
  const [scale, setScale] = React.useState(false);
  window.onmousemove = (e: any) => {
    if (!move && !scale) return;
    const dx = e.clientX - lastPoint.x;
    const dy = e.clientY - lastPoint.y;
    setLastPoint({ x: e.clientX, y: e.clientY });
    move && setMoveTrack({ x: moveTrack.x! + dx, y: moveTrack.y! + dy });
    scale && setScaleTrack({ x: scaleTrack.x + dx, y: scaleTrack.y + dy });
  };
  window.onmouseup = (e: any) => {
    setMove(false);
    setScale(false);
    if (!!props.save) {
      props.save({
        width: scaleTrack.x,
        height: scaleTrack.y,
        x: moveTrack.x,
        y: moveTrack.y,
      });
    }
  };
  return (
    <div
      className="door-dp-container"
      onMouseDown={(e) => {
        e.stopPropagation();
        setMove(true);
        setLastPoint({ x: e.clientX, y: e.clientY });
      }}
      style={{
        transform: `translate(${moveTrack.x}px,${moveTrack.y}px)`,
        width: `${scaleTrack.x}px`,
        height: `${scaleTrack.y}px`,
      }}
    >
      <div
        className="door-dp-bottom-right"
        onMouseDown={(e) => {
          e.stopPropagation();
          setScale(true);
          setLastPoint({ x: e.clientX, y: e.clientY });
        }}
      />
      {props.children}
    </div>
  );
}
