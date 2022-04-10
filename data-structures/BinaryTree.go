package data_structures

import "errors"

type BinaryTreeNode struct {
	Value int
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
}

func (btn *BinaryTreeNode) InsertValue(value int) error {
	err := btn.InsertNode(&BinaryTreeNode{Value: value})
	if err != nil {
		return err
	}

	return nil
}

func (btn *BinaryTreeNode) InsertNode(node *BinaryTreeNode) error {
	newVal := node.Value
	oldVal := btn.Value
	if newVal < oldVal && btn.Left != nil {
		err := btn.Left.InsertNode(node)
		if err != nil {
			return err
		}
		return nil
	} else if newVal < oldVal && btn.Left == nil {
		btn.Left = node
		return nil
	} else if newVal > oldVal && btn.Right != nil {
		err := btn.Right.InsertNode(node)
		if err != nil {
			return err
		}
		return nil
	} else if newVal > oldVal && btn.Right == nil {
		btn.Right = node
		return nil
	}

	return errors.New("value already exists")
}

func (btn *BinaryTreeNode) GetNode(value int) (*BinaryTreeNode, error) {
	if btn.Value == value {
		return btn, nil
	} else if value < btn.Value && btn.Left != nil {
		return btn.Left.GetNode(value)
	} else if value > btn.Value && btn.Right != nil {
		return btn.Right.GetNode(value)
	}
	return nil, errors.New("value could not be found")
}

//func (btn *BinaryTreeNode) Remove(value int) {
//
//}

//func (btn *BinaryTreeNode) GetMaxNode() int {
//
//}
//
//func (btn *BinaryTreeNode) GetMinNode() int {
//
//}
//
//func (btn *BinaryTreeNode) Balance() *BinaryTreeNode {
//
//}
