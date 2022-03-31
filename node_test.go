package capsule

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Node_FeaturesWith(t *testing.T) {
	n := &Node{}
	n.Features = append(n.Features, &Feature{Name: "1", Attributes: Attributes{"a": "A", "b": "B"}})
	n.Features = append(n.Features, &Feature{Name: "2", Attributes: Attributes{"1": "10", "2": "20", "3": 30}})
	n.Features = append(n.Features, &Feature{Name: "3", Attributes: Attributes{"a": "A", "4": 40}})

	ftrs := n.FeaturesWith(Attributes{"a": "A"})
	assert.Equal(t, 2, len(ftrs))
	assert.Equal(t, "1", ftrs[0].Name)
	assert.Equal(t, "3", ftrs[1].Name)

	ftrs = n.FeaturesWith(Attributes{"not": "found"})
	assert.Equal(t, 0, len(ftrs))

	ftrs = n.FeaturesWith(Attributes{"3": 30})
	assert.Equal(t, 1, len(ftrs))
	assert.Equal(t, "2", ftrs[0].Name)
}

func Test_Node_FeaturesWithout(t *testing.T) {
	n := &Node{}
	n.Features = append(n.Features, &Feature{Name: "1", Attributes: Attributes{"a": "A", "b": "B"}})
	n.Features = append(n.Features, &Feature{Name: "2", Attributes: Attributes{"1": "10", "2": "20", "3": 30}})
	n.Features = append(n.Features, &Feature{Name: "3", Attributes: Attributes{"a": "A", "4": 40}})

	ftrs := n.FeaturesWithout(Attributes{"a": "A"})
	assert.Equal(t, 1, len(ftrs))
	assert.Equal(t, "2", ftrs[0].Name)
}
