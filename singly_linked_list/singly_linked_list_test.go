package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToSlice(t *testing.T) {
	list := &LinkedList{}
	list.Insert(1)
	list.Insert(2)
	list.Insert(3)

	result := list.ToSlice()

	assert.Equal(t, []int{1, 2, 3}, result)
}

func TestDisplay(t *testing.T) {
	list := &LinkedList{}
	list.Insert(1)
	list.Insert(2)
	list.Insert(3)

	list.Display()
}

func TestInsertFromSlice(t *testing.T) {
	l := &LinkedList{}
	l.InsertFromSlice([]int{1, 2, 3, 4})

	assert.Equal(t, []int{1, 2, 3, 4}, l.ToSlice())
}
