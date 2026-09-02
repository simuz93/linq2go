package linq2go

import "golang.org/x/exp/constraints"

// Number is the constraint satisfied by every integer and float type
//
//	func (s *slice[T]) Sum[V Number](fn func(T) V) V
type Number interface {
	constraints.Integer | constraints.Float
}
