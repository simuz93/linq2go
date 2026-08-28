package linq2go

import (
	"cmp"
	"slices"
)

/*
Sum returns the sum of the values selected by the function
*/
func (s *Slice[T]) Sum[K Number](fn func(T) K) K {
	var sum K = 0

	for _, v := range s.s {
		sum += fn(v)
	}

	return sum
}

/*
Min returns the minimum value selected by the function, or 0 if the slice is empty
*/
func (s *Slice[T]) Min[K Number](fn func(T) K) K {
	if len(s.s) == 0 {
		return 0
	}

	var min K = fn(s.s[0])

	for i := 1; i < len(s.s); i++ {
		v := fn(s.s[i])
		if v < min {
			min = v
		}
	}

	return min
}

/*
WhereMin use the given function to select a field of each element and returns the element with the minimum value in that field, or default if the slice is empty.
Uses slices.MinFunc
*/
func (s *Slice[T]) WhereMin[K Number](fn func(T) K) T {
	if len(s.s) == 0 {
		return *new(T)
	}

	cmp := func(s1 T, s2 T) int {
		return cmp.Compare(fn(s1), fn(s2))
	}

	return slices.MinFunc(s.s, cmp)
}

/*
Max returns the maximum value selected by the function, or 0 if the slice is empty
*/
func (s *Slice[T]) Max[K Number](fn func(T) K) K {
	if len(s.s) == 0 {
		return 0
	}

	var max K = fn(s.s[0])

	for i := 1; i < len(s.s); i++ {
		v := fn(s.s[i])
		if v > max {
			max = v
		}
	}

	return max
}

/*
WhereMax use the given function to select a field of each element and returns the element with the maximum value in that field, or default if the slice is empty.
Uses slices.MinFunc
*/
func (s *Slice[T]) WhereMax[K Number](fn func(T) K) T {
	if len(s.s) == 0 {
		return *new(T)
	}

	cmp := func(s1 T, s2 T) int {
		return cmp.Compare(fn(s1), fn(s2))
	}

	return slices.MaxFunc(s.s, cmp)
}

/*
Avg returns the average of the values selected by the function, or 0 if the slice is empty
*/
func (s *Slice[T]) Avg[K Number](fn func(T) K) float64 {
	len := len(s.s)
	if len == 0 {
		return 0
	}
	sum := s.Sum(fn)
	if sum == 0 {
		return 0
	}

	return float64(sum) / float64(len)
}

/*
Sum returns the sum of the values selected by the function
*/
func (g *Group[K, V]) Sum[T Number](fn func(V) T) *Map[K, T] {
	result := map[K]T{}

	for k, v := range g.m {
		result[k] = FromSlice(v).Sum(fn)
	}

	return FromMap(result)
}

/*
Sum returns the sum of the values selected by the function
*/
func (g *Group[K, V]) Min[T Number](fn func(V) T) *Map[K, T] {
	result := map[K]T{}

	for k, v := range g.m {
		result[k] = FromSlice(v).Min(fn)
	}

	return FromMap(result)
}

/*
Sum returns the sum of the values selected by the function
*/
func (g *Group[K, V]) Max[T Number](fn func(V) T) *Map[K, T] {
	result := map[K]T{}

	for k, v := range g.m {
		result[k] = FromSlice(v).Max(fn)
	}

	return FromMap(result)
}

/*
Sum returns the sum of the values selected by the function
*/
func (g *Group[K, V]) Avg[T Number](fn func(V) T) *Map[K, float64] {
	result := map[K]float64{}

	for k, v := range g.m {
		result[k] = FromSlice(v).Avg(fn)
	}

	return FromMap(result)
}
