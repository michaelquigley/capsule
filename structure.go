package capsule

// StructuralDefinition is a strategy that is used to build structure from a Node tree.
//
type StructuralDefinition interface {
	Build(*Node, map[string]interface{}) error
}
