// generated by metactl sdk gen 
package mql

import (
    "context"
)

type KubernetesService interface {
	Name() (string)
	GetGetServicesEndpoint() (GetServicesEndpoint)
    GetServices(ctx context.Context, req GetServicesRequest) (rsp GetServicesResponse)
}