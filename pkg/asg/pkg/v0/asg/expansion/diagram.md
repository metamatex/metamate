https://mermaidjs.github.io/mermaid-live-editor

graph TD

Service(Service)
ServiceFilter(ServiceFilter)
ServiceSort(ServiceSort)
ServiceSelect(ServiceSelect)
GetServicesResponse(GetServicesResponse)
GetServicesResponseSelect(GetServicesResponseSelect)
GetServicesRequest(GetServicesRequest)
CreateServicesResponse(CreateServicesResponse)
CreateServicesResponseSelect(CreateServicesResponseSelect)
CreateServicesRequest(CreateServicesRequest)
GetServicesRequestFilter(GetServicesRequestFilter)
CreateServicesRequestFilter(CreateServicesRequestFilter)
GetServicesEndpoint(GetServicesEndpoint)
CreateServicesEndpoint(CreateServicesEndpoint)
GetServicesEndpointFilter(GetServicesEndpointFilter)
CreateServicesEndpointFilter(CreateServicesEndpointFilter)
ResponseMeta(ResponseMeta)
Endpoint(Endpoint)
EndpointFilter(EndpointFilter)
TypeMeta(TypeMeta)

Service --> |GET_RSP| GetServicesResponse
Service --> |FILTER| ServiceFilter
Service --> |SORT| ServiceSort
Service --> |DEP/CREATE_REQ| CreateServicesRequest
Service --> |CREATE_RSP| CreateServicesResponse
Service --> |GET_REQ| GetServicesRequest
Service --> |SELECT| ServiceSelect
Service --> |DEP| ResponseMeta
Service --> |DEP| TypeMeta 

ServiceFilter --> |DEP| CreateServicesRequest
ServiceFilter --> |DEP| GetServicesRequest
ServiceFilter --> |DEP| ServiceSelect

ServiceSelect --> |DEP| GetServicesResponseSelect
ServiceSort --> |DEP| ServiceSelect

CreateServicesEndpoint --> |DEP| Endpoint
CreateServicesEndpoint --> |FILTER| CreateServicesEndpointFilter
CreateServicesEndpointFilter --> |DEP| EndpointFilter
CreateServicesRequest --> |FILTER| CreateServicesRequestFilter
CreateServicesRequestFilter --> |DEP| CreateServicesEndpoint
CreateServicesResponse --> |SELECT| CreateServicesResponseSelect
CreateServicesResponseSelect --> |DEP| CreateServicesRequest
GetServicesEndpoint --> |DEP| Endpoint
GetServicesEndpoint --> |FILTER| GetServicesEndpointFilter
GetServicesEndpointFilter --> |DEP| EndpointFilter
GetServicesRequest --> |FILTER| GetServicesRequestFilter
GetServicesRequestFilter --> |DEP| GetServicesEndpoint
GetServicesResponseSelect --> |DEP| GetServicesRequest 
GetServicesResponse --> |SELECT| GetServicesResponseSelect


Endpoint --> |DEP| Service
Endpoint --> |FILTER| EndpointFilter
EndpointFilter --> |DEP| ServiceFilter

TypeMeta --> |DEP| Service

ResponseMeta --> |DEP| GetServicesResponse
ResponseMeta --> |DEP| CreateServicesResponse

