import React from 'react';
import {Button, Table as AntdTable} from 'antd';
import YAML from 'yaml'
import SyntaxHighlighter from 'react-syntax-highlighter';
import {atomOneLight} from 'react-syntax-highlighter/dist/esm/styles/hljs';
import Sdk from '../../../Sdk_';
import Client from '../../../Client';

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
];

interface Props {
    client: Client
}

interface State {
    loading: boolean
    clientAccounts: Sdk.ClientAccount[],
}

class Table extends React.Component<Props, State> {
    constructor(props: Props) {
        super(props);

        this.load = this.load.bind(this)

        this.state = {
            loading: false,
            clientAccounts: [],
        }

        this.load()
    }

    render() {
        return (
            <div>
                <Button type="primary" onClick={this.load} loading={this.state.loading}>
                    Reload
                </Button>
                <AntdTable size={"small"} columns={columns} pagination={false} dataSource={this.state.clientAccounts}
                       expandedRowRender={(account: Sdk.ClientAccount) => <SyntaxHighlighter language="yaml"
                                                                                   style={atomOneLight}>{YAML.stringify(account)}</SyntaxHighlighter>}/>
            </div>
        );
    }

    load() {
        this.setState({loading: true});

        this.props.client.GetClientAccounts({}).then((rsp: Sdk.GetClientAccountsResponse) => {
            if (!rsp.clientAccounts) {
                return
            }

            this.setState({
                loading: false,
                clientAccounts: rsp.clientAccounts,
            });
        })
    };

}

export default Table;