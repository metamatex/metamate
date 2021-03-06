# MetaMate ![test](https://github.com/metamatex/metamate/workflows/test/badge.svg) ![build](https://github.com/metamatex/metamate/workflows/build/badge.svg)

MetaMate is an open-source semantic service bus and provides you an api for everything.

This monorepo hosts
- `asg` abstract schema graph - an abstraction of MetaMate's data and communication layer
- `gen` generated sdks
- `generic` generic representation of types, powers everything that needs to handle a lot of different entities
- `hackernews-svc` hackernews gateway service
- `kubernetes-svc` discovery service for services running on kubernetes
- `mastodon-svc` mastodon gateway service
- `metactl` MetaMate's cli
- `metamate` semantic service bus

## Showcase

Please see [showcase.metamate.io](https://showcase.metamate.io) for examples applications built with MetaMate

## Installation

#### metamate

osx `brew install metamatex/taps/metamate`

docker `docker run metamatex/metamate:latest`

For all other platforms, please see [releases](https://github.com/metamatex/metamate/releases)

#### metactl

`metactl` is MetaMate's cli. It generates sdks and helps you explore the asg.

osx `brew install metamatex/taps/metactl`

For all other platforms, please see [releases](https://github.com/metamatex/metamate/releases)

## Usage

[How to build an application with MetaMate](https://metamate.io/docs/0.1/how-to-build-an-application-with-metamate/)

## Development

Building requires a couple dependencies

`yarn` is required to build the graphql explorer [Installation](https://classic.yarnpkg.com/en/docs/install/#mac-stable)

`esc` is required to add static assets to the `metamate` binary [Installation](https://github.com/mjibson/esc)

```sh
# yarn
curl -o- -L https://yarnpkg.com/install.sh | bash

# esc
go get -u github.com/mjibson/esc
```

Run `make prepare`

To build `metamate` and `metactl` run `make build`

To test `metamate` and `metactl` run `make test`

## Roadmap

MetaMate aims to provide an abstraction layer for all network connected datastores, which can be databases, websites, apis etc. The challenge here is to derive an api and internal concepts that cover all major use-cases. MetaMate needs to be able to handle different kinds of pagination, entity representations, authentication methods etc.

#### v0.x

The community behind MetaMate has a pretty good understanding of the problem domain by now. The focus on the road to `v1.x` lies on following:

- **streaming** clients want to be able to subscribe to changes in a result set, a service may provide a streaming interface or emulates it MetaMate by polling
- **tracing** developers want to monitor how requests propagate through the stack
- **distribute the ASG via MetaMate** the asg is currently hard-coded, which requires a recompilation of the stack to deliver new types
- **introducing result sets** currently results and paginations are merged into one result set
- **split communication types** currently the types for client to MetaMate communication and MetaMate to service communication are the same which might cause some confusion
- **specification** provide a formal specification for MetaMate's overall api as code

Breaking changes will occur whenever we remove or rename an identifier or change the type or content or a returned value. We already know of a handful of minor breaking changes and aren't able to provide a full stability guarantee yet.

#### v1.x

The community intends to provide a stability guarantee from `v1.x` onwards