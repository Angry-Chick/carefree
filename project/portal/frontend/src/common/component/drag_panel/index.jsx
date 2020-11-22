import React from 'react'

export function DragPanel(props) {
  const [moveTrack, setMoveTrack] = React.useState({ x: 0, y: 0 })
  const [scaleTrack, setScaleTrack] = React.useState({
    x: props.width === 0 ? 300 : props.width,
    y: props.height === 0 ? 150 : props.height,
  })
  const [lastPoint, setLastPoint] = React.useState({ x: 0, y: 0 })
  const [move, setMove] = React.useState(false)
  const [scale, setScale] = React.useState(false)    
    window.onmousemove = (e) => {
      if (!move && !scale) return
      const dx = e.clientX - lastPoint.x
      const dy = e.clientY - lastPoint.y
      setLastPoint({ x: e.clientX, y: e.clientY })
      move && setMoveTrack({ x: moveTrack.x + dx, y: moveTrack.y + dy })
      scale && setScaleTrack({ x: scaleTrack.x + dx, y: scaleTrack.y + dy })
    }
    window.onmouseup = (e) => {
      setMove(false)
      setScale(false)
    }
  return (
    <div
      className="door-dp-container"
      onMouseDown={(e) => {
          e.stopPropagation()
          setMove(true)
          setLastPoint({ x: e.clientX, y: e.clientY })
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
            e.stopPropagation()
            setScale(true)
            setLastPoint({ x: e.clientX, y: e.clientY })
        }}
      />
      {props.children}
    </div>
  )
}
