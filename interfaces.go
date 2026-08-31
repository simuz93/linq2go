package linq2go

import "golang.org/x/exp/constraints"

// Number is the constraint satisfied by every integer and float type
type Number interface {
	constraints.Integer | constraints.Float
}
