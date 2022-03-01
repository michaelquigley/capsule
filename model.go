package capsule

import (
	"fmt"
	"sort"
)

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

func (a Attributes) String() string {
	var keys []string
	for key, _ := range a {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	out := "{"
	i := 0
	for _, k := range keys {
		if i > 0 {
			out += " "
		}
		out += fmt.Sprintf("%v:%v", k, a[k])
		i++
	}
	out += "}"
	return out
}

const capsuleVersion = "v0.1"
