package main

import (
	"errors"
	"reflect"
)

type Iter struct {
	items []interface{}

	// 实际实现均为 []
	filters func(interface{}) bool
	maps    func(interface{}) interface{}

	err error
}

func (i *Iter) Filter(f func(interface{}) bool) *Iter {
	i.filters = f
	return i
}

func (i *Iter) Map(f func(interface{}) interface{}) *Iter {
	i.maps = f
	return i
}

func (i *Iter) Collect() []interface{} {
	// cache
	res := []interface{}{}
	for _, item := range i.items {
		if i.filters != nil && !i.filters(item) {
			continue
		}
		if i.maps != nil {
			v := i.maps(item)
			res = append(res, v)
		} else {
			res = append(res, item)
		}
	}
	return res
}

func (i *Iter) Error() error {
	return i.err
}

func NewIter(items interface{}) *Iter {
	iter := &Iter{}

	// ignore is not a slice or array
	v := reflect.ValueOf(items)
	if v.Kind() != reflect.Slice {
		iter.err = errors.New("Iteror value not supported")
		return iter
	}
	for i := 0; i < v.Len(); i++ {
		iter.items = append(iter.items, v.Index(i).Interface())
	}

	return iter
}
