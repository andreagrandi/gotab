package main

type Item struct {
	Id          string
	Description string
	Price       int
}

type TabInterface interface {
	AddItem(i *Item) error
	IsFull() bool
	Settle() error
	Total() int
}
