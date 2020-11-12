import { Layout, Menu, Breadcrumb, Avatar } from 'antd'
import React from 'react'
import { createFromIconfontCN } from '@ant-design/icons'
import { BookMarkList } from '../book_mark'
import { DragPanel } from '../drag_panel'
import { Note } from '../note'
import { ToDo } from '../todo'
import Axios from 'axios'

const { Sider } = Layout
const { SubMenu } = Menu

export function Sidebar(props: any) {
  const getBookMark = () => {}
  const [collapsed, setCollapsed] = React.useState(true)
  interface valueProps {
    key: string
    content: string | object
  }
  const todoProps: Array<valueProps> = []
  const [todoList, setTodoList] = React.useState(todoProps)
  const bklistProps: string[] = []
  const [bookmarkList, setBookmarkList] = React.useState(bklistProps)
  const noteListProps: Array<valueProps> = []
  const [noteList, setNoteList] = React.useState(noteListProps)
  const onCollapse = (collapsed: boolean) => {
    setCollapsed(collapsed)
  }
  return (
    <Layout style={{ minHeight: '100vh' }}>
      <Sider
        theme="light"
        collapsible
        collapsed={collapsed}
        onCollapse={onCollapse}
      >
        <div className="cf-layout-avatar">
          <Avatar
            style={{ backgroundColor: '#00a2ae', verticalAlign: 'middle' }}
            size={40}
          >
            lijunyi
          </Avatar>
        </div>
        <Menu theme="light" mode="inline">
          <SubMenu key="sub1" title="书签">
            <Menu.Item
              key="1"
              onClick={() => {
                setBookmarkList([...bookmarkList, '1'])
              }}
            >
              添加书签
            </Menu.Item>
          </SubMenu>
          <SubMenu key="sub2" title="便签">
            <Menu.Item
              key="2"
              onClick={() => {
                setNoteList([
                  ...noteList,
                  {
                    key: (+new Date()).toString(),
                    content: '请输入要记录的内容，按鼠标右键保存',
                  },
                ])
              }}
            >
              添加便签
            </Menu.Item>
          </SubMenu>
          <SubMenu key="sub3" title="代办事项">
            <Menu.Item
              key="2"
              onClick={() => {
                setTodoList([
                  ...todoList,
                  {
                    key: (+new Date()).toString(),
                    content: '',
                  },
                ])
              }}
            >
              添加代办事项
            </Menu.Item>
          </SubMenu>
        </Menu>
      </Sider>
      <div>
        {bookmarkList.map((v, i) => {
          return (
            <DragPanel width={300} height={300}>
              <BookMarkList
                delete={() => {
                  let li = [...bookmarkList]
                  li.splice(i, 1)
                  setBookmarkList(li)
                }}
              />
            </DragPanel>
          )
        })}
        {noteList.map((v, i) => {
          return (
            <DragPanel key={v.key} width={300} height={300}>
              <Note
                content={v.content}
                delete={() => {
                  let li = [...noteList]
                  li.splice(i, 1)
                  setNoteList(li)
                }}
              />
            </DragPanel>
          )
        })}
        {todoList.map((v, i) => {
          return (
            <DragPanel key={v.key} width={300} height={300}>
              <ToDo
                content={v.content}
                delete={() => {
                  let li = [...todoList]
                  li.splice(i, 1)
                  setTodoList(li)
                }}
              />
            </DragPanel>
          )
        })}
      </div>
    </Layout>
  )
}
