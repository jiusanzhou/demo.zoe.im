package main

import "errors"

type Balancer interface {
	Next() interface{}
}

type RobinItem struct {
	Value interface{}
	Weight int
}

type RoundRobin struct {
	items []RobinItem
	weighted bool

	_pool []int
	_poolLen int
	_current int
}

func NewRoundRobin(items ...RobinItem) *RoundRobin {
	rr := &RoundRobin{
		items: items,
		_current: -1,
	}

	for _, t := range items {
		if t.Weight > 0 {
			rr.weighted = true
			break
		}
	}

	if !rr.weighted {
		for _, t := range items {
			t.Weight = 1
		}
	}
	
	// generate the _pool []

	// any else?

	rr._poolLen = len(rr._pool)

	return rr
}

func (rr *RoundRobin) Next() (interface{}, error) {
	if len(rr._pool) == 0 {
		return nil, errors.New("empty")
	}

	rr._current += 1

	if rr._current >= rr._poolLen {
		rr._current = 0
	}

	return rr._pool[rr._current], nil
}


// utils

func GetLCM(counts ...int) int {

	return 0
}

func GetSum(counts ...int) int {
	var r int
	for _, i := range counts{
		r += i
	}
	return r
}