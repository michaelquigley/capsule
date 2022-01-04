package capsule

type StructuralDirective interface {
	Build(rootPath string, node *Node) (interface{}, error)
}

type TimelineStructure struct {
	Nodes []*Node
}

type TimelineStructuralDirective struct{}

func (self *TimelineStructuralDirective) Build(rootPath string, node *Node) (interface{}, error) {
	return nil, nil
}

func (self *TimelineStructuralDirective) inventory(rootPath string, node *Node) []string {
	return nil
}
