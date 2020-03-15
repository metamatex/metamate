# metamate

### pkg/v0/boot

the `boot` package starts a MetaMate instance and

- spins up the virtual cluster and deploys embedded services
- creates clients for the MetaMate to be able to communicate with services
- initializes the request-dispatch pipelines and injects them into the request handler
- exposes endpoints

### pkg/v0/business

the `business` package contains all the internal logic and the virtual cluster

MetaMate is using a request pipeline to dispatch incoming requests. Every request creates a `RequestCtx` that propagates through the pipeline

- `funcs` defines functions (and function constructors) that take care of very detailed business logic, they all work on `RequestCtx` what allows to convienently stich them together
- `line` provides a tailored pipeline framework
- `pipeline` constructes very high-level pipelines with `line` and fills with life with `funcs`, this is the heart of MetaMate
- `validation` validates Requests and Responses
- `virtual` contains MetaMate's virtual cluster and its embedded services

### pkg/v0/communication

the `communication` package handles all in- and outgoing communication

#### pkg/v0/communication/clients

the `clients` package contains all clients a MetaMate uses to communicate with services

- `httpjson` communicates json-serialized data over http

#### pkg/v0/communication/servers

the `servers` package contains all the endpoints a MetaMate exposes

visit [metamate.one](http://metamate.one/) to get a live overview of all the exposed endpoints

- `admin` exposes MetaMate's administration interface
- `config` exposes the loaded configuration
- `debug` exposes golang runtime information
- `explorer` exposes MetaMate's graphql explorer
- `graphql` exposes MetaMate's graphql interface
- `httpjson` exposes MetaMate's json over http api
- `index` exposes MetaMate's index page