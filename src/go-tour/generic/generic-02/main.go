package main

import "fmt"

// List represents a singly-linked list that holds
// values of any type.
type List[T any] struct {
	next *List[T]
	val  T
}

func main() {
	emiliaCuteness := List[string]{
		val: "Emilia is the best girl",
		next: &List[string]{
			val: "Indeed she is!",
			next: &List[string]{
				val:  "No doubts about that!",
				next: nil,
			},
		},
	}

	emiliaTan := List[string]{
		next: &emiliaCuteness,
		val:  "Emilia-tan",
	}

	fmt.Printf("%+v %+v\n", emiliaTan, emiliaCuteness)
}
