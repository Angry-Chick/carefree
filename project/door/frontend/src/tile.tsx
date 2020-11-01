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

function Tile() {
  const data = [
    '百度',
    '谷歌',
    '360浏览器',
    '火狐',
    '凤凰网',
    '必应',
    '今日头条',
    '最强大脑',
    '门户',
  ]
  return (
    <div className="door-tile-group">
      {data.map((name, index) => {
        return (
          <div>
            <div
              className={`tile-${index + 1}`}
              style={{
                position: `relative`,
                width: `100%`,
                height: `100%`,
                borderRadius: `4px`,
                // boxShadow: `0 7px 14px 0 rgba(60, 66, 87, 0.1), 0 3px 6px 0 rgba(0, 0, 0, .07)`,
              }}
            >
              <div
                style={{
                  position: `absolute`,
                  width: `100%`,
                  bottom: '5px',
                  textAlign: `center`,
                }}
              >
                {name}
              </div>
            </div>
          </div>
        )
      })}
    </div>
  )
}
