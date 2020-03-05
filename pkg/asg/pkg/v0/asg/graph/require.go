package graph

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
)

//func requireHasUniqueNames(n Node, es ...multiedge.Edge) (err error) {
//	for _, e := range es {
//		names := map[string]bool{}
//
//		for _, n0 := range n.GetNodes(e) {
//			_, ok := names[n0.Name()]
//			if ok {
//				err = errors.New(fmt.Sprintf("%v %v %v %v is not unique", reflect.TypeOf(n).Name(), e, n.Name(), n0.Name()))
//
//				return
//			}
//
//			names[n0.Name()] = true
//		}
//	}
//
//	return
//}

func requireNameIsUpperCamelCase(n Node) (err error) {
	r := regexp.MustCompile(`^[A-Z][a-zA-Z0-9]*$`)

	m := r.FindString(n.Name())

	if m != n.Name() {
		err = errors.New(fmt.Sprintf("%s Name %v must be upper camelcase", reflect.TypeOf(n).Name(), n.Name()))

		return
	}

	return
}

func requireHasValues(n *EnumNode) (err error) {
	if len(n.Data.Values) == 0 {
		err = errors.New(fmt.Sprintf("enum %v has no values", n.Name()))
	}

	return
}

func requireHasNameAndLowerCaseId(n Node) (errs []error) {
	err := requireHasName(n)
	if err != nil {
		errs = append(errs, err)
	}

	err = requireHasLowerCaseId(n)
	if err != nil {
		errs = append(errs, err)
	}

	return
}

func requireHasName(n Node) (err error) {
	if n.Name() == "" {
		err = errors.New(fmt.Sprintf("%v %v has no Name", reflect.TypeOf(n).Name(), n.Name()))

		return
	}

	return
}

//func requireNodes(n Node, es ...multiedge.Edge) (errs []error) {
//	for _, r := range es {
//		if n.GetNodes(r) == nil {
//			errs = append(errs, errors.New(fmt.Sprintf("%v %v has no %v nodes", reflect.TypeOf(n).Name(), n.Name(), r)))
//
//			return
//		}
//	}
//
//	return
//}
//
//func requireNode(n Node, es ...edge.Edge) (errs []error) {
//	for _, r := range es {
//		if n.GetNode(r) == nil {
//			errs = append(errs, errors.New(fmt.Sprintf("%v %v has no %v node", reflect.TypeOf(n).Name(), n.Name(), r)))
//
//			return
//		}
//	}
//
//	return
//}

func requireUniqueValues(n *EnumNode) (errs []error) {
	valueCount := map[string]int{}

	for _, v := range n.Data.Values {
		valueCount[v]++
	}

	for k, v := range valueCount {
		if v != 1 {
			errs = append(errs, errors.New(fmt.Sprintf("%v %v has non unique value %v", reflect.TypeOf(n).Name(), n.Name(), k)))
		}
	}

	return
}

func requireHasLowerCaseId(n Node) (err error) {
	if n.Id() == "" {
		err = errors.New(fmt.Sprintf("%v %v has no Id", reflect.TypeOf(n).Elem().Name(), n.Name()))

		return
	}

	r := regexp.MustCompile(`^(?:[_a-z0-9]+)+$`)
	m := r.FindString(string(n.Id()))

	if m != string(n.Id()) {
		err = errors.New(fmt.Sprintf("%v %v Id %v must be lowercase", reflect.TypeOf(n).Name(), n.Name(), n.Id()))

		return
	}

	return
}

func requireNameIsLowerCamelCase(n *FieldNode) (err error) {
	r := regexp.MustCompile(`^(?:[a-z][a-zA-Z0-9]+)+$`)

	m := r.FindString(n.Name())

	if m != n.Name() {
		err = errors.New(fmt.Sprintf("fieldname %v must be lower camelcase", n.Name()))

		return
	}

	return
}

func requireCrossReference(n *RelationNode) (err error) {
	activePn := n.Edges.Path.Active()
	passivePn := n.Edges.Path.Passive()

	if activePn.Edges.Type.Resolver.From() != passivePn.Edges.Type.Resolver.To() || activePn.Edges.Type.Resolver.To() != passivePn.Edges.Type.Resolver.From() {
		err = errors.New(fmt.Sprintf("node %v paths need to reference each other", n.Name()))
	}

	return
}

func requirePaths(n *RelationNode) (err error) {
	if n.Edges.Path.Active() == nil || n.Edges.Path.Passive() == nil {
		err = errors.New(fmt.Sprintf("node %v both paths need to be set", n.Name()))
	}

	return
}
