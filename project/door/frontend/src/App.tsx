import React from 'react'
import './app.scss'
import { BookMarkCard } from './bookmark_card'
import * as pb from '../../../../api/door/v1/door/door_pb'
import { Input } from 'antd'
import { AudioOutlined } from '@ant-design/icons'
import { DragPanel } from './common/component/drag_panel/index'

const { Search } = Input

function App() {
  return (
    <div className="App">
      <div className="container">
        <BookMarkCard />
        <div
          style={{
            width: `400px`,
          }}
        >
          <DragPanel width={300}>
            <Search
              placeholder="input search text"
              onSearch={(value) => window.open('https://www.baidu.com/wd=123')}
              enterButton
            />
          </DragPanel>
        </div>
      </div>
    </div>
  )
}

export default App
