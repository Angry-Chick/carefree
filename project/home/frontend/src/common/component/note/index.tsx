import React from 'react'
import { Menu, Dropdown } from 'antd'

export function Note(props: any) {
  const colorList = ['lightpink', 'coral', 'aliceblue', 'gold', '#C9A39C']
  const [editable, setEditable] = React.useState(true)
  const [color, setColor] = React.useState(colorList[0])
  const menu = (
    <Menu>
      <Menu.Item
        key="1"
        onClick={() => {
          setEditable(!editable)
        }}
      >
        {(function (editable: boolean) {
          if (!!editable) {
            return '保存'
          } else {
            return '编辑'
          }
        })(editable)}
      </Menu.Item>
      <Menu.Item
        key="2"
        onClick={() => {
          const index = colorList.indexOf(color)
          setColor(
            index < colorList.length - 1 ? colorList[index + 1] : colorList[0]
          )
        }}
      >
        改变背景颜色
      </Menu.Item>
      <Menu.Item
        key="2"
        onClick={() => {
          props.delete()
        }}
      >
        删除组件
      </Menu.Item>
    </Menu>
  )
  return (
    <Dropdown overlay={menu} trigger={['contextMenu']}>
      <div
        className="cf-note-wrap"
        contentEditable={editable}
        style={{ background: color }}
      >
        {props.content}
      </div>
    </Dropdown>
  )
}
