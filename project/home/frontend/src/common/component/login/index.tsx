import {Form, Input, Button, Checkbox, Card, message} from 'antd'
import React from 'react'
import {UserOutlined, LockOutlined} from '@ant-design/icons'
import {Regist} from "../register";
import axios from 'axios'
import * as account from "../../service/cred";

export interface LoginProviderProps {
    children(user: string): JSX.Element
}

export function LoginProvider(props: LoginProviderProps) {
    const [user, setUser] = React.useState('')
    const fetchUser = React.useCallback((userID) => {
        setUser(userID)
    }, [])
    if (user === '') {
        return <LoginForm onLogin={fetchUser}/>
    }
    return <div>{props.children(user)}</div>
}

interface LoginFormProps {
    onLogin(userID: string): void
}

function LoginForm(props: LoginFormProps) {
    const [register, setRegister] = React.useState(false)
    if (!!register) {
        return (<Regist success={() => {
            setRegister(false)
        }}/>)
    }
    const onFinish = (values: any) => {
        props.onLogin('lijunyi')
        axios.post('/v1/namespaces/carefree/login', {
            username: values.username,
            password: values.password,
        })
            .then(function (response) {
                console.log(response);
                // account.defaultCredsProvider.setRefreshToken(response)
            })
            .catch(function (error) {
                console.log(error);
            });
    }
    return (
        <div className="cf-login-wrap">
            <Card
                style={{
                    width: 400,
                    height: 300,
                    position: 'absolute',
                    left: '50%',
                    right: '50%',
                    margin: '300px 0 0 -200px',
                }}
            >
                <p>登录</p>
                <Form
                    name="normal_login"
                    className="login-form"
                    initialValues={{
                        remember: true,
                    }}
                    onFinish={onFinish}
                >
                    <Form.Item
                        name="username"
                        rules={[
                            {
                                required: true,
                                message: '请输入你的用户名!',
                            },
                        ]}
                    >
                        <Input
                            prefix={<UserOutlined className="site-form-item-icon"/>}
                            placeholder="用户名"
                        />
                    </Form.Item>
                    <Form.Item
                        name="password"
                        rules={[
                            {
                                required: true,
                                message: '请输入你的密码!',
                            },
                        ]}
                    >
                        <Input
                            prefix={<LockOutlined className="site-form-item-icon"/>}
                            type="password"
                            placeholder="密码"
                        />
                    </Form.Item>
                    <Form.Item>
                        <Form.Item name="remember" valuePropName="checked" noStyle>
                            <Checkbox>记住密码</Checkbox>
                        </Form.Item>
                        <a className="login-form-forgot" href="">
                            忘记密码？
                        </a>
                    </Form.Item>
                    <Form.Item>
                        <Button
                            type="primary"
                            htmlType="submit"
                            className="login-form-button"
                        >
                            登录
                        </Button>
                        &nbsp;或者 <a onClick={() => {
                        setRegister(true)
                    }}>注册!</a>
                    </Form.Item>
                </Form>
            </Card>
        </div>
    )
}
