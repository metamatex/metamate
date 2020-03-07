import React from 'react';
import {Col, Icon, Layout, Menu, Row} from 'antd';
import {Link} from "react-router-dom";

const {SubMenu} = Menu;

const {Header, Sider, Content} = Layout;

function handleClick(e: any) {
    console.log('click', e);
}

interface Props {
}

class App extends React.Component<Props> {
    constructor(props: Props) {
        super(props)
    }

    state = {
        collapsed: false,
    };

    toggle = () => {
        this.setState({
            collapsed: !this.state.collapsed,
        });
    };

    render() {
        return (
            <Layout className={"app2"}>
                <Sider theme={"light"} trigger={null} collapsible collapsed={this.state.collapsed}
                       className={"app2__left"}>
                    <Icon
                        className="trigger"
                        type={this.state.collapsed ? 'menu-unfold' : 'menu-fold'}
                        onClick={this.toggle}
                    />
                    <Menu mode="inline" defaultSelectedKeys={['1']}>
                        <Menu.Item key="1">
                            <Icon type="user"/>
                            <Link to={"/discovery"}>Discovery</Link>
                        </Menu.Item>
                        <Menu.Item key="2">
                            <Icon type="video-camera"/>
                            <Link to={"/accounts"}>Accounts</Link>
                        </Menu.Item>
                        <Menu.Item key="3">
                            <Icon type="upload"/>
                            <Link to={"/settings"}>Settings</Link>
                        </Menu.Item>
                    </Menu>
                </Sider>
                <Layout>
                    <Header style={{background: '#fff', padding: 0}}>
                    </Header>
                    <Content className={"app2__content"}>
                        <Row>
                            <Col span={24}>
                                {this.props.children}
                            </Col>
                        </Row>
                    </Content>
                </Layout>
            </Layout>
        )
    }
}

export default App;