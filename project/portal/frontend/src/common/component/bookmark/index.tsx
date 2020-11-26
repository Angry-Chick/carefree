import {
  List,
  Avatar,
  Button,
  Skeleton,
  Popover,
  Menu,
  Dropdown,
  Form,
  Input,
} from "antd";
import React from "react";
import * as sid from "../sidebar";
import { DragPanel, locationProps } from "../drag_panel";

interface bookmarkProps {
  index: number;
  delete: (index: number) => void;
  data: sid.sliceBookMark;
  // TODO(ljy):将结构定义在 store/api 中
  save: (index: number, bk: sid.sliceBookMark) => void;
}

export function Bookmark(props: bookmarkProps) {
  let data = props.data
  const [bmData, setBmData] = React.useState(props.data)
  const [liVisible, setLiVisible] = React.useState(false)
  const [hdVisible, setHdVisible] = React.useState(false)
  const saveLoc = (loc: locationProps) => {
    bmData.loc = loc
  };
  const menu = (
    <Menu>
      <Menu.Item
        key="1"
        onClick={() => {
          props.delete(props.index);
        }}
      >
        删除组件
      </Menu.Item>
      <Menu.Item
        key="2"
        onClick={() => {
          props.save(props.index, bmData)
        }}
      >
        保存
      </Menu.Item>
    </Menu>
  );
  return (
    <DragPanel
      save={saveLoc}
      location={{
        x: data.loc.x,
        y: data.loc.y,
        width: data.loc.width,
        height: data.loc.height,
      }}
    >
      <Dropdown overlay={menu} trigger={["contextMenu"]}>
        <div className="cf-book-mark-warp">
          <List
            header={
              <div className="cf-bookmark-header">
                {bmData.title ? bmData.title : "书签"}
                <Popover
                  visible={hdVisible}
                  className="cf-bookmark-header-edit"
                  content={
                    <AlterHeader
                      alterHeader={(header: string) => {
                        setBmData({
                          title: header,
                          items: bmData.items,
                          loc: bmData.loc,
                        })
                        setHdVisible(false)
                      }}
                    />
                  }
                  onVisibleChange={(visible) => {
                    setHdVisible(visible);
                  }}
                  title="编辑标题"
                  trigger="click"
                >
                  <Button type="link">修改</Button>
                </Popover>
              </div>
            }
            className="cf-loadmore-list"
            itemLayout="horizontal"
            dataSource={bmData.items}
            renderItem={(item: sid.bookMarkItem, index: number) => {
              return (
                <List.Item
                  actions={[
                    <Popover
                      visible={liVisible}
                      onVisibleChange={(visible) => {
                        setLiVisible(visible);
                      }}
                      content={
                        <AlterLineItem
                          alterLineItem={(
                            title: string,
                            link: string,
                          ) => {
                            let li = [...bmData.items!];
                            li[index].title = title;
                            li[index].link = link;
                            setBmData({
                              title: bmData.title,
                              items: li,
                              loc: bmData.loc,
                            })
                          }}
                        />
                      }
                      title="编辑书签"
                      trigger="click"
                    >
                      <Button style={{ padding: 0 }} size="small" type="link">
                        修改
                      </Button>
                    </Popover>,
                    <Button
                      style={{ padding: 0, marginRight: 15 }}
                      size="small"
                      type="link"
                      onClick={() => {
                        let li = [...bmData.items!];
                        li.splice(index, 1);
                        setBmData({
                          title: bmData.title,
                          items: li,
                          loc: bmData.loc,
                        })
                      }}
                    >
                      删除
                    </Button>,
                  ]}
                >
                  <Skeleton avatar title={false} loading={false} active>
                    <List.Item.Meta
                      avatar={
                        <Avatar
                          style={{
                            backgroundColor: "#333333",
                            verticalAlign: "middle",
                          }}
                        >
                          {item.title}
                        </Avatar>
                      }
                      title={<a href={item.link}>{item.title}</a>}
                      description={item.link}
                    />
                  </Skeleton>
                </List.Item>
              );
            }}
          />
          <Popover
            content={
              <AddLineItem
                addLineItem={(title: string, link: string) => {
                  const newBM: sid.bookMarkItem =
                  {
                    title: title,
                    link: link,
                  }
                  const li = bmData.items ? [...bmData.items, newBM] : [newBM]
                  setBmData({
                    title: bmData.title,
                    items: li,
                    loc: bmData.loc,
                  })
                }}
              />
            }
            title="增加书签"
            trigger="click"
          >
            <Button className="cf-bookmark-add-item" type="link">
              增加
            </Button>
          </Popover>
        </div>
      </Dropdown>
    </DragPanel >
  );
}

const AlterLineItem = (props: any) => {
  return (
    <Form
      size="small"
      layout={"horizontal"}
      onFinish={(values: any) => { props.alterLineItem(values.title, values.url,); }}
    >
      <Form.Item label="标题" name="title">
        <Input placeholder="请输入标题" />
      </Form.Item>
      <Form.Item label="网址" name="url">
        <Input placeholder="请输入url" />
      </Form.Item>
      <Form.Item>
        <Button size="small" type="primary" htmlType="submit">
          提交
        </Button>
      </Form.Item>
    </Form>
  );
};

const AlterHeader = (props: any) => {
  return (
    <Form
      size="small"
      layout={"horizontal"}
      onFinish={(values: any) => {
        props.alterHeader(values.header);
      }}
    >
      <Form.Item label="标题名称" name="header">
        <Input placeholder="请输入标题" />
      </Form.Item>
      <Form.Item>
        <Button size="small" type="primary" htmlType="submit">
          提交
        </Button>
      </Form.Item>
    </Form>
  );
};

const AddLineItem = (props: any) => {
  return (
    <Form
      size="small"
      layout={"horizontal"}
      onFinish={(values: any) => {
        props.addLineItem(values.title, values.link);
      }}
    >
      <Form.Item label="标题" name="title">
        <Input placeholder="请输入标题" />
      </Form.Item>
      <Form.Item label="网址" name="link">
        <Input placeholder="请输入 link 的网址" />
      </Form.Item>
      <Form.Item>
        <Button size="small" type="primary" htmlType="submit">
          提交
        </Button>
      </Form.Item>
    </Form>
  );
};
