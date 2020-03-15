import React from 'react';
import {Button, Form, Input} from 'antd';

interface Props {
    form: any;
}

interface State {
}

class TokenLoginForm extends React.Component<Props, State> {
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
            <div className={"token-login-form"}>
                <Form onSubmit={this.handleSubmit}>
                    <Form.Item>
                        {getFieldDecorator('token', {
                            rules: [{required: true, message: 'Please input your Password!'}],
                        })(
                            <Input
                                placeholder="Token"
                            />,
                        )}
                    </Form.Item>
                    <Form.Item>
                        <Button type="primary" htmlType="submit" block>
                            Log in
                        </Button>
                    </Form.Item>
                </Form>
            </div>
        );
    }
}

const WrappedTokenLoginForm = Form.create({name: 'login'})(TokenLoginForm);

export default WrappedTokenLoginForm;