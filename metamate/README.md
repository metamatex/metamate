# metamate


### communication

`pkg/v0/communication/clients`

the `clients` package contains all clients a MetaMate uses to communicate with services

`httpjson` fetches json serialized data over http

`pkg/v0/communication/servers`

the `servers` package contains all the tcp endpoints a MetaMate exposes

visit [metamate.one](http://metamate.one/) to get a live overview of all the exposed endpoints

`admin` exposes MetaMate's administration interface

`config` exposes the loaded configuration

`debug` exposes golang runtime information

`explorer` exposes MetaMate's graphql explorer

`graphql` exposes MetaMate's graphql interface

`httpjson` exposes MetaMate's json over http api

`index` exposes MetaMate's index page