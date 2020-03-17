# MetaMate

MetaMate is an open-source semantic service bus and provides you an api for everything.

This monorepo hosts
- `asg` abstract schema graph - an abstraction of MetaMate's data and communication layer
- `auth-svc` authenticates `ClientAccounts`, hashes password and verifies `Tokens`
- `gen` generated sdks
- `generic` generic representation of types, powers everything that needs to handle a lot of different entities
- `kubernetes-svc` discovery service for services running on kubernetes
- `mastodon-svc` gateway service for mastodon
- `metactl` MetaMate's cli
- `metamate` semantic service bus
- `sqlx-svc` storage service that maps mql to sql

## Installation

#### metamate

To spin up a metamate instance simply run `docker run metamatex/metamate:latest`. It's preconfigured with some sane defaults.

#### metactl

`metactl` is MetaMate's cli. It generates sdks and helps you explore the asg.

macOs `brew install metamatex/taps/metactl`

For all other platforms, please see our [releases](https://github.com/metamatex/metamate/releases)

## Development

MetaMate develops in multiple stages

MetaMate tries to provide an abstraction layer for all network connected datastores, which can be databases, websites, apis etc. The challenge here is to derive an api that covers all major use-cases. MetaMate needs to be able to handle different kinds of paginations, entity respresentations and authentication methods.

#### v0.x

The community behind MetaMate has a pretty good understanding of the problem domain by now.

- **credential management** services need credentials to obtain tokens and tokens to make authorizated requests, services will indicate the necessitate for a token
- **streaming** clients want to be able to subscribe to changes in a result set, a service may provide a streaming interface or emulates it MetaMate by polling
- **tracing** developers want to monitor how requests propagate through the stack
- **exposing the ASG via MetaMate** the asg is currently hard-coded, which requires a recompilation of the stack to deliver new types.
- **RBAC** administrator want to be able to set rules of who can access what and what subset
- **introducing result sets** currently results and paginations are merged into one result set
- **split communication types** currently the types for client to MetaMate communication and MetaMate to service communication are the same which might cause some confusion
- **specification** provide a formal specification for MetaMate's overall api as code

Doesn't provide a stability guarantee yet

#### v1.x

Provides a stability guarantee