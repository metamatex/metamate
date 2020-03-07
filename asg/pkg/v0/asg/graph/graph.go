//go:generate go run gen/edges.go
//go:generate go run gen/nodemap.go
//go:generate go run gen/nodeslice.go

package graph

import (
	"context"
	"gopkg.in/yaml.v2"
	"regexp"
)

func Print(i interface{}, exclude ...string) {
	println(Sprint(i, exclude...))
}

func Sprint(i interface{}, exclude ...string) (string) {
	b, err := yaml.Marshal(i)
	if err != nil {
		panic(err)
	}

	re := regexp.MustCompile("(?m)[\r\n]+^.*xxx_unrecognized.*$")
	res := re.ReplaceAll(b, []byte{})

	re = regexp.MustCompile("(?m)[\r\n]+^.*: null.*$")
	res = re.ReplaceAll(res, []byte{})

	re = regexp.MustCompile("(?m)[\r\n]+^.*: {}.*$")
	res = re.ReplaceAll(res, []byte{})

	for _, e := range exclude {
		re = regexp.MustCompile("(?m)[\r\n]+^.*" + e + ".*$")
		res = re.ReplaceAll(res, []byte{})
	}

	return string(res)
}

func GetTypeDependenciesFromTypeIds(rn *RootNode, tnm TypeNodeMap) (TypeNodeMap) {
	ctx := context.Background()

	seenTypes, _ := ctx.Value("seenTypes").(map[NodeId]bool)
	if seenTypes == nil {
		seenTypes = map[NodeId]bool{}
		ctx = context.WithValue(ctx, "seenTypes", seenTypes)
	}

	ids0 := NodeIds{}
	tnm.Each(func(tn *TypeNode) {
		ids0 = append(ids0, tn.Edges.Types.Resolver.dependencies(ctx)...)
	})

	return rn.Types.ByIds(ids0...)
}

func GetEnumDependenciesFromTypeIds(rn *RootNode, tnm TypeNodeMap) (EnumNodeMap) {
	ctx := context.Background()

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

	ids0 := NodeIds{}
	tnm.Each(func(tn *TypeNode) {
		ids0 = append(ids0, tn.Edges.Enums.Resolver.dependencies(ctx)...)
	})

	return rn.Enums.ByIds(ids0...)
}

func GetTypeDependenciesFromEndpointIds(rn *RootNode, enm EndpointNodeMap) (TypeNodeMap) {
	ctx := context.Background()

	seenTypes, _ := ctx.Value("seenTypes").(map[NodeId]bool)
	if seenTypes == nil {
		seenTypes = map[NodeId]bool{}
		ctx = context.WithValue(ctx, "seenTypes", seenTypes)
	}

	ids0 := NodeIds{}
	enm.Each(func(en *EndpointNode) {
		ids0 = append(ids0, en.Edges.Types.Resolver.dependencies(ctx)...)
	})

	return rn.Types.ByIds(ids0...)
}

func GetEnumDependenciesFromEndpointIds(rn *RootNode, enm EndpointNodeMap) (EnumNodeMap) {
	ctx := context.Background()

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

	ids0 := NodeIds{}
	enm.Each(func(en *EndpointNode) {
		ids0 = append(ids0, en.Edges.Enums.Resolver.dependencies(ctx)...)
	})

	return rn.Enums.ByIds(ids0...)
}

