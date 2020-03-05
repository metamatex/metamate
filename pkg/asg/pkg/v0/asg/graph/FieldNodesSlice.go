package graph


func (ns FieldNodeSlice) init(rn *RootNode, parentId NodeId) {
	for _, n := range ns {
		n.Init(rn, parentId)
	}
}