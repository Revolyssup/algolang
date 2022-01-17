package tree

import (
	"errors"
	"reflect"
)

type Binary struct {
	Root *Node
	t    reflect.Type
}
type Node struct {
	Left  *Node
	Right *Node
	Value interface{}
	t     reflect.Type
}

func (node *Node) InsertLeft(val interface{}) error {
	if node.t.String() != reflect.TypeOf(val).String() {
		return errors.New("Expected value of type: " + node.t.String() + ", Gotten: " + reflect.TypeOf(val).String())
	}
	node.Left = &Node{
		Value: val,
		t:     node.t,
	}
	return nil
}
func (node *Node) InsertRight(val interface{}) error {
	if node.t.String() != reflect.TypeOf(val).String() {
		return errors.New("Expected value of type: " + node.t.String() + ", Gotten: " + reflect.TypeOf(val).String())
	}
	node.Right = &Node{
		Value: val,
		t:     node.t,
	}
	return nil
}

const ( //All possible values of order passed in Traverse Method.
	//Assigned three digit values can be read as [position of left][position of right][position of centre]
	CLR = 231
	CRL = 321
	LCR = 132
	LRC = 123
	RCL = 312
	RLC = 213
)

func (node *Node) Traverse(fn NodeCallbackFn, order int) error {
	left, right, curr := generateOrder(order)
	for !(left < 1 && right < 1 && curr < 1) {
		if left == 1 {
			if node.Left != nil {
				err := node.Left.Traverse(fn, order)
				if err != nil {
					return err
				}
			}
			left = 0
		}
		if right == 1 {
			if node.Right != nil {
				err := node.Right.Traverse(fn, order)
				if err != nil {
					return err
				}
			}
			right = 0
		}
		if curr == 1 {
			err := fn(node)
			if err != nil {
				return err
			}
			curr = 0
		}
		left--
		right--
		curr--
	}
	return nil
}
func generateOrder(order int) (left, right, curr int) {
	curr = order % 10
	order /= 10

	right = order % 10
	order /= 10

	left = order
	return left, right, curr
}

func (bt *Binary) Traverse(fn NodeCallbackFn, order int) error {
	return bt.Root.Traverse(fn, order)
}

func NewBinary(value interface{}, dtype interface{}) *Binary {
	t := reflect.TypeOf(dtype)
	return &Binary{
		t: t,
		Root: &Node{
			Value: value,
			t:     t,
		},
	}
}

func (bt *Binary) Type() string {
	return bt.t.String()
}
