package capsule

import "fmt"

// Dump the logical contents of a Model for debugging and inspection.
//
func Dump(model *Model) string {
	if model == nil {
		return ""
	}
	out := fmt.Sprintf("Model:\n")
	out += fmt.Sprintf("\tCapsule\n\t\tVersion: '%v'\n", model.Capsule.Version)
	return out
}
