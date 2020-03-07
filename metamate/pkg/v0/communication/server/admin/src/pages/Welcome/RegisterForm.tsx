import React from 'react';
import {Button, Form, Input, message} from 'antd';
import Client from '../../Client';
import Sdk from './../../Sdk_';
import {FormComponentProps} from 'antd/es/form';


interface UserFormProps extends FormComponentProps {
    client: Client;
}

interface Props {
    form: any;
}

interface State {
}

class RegisterForm extends React.Component<UserFormProps, State> {
    constructor(props: UserFormProps) {
        super(props);

        this.state = {}

        this.handleSubmit = this.handleSubmit.bind(this)
    }

    handleSubmit(e: any) {
        console.log("hey")
        e.preventDefault();
        this.props.form.validateFields((err: any, values: any) => {
            if (err) {
                message.error(err)

                return
            }

            var req: Sdk.PostClientAccountsRequest = {
                clientAccounts: [
                    {
                        alternativeIds: [
                            {
                                kind: "email",
                                email: {
                                    value: values.email,
                                },
                            }
                        ],
                        password: {
                            value: values.password,
                        },
                    },
                ]
            }

            this.props.client.PostClientAccounts(req).then((rsp) => console.log(rsp))
        });
    };

    render() {
        const {getFieldDecorator} = this.props.form;

        return (
            <div className={"register-form"}>
                <Form onSubmit={this.handleSubmit}>
                    <Form.Item>
                        {getFieldDecorator('email', {
                            rules: [{required: true, message: 'Please input your email!'}],
                            initialValue: "ph.woerdehoff@gmail.com",
                        })(
                            <Input
                                placeholder="Email"
                            />,
                        )}
                    </Form.Item>
                    <Form.Item>
                        {getFieldDecorator('password', {
                            rules: [{required: true, message: 'Please input your Password!'}],
                            initialValue: "passwort1",
                        })(
                            <Input
                                type="password"
                                placeholder="Password"
                            />,
                        )}
                    </Form.Item>
                    <Form.Item>
                        {getFieldDecorator('password-repeat', {
                            rules: [{required: true, message: 'Please input your Password!'}],
                            initialValue: "passwort1",
                        })(
                            <Input
                                type="password"
                                placeholder="Repeat password"
                            />,
                        )}
                    </Form.Item>
                    <Form.Item>
                        <Button type="primary" htmlType="submit" block>
                            Create account
                        </Button>
                    </Form.Item>
                </Form>
            </div>
        );
    }
}

const WrappedLoginForm = Form.create<UserFormProps>({name: 'login'})(RegisterForm);

export default WrappedLoginForm;