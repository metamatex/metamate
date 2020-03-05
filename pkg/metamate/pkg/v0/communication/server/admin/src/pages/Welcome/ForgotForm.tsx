import React from 'react';
import {Button, Checkbox, Form, Icon, Input} from 'antd';

interface Props {
    form: any;
}

interface State {
}

class ForgotForm extends React.Component<Props, State> {
    handleSubmit(e: any) {
        e.preventDefault();
        this.props.form.validateFields((err: any, values: any) => {
            if (!err) {
                console.log('Received values of form: ', values);
            }
        });
    };

    render() {
        const {getFieldDecorator} = this.props.form;

        return (
            <div className={"email-login-form"}>
                <Form onSubmit={this.handleSubmit}>
                    <Form.Item>
                        {getFieldDecorator('email', {
                            rules: [{required: true, message: 'Please input your email!'}],
                        })(
                            <Input
                                placeholder="Email"
                            />,
                        )}
                    </Form.Item>
                    <Form.Item>
                        <Button type="primary" htmlType="submit" block>
                            Sent password
                        </Button>
                    </Form.Item>
                </Form>
            </div>
        );
    }
}

const WrappedLoginForm = Form.create({name: 'login'})(ForgotForm);

export default WrappedLoginForm;