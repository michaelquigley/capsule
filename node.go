package capsule

import "path/filepath"

type Node struct {
	Path     string
	Features []*Feature
	Parent   *Node
	Children []*Node
}

func (n *Node) FullPath() string {
	if n.Parent != nil {
		return filepath.ToSlash(filepath.Join(n.Parent.FullPath(), n.Path))
	} else {
		return filepath.ToSlash(n.Path)
	}
}

func (n *Node) FeatureNamed(name string) *Feature {
	for _, ftr := range n.Features {
		if ftr.Name == name {
			return ftr
		}
	}
	return nil
}

func (n *Node) FeaturesWith(attrs Attributes) []*Feature {
	var matches []*Feature
	for _, ftr := range n.Features {
		match := true
		for k, v := range attrs {
			if ftrV, found := ftr.Attributes[k]; found {
				if ftrV != v {
					match = false
					break
				}
			} else {
				match = false
				break
			}
		}
		if match {
			matches = append(matches, ftr)
		}
	}
	return matches
}

func (n *Node) FeaturesWithout(attrs Attributes) []*Feature {
	var matches []*Feature
	for _, ftr := range n.Features {
		match := true
		for k, v := range attrs {
			if ftrV, found := ftr.Attributes[k]; found {
				if ftrV != v {
					match = false
					break
				}
			} else {
				match = false
				break
			}
		}
		if !match {
			matches = append(matches, ftr)
		}
	}
	return matches
}
