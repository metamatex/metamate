package graph

import (
	"context"
)

func init() {
	TypeToTypesIdResolver_Misses = func(ctx context.Context, r TypeToTypesIdResolver) (ids []NodeId) {
		r.n.Edges.Fields.Holds().Each(func(fn *FieldNode) {
			id := fn.Edges.Type.Resolver.Holds()

			if id == "" {
				return
			}

			if !r.n.root.Types.HasId(id) {
				ids = append(ids, id)
			}
		})

		return
	}

	TypeToTypesIdResolver_Dependencies = func(ctx context.Context, r TypeToTypesIdResolver) (ids []NodeId) {
		seenTypes, _ := ctx.Value("seenTypes").(map[NodeId]bool)
		if seenTypes == nil {
			seenTypes = map[NodeId]bool{}
		}

		r.n.Edges.Fields.Holds().Each(func(fn *FieldNode) {
			holdsId := fn.Edges.Type.Resolver.Holds()
			if holdsId == "" {
				return
			}

			if !r.n.root.Types.HasId(holdsId) {
				return
			}

			holds := fn.Edges.Type.Holds()
			if holds == nil {
				return
			}

			_, ok := seenTypes[holds.Id()]
			if ok {
				return
			}

			ids = append(ids, holds.Id())

			seenTypes[holds.Id()] = true

			ctx = context.WithValue(ctx, "seenTypes", seenTypes)

			ids = append(ids, holds.Edges.Types.Resolver.dependencies(ctx)...)
		})

		return
	}

	TypeToEnumsIdResolver_Dependencies = func(ctx context.Context, r TypeToEnumsIdResolver) (ids []NodeId) {
		seenTypes, _ := ctx.Value("seenTypes").(map[NodeId]bool)
		if seenTypes == nil {
			seenTypes = map[NodeId]bool{}
			ctx = context.WithValue(ctx, "seenTypes", seenTypes)
		}

		seenEnums, _ := ctx.Value("seenEnums").(map[NodeId]bool)
		if seenEnums == nil {
			seenEnums = map[NodeId]bool{}
			ctx = context.WithValue(ctx, "seenEnums", seenEnums)
		}

		r.n.Edges.Fields.Holds().Each(func(fn *FieldNode) {
			holdsType := fn.Edges.Type.Holds()
			if holdsType != nil {
				_, ok := seenTypes[holdsType.Id()]
				if ok {
					return
				}
				seenTypes[holdsType.Id()] = true

				ids = append(ids, holdsType.Edges.Enums.Resolver.dependencies(ctx)...)
			}

			holdsEnum := fn.Edges.Enum.Holds()
			if holdsEnum != nil {
				_, ok := seenEnums[holdsEnum.Id()]
				if ok {
					return
				}
				seenEnums[holdsEnum.Id()] = true

				ids = append(ids, holdsEnum.Id())
			}
		})

		return
	}

	EndpointToEnumsIdResolver_Dependencies = func(ctx context.Context, r EndpointToEnumsIdResolver) (ids []NodeId) {
		seenTypes, _ := ctx.Value("seenTypes").(map[NodeId]bool)
		if seenTypes == nil {
			seenTypes = map[NodeId]bool{}
			ctx = context.WithValue(ctx, "seenTypes", seenTypes)
		}

		seenEnums, _ := ctx.Value("seenEnums").(map[NodeId]bool)
		if seenEnums == nil {
			seenEnums = map[NodeId]bool{}
			ctx = context.WithValue(ctx, "seenEnums", seenEnums)
		}

		responseIds := r.n.Edges.Type.Response().Edges.Enums.Resolver.dependencies(ctx)
		requestIds := r.n.Edges.Type.Request().Edges.Enums.Resolver.dependencies(ctx)

		return MergeNodeIds(responseIds, requestIds)
	}

	EndpointToTypesIdResolver_Dependencies = func(ctx context.Context, r EndpointToTypesIdResolver) (ids []NodeId) {
		seenTypes, _ := ctx.Value("seenTypes").(map[NodeId]bool)
		if seenTypes == nil {
			seenTypes = map[NodeId]bool{}
			ctx = context.WithValue(ctx, "seenTypes", seenTypes)
		}

		request := r.n.Edges.Type.Request()
		response := r.n.Edges.Type.Response()

		seenTypes[request.Id()] = true
		seenTypes[response.Id()] = true

		requestIds := request.Edges.Types.Resolver.dependencies(ctx)
		responseIds := response.Edges.Types.Resolver.dependencies(ctx)

		return MergeNodeIds([]NodeId{request.Id(), response.Id()}, requestIds, responseIds)
	}
}























