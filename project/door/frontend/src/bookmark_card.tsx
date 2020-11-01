import React, { useState } from 'react'
import './app.scss'
import { DragPanel } from './common/component/drag_panel/index'
import { Menu, Button, Card, Avatar, Carousel, Tag, List } from 'antd'

export function BookMarkCard() {
  const title = '今日愿望'
  const data = [
    {
      title: '谷歌',
      url: '//www.google.com',
      image: '',
    },
    {
      title: '360',
      url: '//www.360.com',
      image: '',
    },
    {
      title: '必应',
      url: '//www.biying.com',
      image: '',
    },
  ]
  return (
    <DragPanel width={300} height={300}>
      <Card
        title={title}
        bordered={false}
        style={{
          height: `100%`,
        }}
      >
        <div
          className={`bookmarkcard`}
          style={{
            width: `100%`,
            height: `100%`,
          }}
        >
          {data.map((value, index) => {
            return (
              <Card.Grid
                style={{
                  width: `40%`,
                  height: `60px`,
                  marginLeft: `5px`,
                }}
              >
                <Card.Meta
                  avatar={<Avatar>{value.title}</Avatar>}
                  title={value.title}
                />
              </Card.Grid>
            )
          })}
        </div>
      </Card>
    </DragPanel>
  )
}
