import React from 'react';
import {Col, Row} from 'antd';
import Client from '../Client';
import {Table as ClientAccountsTable} from '../components/types/ClientAccounts';
import {Table as ServiceAccountsTable, Form as ServiceAccountsForm} from '../components/types/ServiceAccounts';
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
];

interface Props {
    client: Client
}

interface State {
    serviceAccounts: Sdk.ServiceAccount[]
}

class Accounts extends React.Component<Props, State> {
    constructor(props: Props) {
        super(props);

        // this.props.client.GetClientAccounts({
        //     select: {
        //         clientAccounts: {
        //             alternativeIds: {
        //                 kind: true,
        //                 email: {
        //                     value: true,
        //                 }
        //             }
        //         }
        //     }
        // }).then((rsp: Sdk.GetClientAccountsResponse) => {
        //     if (!rsp.clientAccounts) {
        //         return
        //     }
        //
        //     this.setState({
        //         clientAccounts: rsp.clientAccounts,
        //     });
        // })
    }

    render() {

        // if (!this.state || !this.state.clientAccounts || !this.state.serviceAccounts) {
        //     return <h2>Accounts</h2>
        // }

        return (
            <div>
                <Row>
                    <h2>ClientAccounts</h2>
                    <ClientAccountsTable client={this.props.client}/>
                    <h2>ServiceAccounts</h2>
                </Row>
                <Row gutter={32}>
                    <Col span={6}>
                        <ServiceAccountsForm client={this.props.client}/>
                    </Col>
                    <Col span={18}>
                        <ServiceAccountsTable client={this.props.client}/>
                    </Col>
                </Row>
            </div>
        );
    }

}

export default Accounts;