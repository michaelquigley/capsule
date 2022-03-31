package capsule

type Feature struct {
	Name       string
	Attributes Attributes
	Object     interface{}
}

type Features []*Feature

func (ftrs Features) With(attrs Attributes) Features {
	return ftrs.matchAttributes(true, attrs)
}

func (ftrs Features) Without(attrs Attributes) Features {
	return ftrs.matchAttributes(false, attrs)
}

func (ftrs Features) Named(name string) *Feature {
	matches := ftrs.NameIn([]string{name})
	if len(matches) == 1 {
		return matches[0]
	}
	return nil
}

func (ftrs Features) NameIn(names []string) Features {
	var matches Features
	for _, ftr := range ftrs {
		for _, name := range names {
			if name == ftr.Name {
				matches = append(matches, ftr)
				break
			}
		}
	}
	return matches
}

func (ftrs Features) NameNotIn(names []string) Features {
	var matches Features
	for _, ftr := range ftrs {
		match := false
		for _, name := range names {
			if name == ftr.Name {
				match = true
				break
			}
		}
		if !match {
			matches = append(matches, ftr)
		}
	}
	return matches
}

func (ftrs Features) matchAttributes(include bool, attrs Attributes) Features {
	var matches Features
	for _, ftr := range ftrs {
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
		if match == include {
			matches = append(matches, ftr)
		}
	}
	return matches
}
