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

const CapsuleFeature = "capsule.yaml"
const capsuleVersion = "v0.1"
