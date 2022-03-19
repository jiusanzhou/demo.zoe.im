package main

import (
	"reflect"
	"testing"
)

func TestProcess(t *testing.T) {
	cases := []struct {
		name   string
		value  interface{}
		filter func(value interface{}) bool
		fmap   func(value interface{}) interface{}
		result interface{}
	}{
		{
			"Normal", []int{0, 1, 2, 3}, nil, nil,
			[]interface{}{0, 1, 2, 3},
		},
		{
			"Filter", []int{0, 1, 2, 3}, func(v interface{}) bool {
				return v.(int) > 0
			}, nil,
			[]interface{}{1, 2, 3},
		},
		{
			"Convert", []int{0, 1, 2, 3}, func(v interface{}) bool {
				return v.(int) > 0
			}, func(v interface{}) interface{} {
				return v.(int) * 10
			},
			[]interface{}{10, 20, 30},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			r := NewIter(c.value).Filter(c.filter).Map(c.fmap).Collect()
			if !reflect.DeepEqual(c.result, r) {
				t.Errorf("want: %v, got: %v", c.result, r)
			}
		})
	}
}

func TestNewIter(t *testing.T) {
	cases := []struct {
		name        string
		value       interface{}
		expectError bool
	}{
		{
			"Nil Error", nil, true,
		},
		{
			"[]int OK", []int{}, false,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			x := NewIter(c.value)
			if c.expectError != (x.Error() != nil) {
				t.Errorf("want error %v, got %v", c.expectError, x.Error())
			}
		})
	}
}
