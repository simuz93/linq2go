package linq2go

import (
	"cmp"
	"slices"
)

// Sum returns the sum of the values selected by fn
func (s *slice[T]) Sum[K Number](fn func(T) K) K {
	var sum K = 0

	for _, v := range s.s {
		sum += fn(v)
	}

	return sum
}

// Min returns the smallest value selected by fn, or 0 if the slice is empty
func (s *slice[T]) Min[K Number](fn func(T) K) K {
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

// WhereMin returns the element whose value selected by fn is the smallest, or the zero value if the slice is empty
func (s *slice[T]) WhereMin[K Number](fn func(T) K) T {
	if len(s.s) == 0 {
		return *new(T)
	}

	cmp := func(s1 T, s2 T) int {
		return cmp.Compare(fn(s1), fn(s2))
	}

	return slices.MinFunc(s.s, cmp)
}

// Max returns the largest value selected by fn, or 0 if the slice is empty
func (s *slice[T]) Max[K Number](fn func(T) K) K {
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

// WhereMax returns the element whose value selected by fn is the largest, or the zero value if the slice is empty
func (s *slice[T]) WhereMax[K Number](fn func(T) K) T {
	if len(s.s) == 0 {
		return *new(T)
	}

	cmp := func(s1 T, s2 T) int {
		return cmp.Compare(fn(s1), fn(s2))
	}

	return slices.MaxFunc(s.s, cmp)
}

// Avg returns the average of the values selected by fn, or 0 if the slice is empty
func (s *slice[T]) Avg[K Number](fn func(T) K) float64 {
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

// Sum returns, for each group, the sum of the values selected by fn
func (g *group[K, V]) Sum[T Number](fn func(V) T) *dictionary[K, T] {
	result := map[K]T{}

	for k, v := range g.m {
		result[k] = newSlice(v).Sum(fn)
	}

	return newDictionary(result)
}

// Min returns, for each group, the smallest value selected by fn
func (g *group[K, V]) Min[T Number](fn func(V) T) *dictionary[K, T] {
	result := map[K]T{}

	for k, v := range g.m {
		result[k] = newSlice(v).Min(fn)
	}

	return newDictionary(result)
}

// Max returns, for each group, the largest value selected by fn
func (g *group[K, V]) Max[T Number](fn func(V) T) *dictionary[K, T] {
	result := map[K]T{}

	for k, v := range g.m {
		result[k] = newSlice(v).Max(fn)
	}

	return newDictionary(result)
}

// Avg returns, for each group, the average of the values selected by fn
func (g *group[K, V]) Avg[T Number](fn func(V) T) *dictionary[K, float64] {
	result := map[K]float64{}

	for k, v := range g.m {
		result[k] = newSlice(v).Avg(fn)
	}

	return newDictionary(result)
}
