package capsule

const capsuleVersion = "v1"

type Capsule struct {
	Version string
}

type Model struct {
	SrcPath   string
	Capsule   *Capsule
	Root      *Node
	Structure map[string]interface{}
}

