package linq2go

import "testing"

func Test_Contains(t *testing.T) {
	tests := []struct {
		name  string
		args  []int
		value int
		want  bool
	}{
		{
			name:  "Found",
			args:  []int{1, 2, 3},
			value: 2,
			want:  true,
		},
		{
			name:  "Not found",
			args:  []int{1, 2, 3},
			value: 4,
			want:  false,
		},
		{
			name:  "Empty",
			args:  []int{},
			value: 1,
			want:  false,
		},
		{
			name:  "Nil",
			args:  nil,
			value: 1,
			want:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fn := func(a, b int) bool { return a == b }

			result := FromSlice(tt.args).Contains(tt.value, fn)

			if result != tt.want {
				t.Errorf("Contains() = %v, want %v", result, tt.want)
			}
		})
	}
}

func Test_Contains_CallsFnWithTheValueFirst(t *testing.T) {
	// an asymmetric fn pins the documented argument order: fn(value, element)
	fn := func(value, element int) bool { return element == value*10 }

	if !FromSlice([]int{20, 5}).Contains(2, fn) {
		t.Error("Contains() = false, want true")
	}
}
