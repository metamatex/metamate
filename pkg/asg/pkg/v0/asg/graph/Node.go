package graph

type Node interface {
	Type() string
	SetId(NodeId)
	Id() NodeId
	SetRoot(*RootNode)
	Name() string
	PluralName() string
	FieldName() string
	SetName(string)
	Flags() FlagsContainer
	Validate() []error
	GetEdges() interface{}
	GetData() interface{}
	Wire()
	Print()
}