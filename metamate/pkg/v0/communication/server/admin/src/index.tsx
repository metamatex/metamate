import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import * as serviceWorker from './serviceWorker';
import {message} from 'antd';
import {BrowserRouter, Route} from "react-router-dom";
import App from "./App";
import Client from "./Client";
import Welcome from "./pages/Welcome";
import RegisterForm from "./pages/Welcome/RegisterForm";
import Discovery from "./pages/Discovery";
import Accounts from "./pages/Accounts";

import 'antd/dist/antd.css';

let c = new Client("", "http://192.168.99.100:32000/httpjson", (meta => {
    if (!meta.errors) {
        return
    }

    for (let err of meta.errors) {
        message.error(err.message)
    }
}));

function getWelcome(c: Client): any {
    return <Welcome forms={{
        email: "",
        token: "",
        register: <RegisterForm client={c}/>
    }}/>
}

const Root = () => {
    return (
        <BrowserRouter basename="/admin">
            <Route exact path="/" component={(props: any) => getWelcome(c)}/>
            <Route path="/discovery" component={(props: any) => <App children={<Discovery client={c}/>}/>}/>
            <Route path="/accounts" component={(props: any) => <App children={<Accounts client={c}/>}/>}/>
            {/*<Route path="/profile/:id" render={props => {return <Profile {...props}/>}}/>*/}
        </BrowserRouter>

    )
};

ReactDOM.render(
    <Root/>,
    document.querySelector('#root')
);

// If you want your app to work offline and load faster, you can change
// unregister() to register() below. Note this comes with some pitfalls.
// Learn more about service workers: https://bit.ly/CRA-PWA
serviceWorker.unregister();
