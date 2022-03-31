package capsule

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Note_FeaturesApi(t *testing.T) {
	n := &Node{}
	n.Features = append(n.Features, &Feature{Name: "1", Attributes: Attributes{"a": "A", "b": "B"}})
	n.Features = append(n.Features, &Feature{Name: "2", Attributes: Attributes{"1": "10", "2": "20", "3": 30}})
	n.Features = append(n.Features, &Feature{Name: "3", Attributes: Attributes{"a": "A", "4": 40}})

	t.Run("with", func(t *testing.T) {
		ftrs := n.Features.With(Attributes{"a": "A"})
		assert.Equal(t, 2, len(ftrs))
		assert.Equal(t, "1", ftrs[0].Name)
		assert.Equal(t, "3", ftrs[1].Name)

		ftrs = n.Features.With(Attributes{"not": "found"})
		assert.Equal(t, 0, len(ftrs))

		ftrs = n.Features.With(Attributes{"3": 30})
		assert.Equal(t, 1, len(ftrs))
		assert.Equal(t, "2", ftrs[0].Name)
	})

	t.Run("without", func(t *testing.T) {
		ftrs := n.Features.Without(Attributes{"a": "A"})
		assert.Equal(t, 1, len(ftrs))
		assert.Equal(t, "2", ftrs[0].Name)
	})

	t.Run("with without", func(t *testing.T) {
		ftrs := n.Features.With(Attributes{"a": "A"}).Without(Attributes{"4": 40})
		assert.Equal(t, 1, len(ftrs))
		assert.Equal(t, "1", ftrs[0].Name)
	})

	t.Run("name in", func(t *testing.T) {
		ftrs := n.Features.NameIn([]string{"1", "2"})
		assert.Equal(t, 2, len(ftrs))
		assert.Equal(t, "1", ftrs[0].Name)
		assert.Equal(t, "2", ftrs[1].Name)
	})

	t.Run("name not in", func(t *testing.T) {
		ftrs := n.Features.NameNotIn([]string{"1", "2"})
		assert.Equal(t, 1, len(ftrs))
		assert.Equal(t, "3", ftrs[0].Name)
	})

	t.Run("with name in", func(t *testing.T) {
		ftrs := n.Features.With(Attributes{"a": "A"}).NameIn([]string{"3"})
		assert.Equal(t, 1, len(ftrs))
		assert.Equal(t, "3", ftrs[0].Name)
	})
}
