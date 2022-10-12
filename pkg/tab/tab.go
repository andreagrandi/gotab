package tab

type Item struct{}

type TabInterface interface {
	AddItem(i *Item) error
	IsFull() bool
	Settle() error
	Total() int
}
