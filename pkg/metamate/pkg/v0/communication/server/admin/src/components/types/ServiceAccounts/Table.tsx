import React from 'react';
import {Button, Col, Input, Row, Table as AntdTable} from 'antd';
import YAML from 'yaml'
import Sdk from './../../../Sdk_';
import Client from './../../../Client';

const columns = [
    {
        title: 'service',
        dataIndex: 'id.serviceName',
        key: 'service',

    },
    {
        title: 'id',
        dataIndex: 'id.value',
        key: 'id',

    },
    {
        title: 'url',
        dataIndex: 'url.value',
        key: 'url',

    },
    {
        title: 'handle',
        dataIndex: 'handle',
        key: 'handle',

    },
    {
        title: 'password',
        dataIndex: 'password.value',
        key: 'password',

    },
];

interface Props {
    client: Client
}

interface State {
    loading: boolean
    serviceAccounts: Sdk.ServiceAccount[],
}

class Table extends React.Component<Props, State> {
    constructor(props: Props) {
        super(props);

        this.load = this.load.bind(this)
        this.expand = this.expand.bind(this)

        this.state = {
            loading: false,
            serviceAccounts: [],
        }

        this.load()
    }

    expand(account: Sdk.ServiceAccount) {
        return <Row>
            <Col span={12}></Col>
            <Col span={12}>
                <Input.TextArea autoSize={true}>{YAML.stringify(account)}</Input.TextArea>
                <Button className={"types__table__reload"} type="primary" onClick={this.load}
                        loading={this.state.loading}>
                    Update
                </Button>
                <Button className={"types__table__reload"} type="primary" onClick={this.load}
                        loading={this.state.loading}>
                    Delete
                </Button>
            </Col>
        </Row>
    }

    render() {
        return (
            <div className={"types__table"}>
                <AntdTable size={"small"} columns={columns} pagination={false} dataSource={this.state.serviceAccounts}
                           expandedRowRender={this.expand}/>
                <Button className={"types__table__reload"} type="primary" onClick={this.load}
                        loading={this.state.loading}>
                    Reload
                </Button>
            </div>
        );
    }

    load() {
        this.setState({loading: true});

        this.props.client.GetServiceAccounts({}).then((rsp: Sdk.GetServiceAccountsResponse) => {
            if (!rsp.serviceAccounts) {
                return
            }

            this.setState({
                loading: false,
                serviceAccounts: rsp.serviceAccounts,
            });
        })
    };

}

export default Table;