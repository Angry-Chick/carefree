import React, { useState } from 'react'
import './app.scss'
import { DragPanel } from './common/component/drag_panel/index'
import { BookMarkCard } from './bookmark_card'
import { Menu, Button, Card, Avatar, Carousel, Tag, List } from 'antd'
import {
  AppstoreOutlined,
  MenuUnfoldOutlined,
  MenuFoldOutlined,
  PieChartOutlined,
  DesktopOutlined,
  ContainerOutlined,
  MailOutlined,
} from '@ant-design/icons'

export function Nav() {
  const [collapsed, setCollapsed] = useState(false)
  return (
    <div style={{ width: `200px` }}>
      <Button
        size="large"
        type="text"
        onClick={() => setCollapsed(!collapsed)}
        style={{
          position: 'absolute',
          top: `4px`,
          left: `4px`,
          marginBottom: 16,
        }}
      >
        {React.createElement(collapsed ? AppstoreOutlined : AppstoreOutlined)}
      </Button>
      <Menu
        style={{
          display: `${collapsed ? 'block' : 'none'}`,
          position: 'absolute',
          top: `100px`,
          left: `0px`,
        }}
        defaultSelectedKeys={['1']}
        defaultOpenKeys={['sub1']}
        mode="vertical-left"
        theme="dark"
      >
        <Menu.Item key="1">Option 1</Menu.Item>
        <Menu.Item key="2" icon={<DesktopOutlined />}>
          Option 2
        </Menu.Item>
        <Menu.Item key="3" icon={<ContainerOutlined />}>
          Option 3
        </Menu.Item>
      </Menu>
    </div>
  )
}
