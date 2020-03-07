import React from 'react';
import {Button, Checkbox, Col, Form as AntForm, Input, Row} from 'antd';
import Client from '../../../Client';
import Sdk from '../../../Sdk_';
import {FormComponentProps} from 'antd/es/form';

interface SerivceAccountsFormProps extends FormComponentProps {
    client: Client;
}

interface Props {
    form: any;
}

interface State {
    loading: boolean
    services: Sdk.Service[]
}

class Form extends React.Component<SerivceAccountsFormProps, State> {
    constructor(props: SerivceAccountsFormProps) {
        super(props);

        this.load = this.load.bind(this);
        this.submit = this.submit.bind(this);

        this.state = {
            loading: false,
            services: [],
        };

        this.load()
    }

    load() {
        this.setState({loading: true});

        this.props.client.GetServices({
            filter: {
                endpoints: {
                    postServiceAccounts: {
                        set: true,
                    }
                }
            },
        }).then((rsp: Sdk.GetServicesResponse) => {
            if (!rsp.services) {
                return
            }

            this.setState({
                loading: false,
                services: rsp.services,
            });
        })
    };

    submit(e: any) {
        e.preventDefault();
        this.props.form.validateFields((err: any, values: any) => {
            if (err) {
                return
            }

            console.log(values)

            var req: Sdk.PostServiceAccountsRequest = {
                serviceFilter: {
                    id: {
                        value: {
                            in: values.serviceIds.map((id: Sdk.ServiceId) => id.value),
                        }
                    },
                },
                select: {
                    meta: {
                        selectAll: true,
                        errors: {
                            kind: true,
                            message: {
                                formatting: true,
                                value: true,
                            },
                        },
                    },
                },
                serviceAccounts: [
                    {
                        password: {
                            value: values.password,
                        },
                        url: {
                            value: values.url,
                        },
                        handle: values.handle,
                    },
                ]
            }

            this.props.client.PostServiceAccounts(req).then((rsp) => console.log(rsp))
        });
    };

    render() {
        const {getFieldDecorator} = this.props.form;

        return (
            <div className={"register-form"}>
                <AntForm onSubmit={this.submit}>
                    <AntForm.Item label="Services">
                        {getFieldDecorator('serviceIds', {})(
                            <Checkbox.Group style={{width: '100%'}}>
                                <Row>
                                    {this.state.services.map((service: Sdk.Service) => {
                                        return <Checkbox
                                                value={service.id}>{service.id ? service.id.serviceName : ""}/{service.id ? service.id.value : ""}</Checkbox>
                                    })}
                                </Row>
                            </Checkbox.Group>,
                        )}
                    </AntForm.Item>
                    <AntForm.Item>
                        {getFieldDecorator('bla.blub', {
                            rules: [{required: true}],
                            initialValue: "",
                        })(
                            <Input
                                placeholder="bla"
                            />,
                        )}
                    </AntForm.Item>
                    <AntForm.Item>
                        {getFieldDecorator('url', {
                            rules: [{required: true}],
                            initialValue: "",
                        })(
                            <Input
                                placeholder="Url"
                            />,
                        )}
                    </AntForm.Item>
                    <AntForm.Item>
                        {getFieldDecorator('handle', {
                            rules: [{required: true}],
                            initialValue: "",
                        })(
                            <Input
                                placeholder="Handle"
                            />,
                        )}
                    </AntForm.Item>
                    <AntForm.Item>
                        {getFieldDecorator('password', {
                            rules: [{required: true}],
                            initialValue: "",
                        })(
                            <Input
                                type="password"
                                placeholder="Password"
                            />,
                        )}
                    </AntForm.Item>
                    <AntForm.Item>
                        <Button type="primary" htmlType="submit" block>
                            Post ServiceAccount
                        </Button>
                    </AntForm.Item>
                </AntForm>
            </div>
        );
    }
}


export default AntForm.create<SerivceAccountsFormProps>({name: 'login'})(Form);