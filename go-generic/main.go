package main

// func largest[T int | byte ](items []T) T {
// 	one := items[0]
// 	for _, item := range items {
// 		if item > one {
// 			one = item
// 		}
// 	}
// 	return one
// }

// type Stringer interface {
// 	String() string
// }

// func echo[T Stringer](s T) {
// 	print(s.String())
// }

func main() {
	// println(largest[int]([]int{1, 2, 3}))
	// println(largest([]byte{'a', 'b', 'c'}))
	println(NewValue(0).Or(1).Value())
	println(NewValue("").Or("default").Value())
	// println((*int)(nil))
	// println(NewValue[*int](nil).Unwrap(nil, nil).Value())
}
