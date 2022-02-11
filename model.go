package capsule

type Model struct {
	Path       string
	Capsule    *Capsule
	Root       *Node
	Structures map[string]interface{}
}

type Capsule struct {
	Version string
}

type Feature struct {
	Name       string
	Attributes Attributes
	Object     interface{}
}

type Attributes map[string]interface{}

type Node struct {
	Path     string
	Features []*Feature
	Parent   *Node
	Children []*Node
}

const capsuleVersion = "v0.1"
