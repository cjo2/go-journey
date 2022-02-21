package algorithms

import (
	data_structures "github.com/cjo2/go-journey/data-structures"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRearrange(t *testing.T) {
	l := data_structures.LinkedList{}
	l.Append(&data_structures.Node{Value: 1})
	l.Append(&data_structures.Node{Value: 2})
	l.Append(&data_structures.Node{Value: 3})
	l.Append(&data_structures.Node{Value: 4})
	l.Append(&data_structures.Node{Value: 5})

	// Odd number length
	node, _ := l.GetNth(0)
	assert.Equal(t, 1, node.Value)
	node, _ = l.GetNth(1)
	assert.Equal(t, 2, node.Value)
	node, _ = l.GetNth(2)
	assert.Equal(t, 3, node.Value)
	node, _ = l.GetNth(3)
	assert.Equal(t, 4, node.Value)
	node, _ = l.GetNth(4)
	assert.Equal(t, 5, node.Value)

	Rearrange(l)

	node, _ = l.GetNth(0)
	assert.Equal(t, 1, node.Value)
	node, _ = l.GetNth(1)
	assert.Equal(t, 5, node.Value)
	node, _ = l.GetNth(2)
	assert.Equal(t, 2, node.Value)
	node, _ = l.GetNth(3)
	assert.Equal(t, 4, node.Value)
	node, _ = l.GetNth(4)
	assert.Equal(t, 3, node.Value)

	// Even number length
	l = data_structures.LinkedList{}
	l.Append(&data_structures.Node{Value: 1})
	l.Append(&data_structures.Node{Value: 2})
	l.Append(&data_structures.Node{Value: 3})
	l.Append(&data_structures.Node{Value: 4})

	node, _ = l.GetNth(0)
	assert.Equal(t, 1, node.Value)
	node, _ = l.GetNth(1)
	assert.Equal(t, 2, node.Value)
	node, _ = l.GetNth(2)
	assert.Equal(t, 3, node.Value)
	node, _ = l.GetNth(3)
	assert.Equal(t, 4, node.Value)

	Rearrange(l)

	node, _ = l.GetNth(0)
	assert.Equal(t, 1, node.Value)
	node, _ = l.GetNth(1)
	assert.Equal(t, 4, node.Value)
	node, _ = l.GetNth(2)
	assert.Equal(t, 2, node.Value)
	node, _ = l.GetNth(3)
	assert.Equal(t, 3, node.Value)

}
