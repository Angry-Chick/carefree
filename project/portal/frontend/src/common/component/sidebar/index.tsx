import { Layout, Menu, Avatar } from "antd";
import React from "react";
import { Bookmark } from "../bookmark";
import { DragPanel } from "../drag_panel";
import { Note } from "../note";
import { ToDo } from "../todo";
import axios from "axios";

const { Sider } = Layout;
const { SubMenu } = Menu;

export function Sidebar(props: any) {
  const [collapsed, setCollapsed] = React.useState(true);
  interface valueProps {
    key: string;
    content: string | object;
  }
  const todoProps: Array<valueProps> = [];
  const [todoList, setTodoList] = React.useState(todoProps);
  const noteListProps: Array<valueProps> = [];
  const [noteList, setNoteList] = React.useState(noteListProps);
  const [slice, setSlice] = React.useState<sliceProps>({
    name: "",
    background: "",
    bookmarks: [],
  });
  const onCollapse = (collapsed: boolean) => {
    setCollapsed(collapsed);
  };
  React.useEffect(() => {
    fetchSlice(props.user)
      .then((res) => {
        setSlice({
          name: res.name,
          background: res.background,
          bookmarks: res.bookmarks,
        });
      })
      .catch((err) => {
        console.log(err);
      });
  }, []);

  const deleteBookmark = (index: number) => {
    let li = [...slice.bookmarks!];
    li.splice(index, 1);
    updateSlice({
      name: slice.name,
      background: slice.background,
      bookmarks: li,
    });
    setSlice({
      name: slice.name,
      background: slice.background,
      bookmarks: li,
    });
  };
  const saveBookmark = (index: number, bk: sliceBookMark) => {
    let ns: sliceProps = slice;
    ns.bookmarks![index] = bk
    setSlice(ns)
    updateSlice(ns)
  };
  const createBookmark = () => {
    const newBK: sliceBookMark = {
      title: Date.now().toString(),
      loc: { x: 0, y: 0, width: 300, height: 300 },
    };
    const ns: sliceProps = {
      name: slice.name,
      background: slice.background,
      bookmarks: slice.bookmarks ? [...slice.bookmarks, newBK] : [newBK],
    };
    updateSlice(ns);
    setSlice(ns);
  };
  return (
    <Layout style={{ minHeight: "100vh" }}>
      <Sider
        theme="light"
        collapsible
        collapsed={collapsed}
        onCollapse={onCollapse}
      >
        <div className="cf-layout-avatar">
          <Avatar
            style={{ backgroundColor: "#00a2ae", verticalAlign: "middle" }}
            size={40}
          >
            lijunyi
          </Avatar>
        </div>
        <Menu theme="light" mode="inline">
          <SubMenu key="sub1" title="书签">
            <Menu.Item key="1" onClick={createBookmark}>
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
                    content: "请输入要记录的内容，按鼠标右键保存",
                  },
                ]);
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
                    content: "",
                  },
                ]);
              }}
            >
              添加代办事项
            </Menu.Item>
          </SubMenu>
        </Menu>
      </Sider>
      <div>
        {slice.bookmarks
          ? slice.bookmarks.map((v, i) => {
            return (
              <Bookmark
                key={v.title}
                data={v}
                index={i}
                delete={deleteBookmark}
                save={saveBookmark}
              />
            );
          })
          : ""}
        {noteList.map((v, i) => {
          return (
            <Note
              key={v.key}
              content={v.content}
              delete={() => {
                let li = [...noteList];
                li.splice(i, 1);
                setNoteList(li);
              }}
            />
          );
        })}
        {todoList.map((v, i) => {
          return (
            <ToDo
              key={v.key}
              content={v.content}
              delete={() => {
                let li = [...todoList];
                li.splice(i, 1);
                setTodoList(li);
              }}
            />
          );
        })}
      </div>
    </Layout >
  );
}

interface sliceProps {
  name: string;
  background?: string;
  bookmarks?: Array<sliceBookMark>;
}

export interface sliceBookMark {
  title: string;
  items?: Array<bookMarkItem>;
  loc: sliceLocation;
}

interface sliceLocation {
  x: number;
  y: number;
  width: number;
  height: number;
}

export interface bookMarkItem {
  title: string;
  image?: string;
  link: string;
}

const updateSlice = (slice: sliceProps) =>
  axios
    .post("/api/carefree.project.portal.slice.v1.SliceService/UpdateSlice", {
      slice: slice,
    })
    .then((res) => {
      return res.data;
    })
    .catch(function (error) {
      console.log(error);
    });

async function fetchSlice(userID: string) {
  try {
    const user = await getUser(`users/${userID}`);
    const sid = user.my_spaces[0];
    const slice = await getSlice(`spaces/${sid}/slices/${sid}`);
    return slice;
  } catch (err) {
    console.log(err);
  }
}

const getSlice = (sliceName: string) =>
  axios
    .post("/api/carefree.project.portal.slice.v1.SliceService/GetSlice", {
      name: sliceName,
    })
    .then((res) => {
      return res.data;
    })
    .catch(function (error) {
      console.log(error);
    });

const getUser = (userName: string) =>
  axios
    .post("/api/carefree.project.portal.user.v1.UserService/GetUser", {
      name: userName,
    })
    .then((res) => {
      return res.data;
    })
    .catch(function (error) {
      console.log(error);
    });
