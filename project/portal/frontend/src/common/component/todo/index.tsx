import { List, Button, Checkbox, Input, Menu, Dropdown } from 'antd'
import React from 'react'
import { DragPanel } from "../drag_panel";

export function ToDo(props: any) {
  const liProps: Array<string> = []
  const [lineItem, setLineItem] = React.useState(liProps)
  const menu = (
    <Menu>
      <Menu.Item
        key="1"
        onClick={() => {
          props.delete()
        }}
      >
        删除组件
      </Menu.Item>
    </Menu>
  )
  let inputValue = ''
  return (
    <DragPanel >
      <Dropdown className="cf-todo-warp" overlay={menu} trigger={['contextMenu']}>
        <div>
          <List
            header="TODO"
            className="cf-todo-list"
            itemLayout="horizontal"
            dataSource={lineItem}
            renderItem={(item: any) => {
              return (
                <List.Item
                  actions={[
                    <Button
                      className="cf-todo-delete-item"
                      size="small"
                      type="link"
                      onClick={() => {
                        let li = [...lineItem]
                        li.splice(lineItem.indexOf(item), 1)
                        setLineItem(li)
                      }}
                    >
                      删除
                  </Button>,
                  ]}
                >
                  <div className="cf-todo-list-item">
                    <Checkbox>{item}</Checkbox>
                  </div>
                </List.Item>
              )
            }}
          />
          <Input
            allowClear
            className="cf-add-todo-submit"
            placeholder="按回车提交"
            defaultValue={inputValue}
            onChange={(e) => {
              inputValue = e.target.value
            }}
            onPressEnter={(e) => {
              e.stopPropagation()
              setLineItem([...lineItem, inputValue])
            }}
          />
        </div>
      </Dropdown>
    </DragPanel>
  )
}
