import React from 'react';
import {Col, Row} from 'antd';
import EmailLoginForm from "./Welcome/EmailLoginForm";
import TokenLoginForm from "./Welcome/TokenLoginForm";
import RegisterForm from "./Welcome/RegisterForm";
import ForgotForm from "./Welcome/ForgotForm";

const
    register = "register",
    login = "login",
    forgot = "forgot",
    email = "email",
    token = "token"

var forms: { [key: string]: any; } = {
    email: <EmailLoginForm/>,
    token: <TokenLoginForm/>,
}

interface Tip {
    ImgUrl: string
    Title: string
    Text: string
}

var tips: Tip[] = [
    {
        ImgUrl: "https://metamate-io.netlify.com/images/icons/services_service.svg",
        Title: "Service Discovery",
        Text: "Discovery services supply your MetaMate with services it can talk to. These services can be publicly routable or in a private network",
    },
    {
        ImgUrl: "https://metamate-io.netlify.com/images/icons/cloud_distributed.svg",
        Title: "Distributed tracing",
        Text: "MetaMate and all it's services are instrumented. Reasoning about the inter-service communication is a breeze",
    },
    {
        ImgUrl: "https://metamate-io.netlify.com/images/icons/schema_community.svg",
        Title: "Community driven",
        Text: "Everyone can contribute new types and fields by simply opening a pull request",
    },
    {
        ImgUrl: "https://metamate-io.netlify.com/images/icons/schema_infinitely.svg",
        Title: "Infinitely backward compatible",
        Text: "Backward compatibility of the schema is programmatically enforced. Services and applications built will work forever",
    },
]

interface Props {
    forms: {
        email: any;
        token: any;
        register: any;
    }
}

interface State {
    activeLoginForm: string
    activeSection: string
    sections: { [key: string]: any; }
    tip: Tip
}

function getRandomInt(min: number, max: number) {
    min = Math.ceil(min);
    max = Math.floor(max);
    return Math.floor(Math.random() * (max - min + 1)) + min;
}

class Welcome extends React.Component<Props, State> {
    constructor(props: Props) {
        super(props);

        this.getLogin = this.getLogin.bind(this);
        this.getRegister = this.getRegister.bind(this);
        this.getForgot = this.getForgot.bind(this);

        this.state = {
            activeLoginForm: email,
            activeSection: register,
            sections: {
                register: this.getRegister,
                login: this.getLogin,
                forgot: this.getForgot,
            },
            tip: tips[getRandomInt(0, tips.length - 1)]
        };
    }

    render() {
        return (
            <div id={"bg-image"}>
                <div id="bg-image-wrapped"></div>
                <div className={"welcome"}>
                    <Row>
                        <Col span={10} offset={7}>
                            <Row type="flex">
                                <Col span={14} className={"welcome__left panel"}>
                                    <div className="panel__inner">
                                        {/*<p className={"welcome__host"}>@metamate.one</p>*/}
                                        {this.state.sections[this.state.activeSection]()}
                                    </div>
                                </Col>
                                <Col span={10} className={"welcome__right panel"}>
                                    <div className="panel__inner">
                                        <Row>
                                            <Col span={8} offset={8}>
                                                <img src={this.state.tip.ImgUrl} alt=""/>
                                            </Col>
                                            <Col span={24}>
                                                <h3>{this.state.tip.Title}</h3>
                                                <p>{this.state.tip.Text}</p>
                                            </Col>
                                        </Row>
                                    </div>
                                </Col>
                            </Row>
                        </Col>
                    </Row>
                </div>
            </div>
        );
    }

    getLogin() {
        var form = forms[this.state.activeLoginForm]

        return <div className={"login"}>
            <Row>
                <Col span={12}>
                    <h2>Sign in</h2>
                </Col>
                <Col span={12}>
                    <p className={"welcome__host"}>@metamate.one</p>
                </Col>
            </Row>
            <div className="login__sections">
                <a href="#"
                   className={this.state.activeLoginForm == email ? "login__section login__section--active" : "login__section"}
                   onClick={(e) => this.setState({activeLoginForm: email})}>Email</a>
                <a href="#"
                   className={this.state.activeLoginForm == token ? "login__section login__section--active" : "login__section"}
                   onClick={(e) => this.setState({activeLoginForm: token})}>Token</a>
            </div>
            {form}
            <a href="#" onClick={(e) => this.setState({activeSection: register})}>Create account</a>
            <a href="#" className={"login__forgot"} onClick={(e) => this.setState({activeSection: forgot})}>Forgot password?</a>
        </div>
    }

    getRegister() {
        return <div>
            <Row>
                <Col span={12}>
                    <h2>Register</h2>
                </Col>
                <Col span={12}>
                    <p className={"welcome__host"}>@metamate.one</p>
                </Col>
            </Row>
            {this.props.forms.register}
            <a href="#" onClick={(e) => this.setState({activeSection: login})}>Back</a>
        </div>
    }

    getForgot() {
        return <div>
            <Row>
                <Col span={12}>
                    <h2>Forgot password</h2>
                </Col>
                <Col span={12}>
                    <p className={"welcome__host"}>@metamate.one</p>
                </Col>
            </Row>
            <ForgotForm/>
            <a href="#" onClick={(e) => this.setState({activeSection: login})}>Back</a>
        </div>
    }
}

export default Welcome;