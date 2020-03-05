import React from 'react';
import {Divider, Table, Tag} from 'antd';
import YAML from 'yaml'
import SyntaxHighlighter from 'react-syntax-highlighter';
import {atomOneLight} from 'react-syntax-highlighter/dist/esm/styles/hljs';
import Client from '../Client';
import Sdk from './../Sdk_';

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
        title: 'transport',
        dataIndex: 'transport',
        key: 'transport',

    },
    {
        title: 'isVirtual',
        dataIndex: 'isVirtual',
        key: 'isVirtual',
        render: (text: boolean) => text ? "true" : "false",
    },
    {
        title: 'port',
        dataIndex: 'port',
        key: 'port',

    },
];

interface Props {
    client: Client
}

interface State {
    svcs: Sdk.Service[]
}

class Discovery extends React.Component<Props, State> {
    constructor(props: Props) {
        super(props);

        let rsp = this.props.client.GetServices({}).then((rsp: Sdk.GetServicesResponse) => {
            if (!rsp.services) {
                return
            }

            this.setState({
                svcs: rsp.services,
            });
        })
    }

    render() {
        if (!this.state || !this.state.svcs) {
            return <div><h2>Discovery</h2></div>
        }

        return (
            <div>
                <h2>Discovery</h2>
                <Table size={"small"} columns={columns} pagination={false} dataSource={this.state.svcs} expandedRowRender={(svc: Sdk.Service) => <SyntaxHighlighter language="javascript" style={atomOneLight}>{YAML.stringify(svc)}</SyntaxHighlighter>}/>
            </div>
        );
    }

}

export default Discovery;