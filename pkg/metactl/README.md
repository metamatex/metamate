# metactl

metactl is a tool for interacting with the MetaMate ecosystem

Use metactl to:

- Generate files by populating templates with entities from the schema
- Install new services to a MetaMate instance
- Query a MetaMate instance
- Run tests

## installation

## usage

### gen

`metactl gen` generates files by populating templates with specified entities or endpoints from the schema

It expects a config file that specifies tasks on what should be rendered with what data.

```yaml
- template: .make/gotpl/schema.gotpl
  out: "{{ .Version }}/proto2/schema_.proto"
  typeFlags:
    or: [all]
  endpointFlags:
    or: [all]
  iterate: false
```

The templates are expected to use the [golang templating language](https://golang.org/pkg/text/template).

An addition to the native template functions, `metactl` comes the following convenience functions: 

[sprig - Useful template functions for Go templates](http://masterminds.github.io/sprig/)

`metactl` template functions

```
plural
        returns the plural of a noun; apple -> apples, repository -> repositories
toLowerCamel
        transforms a string to lower camel; whateverRequest
toCamel
        transforms a string to camel; WhateverRequest
toCamel
        transforms a string to streaming snake; WHATEVER_REQUEST
```

### test

`metactl test` queries a MetaMate instance for `SpecificationDescriptions` and displays them stylized. E.g. having the [specification-svc](https://github.com/metamatex/specification-svc) deployed in your MetaMate environment returns an ouput similar to this:

```
describe endpoints - get
   when a service returns an error
    ✔ it is returned to the client
   when getting whatevers
    ✔ it gets whatevers
describe endpoints - update
   when a service returns an error
    ✔ it is returned to the client
   when updating whatever by service id with service filter
    ✔ the updated version is returned
describe endpoints - create
   when creating whatever with select
    ✔ the selected fields are returned
   when a service returns an error
    ✔ it is returned to the client
   when creating whatever
    ✔ it is present afterwards
   when creating whatever with name id and with service filter
    ✔ whatever is only created for the given service
describe endpoints - get list by id
   when a service returns an error
    ✔ it is returned to the client
   when getting a whatever list with service id and with service filter
    ✔ it returns a list of whatevers
   when getting a whatever list with select
    ✔ the selected fields are returned
...
```

### get

`metactl get` queries a MetaMate instance for the specified type. `metactl get services` could return something like this:

```
services:
- name: discovery
  host: discovery
  endpoints:
  - kind: ENDPOINT_GET_SERVICES_ENDPOINT
  - kind: ENDPOINT_STREAM_SERVICES_ENDPOINT- name: specification-svc
  host: specification-svc
  endpoints:
  - kind: ENDPOINT_GET_SPECIFICATION_DESCRIPTIONS_ENDPOINT
  - kind: ENDPOINT_CREATE_WHATEVERS_ENDPOINT
  - kind: ENDPOINT_GET_WHATEVERS_ENDPOINT
  - kind: ENDPOINT_GET_WHATEVER_BY_ID_ENDPOINT
  - kind: ENDPOINT_GET_WHATEVER_LIST_BY_ID_ENDPOINT
  - kind: ENDPOINT_UPDATE_WHATEVERS_ENDPOINT
  - kind: ENDPOINT_DELETE_WHATEVERS_ENDPOINT
  - kind: ENDPOINT_PASS_CREATE_WHATEVERS_REQUEST_ENDPOINT
```


