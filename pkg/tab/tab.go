package tab

import "time"

type Item struct{}

type TabInterface interface {
	AddItem(i *Item) error
	IsFull() bool
	Settle() error
	Total() int
}

type Tab struct {
	Id         string
	Currency   string
	Limit      int
	CreatedAt  time.Time
	Items      []*Item
	Status     string
	CustomerId string
}
