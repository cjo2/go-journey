package data_structures

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBinaryTreeNode_InsertNode(t *testing.T) {
	head := &BinaryTreeNode{Value: 2}
	zero := &BinaryTreeNode{Value: 0}
	head.InsertNode(zero)
	head.InsertNode(&BinaryTreeNode{Value: 4})
	head.InsertNode(&BinaryTreeNode{Value: -1})
	head.InsertNode(&BinaryTreeNode{Value: 1})
	head.InsertNode(&BinaryTreeNode{Value: 3})
	head.InsertNode(&BinaryTreeNode{Value: 5})

	assert.Equal(t, 0, head.Left.Value)
	assert.Equal(t, -1, head.Left.Left.Value)
	assert.Equal(t, 1, head.Left.Right.Value)
	assert.Equal(t, 4, head.Right.Value)
	assert.Equal(t, 3, head.Right.Left.Value)
	assert.Equal(t, 5, head.Right.Right.Value)

	fmt.Printf("%v %v\n", zero, head.Left)
	fmt.Printf("%v %v", *zero, *head.Left)

	err := head.InsertNode(&BinaryTreeNode{Value: 2})
	assert.Error(t, err)

	err = head.InsertNode(&BinaryTreeNode{Value: 5})
	assert.Error(t, err)
}

func TestBinaryTreeNode_GetNode(t *testing.T) {
	head := &BinaryTreeNode{Value: 2}
	head.InsertNode(&BinaryTreeNode{Value: 0})
	head.InsertNode(&BinaryTreeNode{Value: 4})
	head.InsertNode(&BinaryTreeNode{Value: -1})
	head.InsertNode(&BinaryTreeNode{Value: 1})
	head.InsertNode(&BinaryTreeNode{Value: 3})
	head.InsertNode(&BinaryTreeNode{Value: 5})

	node, err := head.GetNode(0)
	assert.Equal(t, 0, node.Value)
	assert.NoError(t, err)

	node, err = head.GetNode(4)
	assert.Equal(t, 4, node.Value)
	assert.NoError(t, err)

	node, err = head.GetNode(-1)
	assert.Equal(t, -1, node.Value)
	assert.NoError(t, err)

	node, err = head.GetNode(5)
	assert.Equal(t, 5, node.Value)
	assert.NoError(t, err)

	node, err = head.GetNode(7)
	assert.Nil(t, node)
	assert.Error(t, err)
}
