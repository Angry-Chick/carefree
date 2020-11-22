import React, {useState} from 'react';
import {
    Form,
    Input,
    message,
    Button, Card,
} from 'antd';
import axios from 'axios'

const formItemLayout = {
    labelCol: {
        xs: {
            span: 24,
        },
        sm: {
            span: 8,
        },
    },
    wrapperCol: {
        xs: {
            span: 24,
        },
        sm: {
            span: 16,
        },
    },
};
const tailFormItemLayout = {
    wrapperCol: {
        xs: {
            span: 24,
            offset: 0,
        },
        sm: {
            span: 16,
            offset: 8,
        },
    },
};

export function Regist(props: any) {
    const onFinish = (values: any) => {
        axios.post('/api/carefree.project.portal.v1.PortalService/SignUp', {
            username: values.username,
            password: values.password,
        }, {headers: {'Content-Type': 'application/json'}}).then(
            function (resp) {
                message.success('注册成功')
                props.success()
            }
        ).catch(
            function (error) {
                console.log(error)
            }
        )
    };
    return (
        <div className="cf-register-wrap">
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
                <p>注册</p>
                <Form
                    className='register-form'
                    {...formItemLayout}
                    name="register"
                    onFinish={onFinish}
                >
                    <Form.Item
                        name="username"
                        label={
                            <span>用户名</span>
                        }
                        rules={[
                            {
                                required: true,
                                message: '请输入用户名!',
                            },
                        ]}
                    >
                        <Input/>
                    </Form.Item>
                    <Form.Item
                        name="password"
                        label="密码"
                        rules={[
                            {
                                required: true,
                                message: '请输入密码!',
                            },
                        ]}
                        hasFeedback
                    >
                        <Input.Password/>
                    </Form.Item>

                    <Form.Item
                        name="confirm"
                        label="确认"
                        dependencies={['password']}
                        hasFeedback
                        rules={[
                            {
                                required: true,
                                message: '请再次输入密码!',
                            },
                            ({getFieldValue}) => ({
                                validator(rule, value) {
                                    if (!value || getFieldValue('password') === value) {
                                        return Promise.resolve();
                                    }

                                    return Promise.reject('The two passwords that you entered do not match!');
                                },
                            }),
                        ]}
                    >
                        <Input.Password/>
                    </Form.Item>

                    <Form.Item {...tailFormItemLayout}>
                        <Button type="primary" htmlType="submit">
                            Register
                        </Button>
                    </Form.Item>
                </Form>
            </Card>
        </div>
    );
}
