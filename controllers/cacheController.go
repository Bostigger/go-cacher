package controllers

import "github.com/bostigger/go-cacher/model"

func NewQueue() model.Queue {
	head := &model.Node{}
	tail := &model.Node{}

	head.Right = tail
	tail.Left = head
	return model.Queue{
		Head: head,
		Tail: tail,
	}
}
func NewCache() model.Cache {
	return model.Cache{
		Queue: NewQueue(),
		Hash:  model.Hash{},
	}
}
