package capsule

type Model struct {
	SrcPath string
	Root    *Node
}

type Node struct {
	Path       string
	Capsule    *Capsule
	Properties []*Property
	Parent     *Node
	Children   []*Node
}

type Capsule struct {
	Version   string
	Structure string
}

type Property struct {
	Name string
	Type string
}

// Parse a source path into a Model.
//
func Parse(srcPath string) (*Model, error) {
	return nil, nil
}
