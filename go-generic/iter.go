package main


// Iter is the struct to process items
type Iter[T1 any] struct {
	items []T1

	// TODO: array should be
	filters func(T1) bool

	err error
}

// Filter register the filter function
// TODO: filter should call as the map chain
func (i *Iter[T1]) Filter(f func(T1) bool) *Iter[T1] {
	i.filters = f
	return i
}

// Map register the map function to convert item
// TODO: map should call as a chain
func (i *Iter[T1]) Map[T2 any](f func(T1) T2) []T2 {
	// collect items

	res := []T2{}
	for _, item := range i.items {
		if i.filters != nil && !i.filters(item) {
			continue
		}
		res = append(res, f(item))
	}
	return res
	return i
}

// Collect to take items out
func (i *Iter[T1]) Collect() []T1 {
	if i.filters == nil {
		return i.items
	}

	res := []T1{}
	for _, item := range i.items {
		if i.filters != nil && !i.filters(item) {
			continue
		}
		res = append(res, item)
	}
	return res
}

// Error get the init error
func (i *Iter[T1]) Error() error {
	return i.err
}

// NewIter reutrn a new iter
func NewIter[T1 any](items []T1) *Iter[T1] {
	iter := &Iter[T1]{
		items: items,
	}

	return iter
}
