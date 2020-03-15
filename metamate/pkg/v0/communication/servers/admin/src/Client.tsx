import * as axios from 'axios';
import Sdk from './Sdk_';

class Client {
    client: axios.AxiosInstance;
    token: string;
    addr: string;
    rspMetaHandler: (meta: Sdk.ResponseMeta) => void;

    constructor(token: string, addr: string, h: (meta: Sdk.ResponseMeta) => void) {
        this.token = token;
        this.addr = addr;
        this.client = axios.default.create();
        this.rspMetaHandler = h
    }

    async GetWhatevers(req: Sdk.GetWhateversRequest): Promise<Sdk.GetWhateversResponse | undefined> {
        return undefined
    }

    async GetServices(req: Sdk.GetServicesRequest): Promise<Sdk.GetServicesResponse> {
        let rsp = await this.client.request<Sdk.GetServicesResponse>({
            url: this.addr,
            method: "post",
            data: req,
            headers: {
                "X-MetaMate-Type": "GetServicesRequest",
                "Content-Type": "application/json; charset=utf-8",
            },
        })

        if ((rsp) && (rsp.data) && (rsp.data.meta)) {
            this.rspMetaHandler(rsp.data.meta)
        }

        return rsp.data
    }

    async GetServiceAccounts(req: Sdk.GetServiceAccountsRequest): Promise<Sdk.GetServiceAccountsResponse> {
        let rsp = await this.client.request<Sdk.GetServiceAccountsResponse>({
            url: this.addr,
            method: "post",
            data: req,
            headers: {
                "X-MetaMate-Type": "GetServiceAccountsRequest",
                "Content-Type": "application/json; charset=utf-8",
            },
        })

        if ((rsp) && (rsp.data) && (rsp.data.meta)) {
            this.rspMetaHandler(rsp.data.meta)
        }

        return rsp.data
    }

    async PostServiceAccounts(req: Sdk.PostServiceAccountsRequest): Promise<Sdk.PostServiceAccountsResponse> {
        let rsp = await this.client.request<Sdk.PostServiceAccountsResponse>({
            url: this.addr,
            method: "post",
            data: req,
            headers: {
                "X-MetaMate-Type": "PostServiceAccountsRequest",
                "Content-Type": "application/json; charset=utf-8",
            },
        })

        if ((rsp) && (rsp.data) && (rsp.data.meta)) {
            this.rspMetaHandler(rsp.data.meta)
        }

        return rsp.data
    }

    async GetClientAccounts(req: Sdk.GetClientAccountsRequest): Promise<Sdk.GetClientAccountsResponse> {
        let rsp = await this.client.request<Sdk.GetClientAccountsResponse>({
            url: this.addr,
            method: "post",
            data: req,
            headers: {
                "X-MetaMate-Type": "GetClientAccountsRequest",
                "Content-Type": "application/json; charset=utf-8",
            },
        })

        if ((rsp) && (rsp.data) && (rsp.data.meta)) {
            this.rspMetaHandler(rsp.data.meta)
        }

        return rsp.data
    }

    async PostClientAccounts(req: Sdk.PostClientAccountsRequest): Promise<Sdk.PostClientAccountsResponse> {
        let rsp = await this.client.request<Sdk.PostClientAccountsResponse>({
            url: this.addr,
            method: "post",
            data: req,
            headers: {
                "X-MetaMate-Type": "PostClientAccountsRequest",
                "Content-Type": "application/json; charset=utf-8",
            },
        })

        if ((rsp) && (rsp.data) && (rsp.data.meta)) {
            this.rspMetaHandler(rsp.data.meta)
        }

        return rsp.data
    }
}

export default Client;