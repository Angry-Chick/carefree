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

export function BookMarkList(props: any) {
  const [header, setHeader] = React.useState({
    title: "",
    visible: false,
  });
  const menu = (
    <Menu>
      <Menu.Item
        key="1"
        onClick={() => {
          props.delete();
        }}
      >
        删除组件
      </Menu.Item>
    </Menu>
  );
  interface liValue {
    title: string;
    description: string;
    color: string;
    visible: boolean;
  }
  const liProps: Array<liValue> = [];
  const [lineItem, setLineItem] = React.useState(liProps);
  return (
    <Dropdown overlay={menu} trigger={["contextMenu"]}>
      <div className="cf-book-mark-warp">
        <List
          header={
            <div className="cf-bookmark-header">
              {!!header.title ? header.title : "书签"}
              <Popover
                visible={header.visible}
                className="cf-bookmark-header-edit"
                content={
                  <AlterHeader
                    alterHeader={(header: string) => {
                      setHeader({ title: header, visible: false });
                    }}
                    header={header.title}
                  />
                }
                onVisibleChange={(visible) => {
                  setHeader({ title: header.title, visible: visible });
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
          dataSource={lineItem}
          renderItem={(item: any) => {
            return (
              <List.Item
                actions={[
                  <Popover
                    visible={item.visible}
                    onVisibleChange={(visible) => {
                      let li = [...lineItem];
                      li[lineItem.indexOf(item)].visible = visible;
                      setLineItem(li);
                    }}
                    content={
                      <AlterLineItem
                        alterLineItem={(
                          title: string,
                          url: string,
                          color: string
                        ) => {
                          let li = [...lineItem];
                          li[lineItem.indexOf(item)].title = title;
                          li[lineItem.indexOf(item)].description = url;
                          li[lineItem.indexOf(item)].visible = false;
                          li[lineItem.indexOf(item)].color = color;
                          setLineItem(li);
                        }}
                        title={item.title}
                        url={item.description}
                        color={item.color}
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
                      let li = [...lineItem];
                      li.splice(lineItem.indexOf(item), 1);
                      setLineItem(li);
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
                          backgroundColor: item.color,
                          verticalAlign: "middle",
                        }}
                      >
                        {item.title}
                      </Avatar>
                    }
                    title={<a href={item.description}>{item.title}</a>}
                    description={item.description}
                  />
                </Skeleton>
              </List.Item>
            );
          }}
        />
        <Popover
          content={
            <AddLineItem
              addLineItem={(title: string, url: string, color: string) => {
                let li = [
                  ...lineItem,
                  {
                    title: title,
                    description: url,
                    visible: false,
                    color: color,
                  },
                ];
                setLineItem(li);
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
  );
}

const AlterLineItem = (props: any) => {
  return (
    <Form
      size="small"
      layout={"horizontal"}
      onFinish={(values: any) => {
        console.log(values);
        props.alterLineItem(
          values.title ? values.title : props.title,
          values.url ? values.url : props.url,
          values.color ? values.color : props.color
        );
      }}
    >
      <Form.Item label="标题" name="title">
        <Input placeholder="请输入标题" defaultValue={props.title} />
      </Form.Item>
      <Form.Item label="网址" name="url">
        <Input placeholder="请输入url" defaultValue={props.url} />
      </Form.Item>
      <Form.Item label="颜色" name="color">
        <Input placeholder="设置颜色" defaultValue={props.color} />
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
        <Input placeholder="请输入标题" defaultValue={props.header} />
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
        props.addLineItem(values.title, values.url, values.color);
      }}
    >
      <Form.Item label="标题" name="title">
        <Input placeholder="请输入标题" />
      </Form.Item>
      <Form.Item label="网址" name="url">
        <Input placeholder="请输入url" />
      </Form.Item>
      <Form.Item label="颜色" name="color">
        <Input placeholder="设置颜色" />
      </Form.Item>
      <Form.Item>
        <Button size="small" type="primary" htmlType="submit">
          提交
        </Button>
      </Form.Item>
    </Form>
  );
};
