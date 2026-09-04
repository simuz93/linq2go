package linq2go

import "testing"

func Test_Any(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want bool
	}{
		{
			name: "Any true",
			args: []int{1, 2, 3, 4},
			want: true,
		},
		{
			name: "Any false",
			args: []int{1, 2, 1, 2},
			want: false,
		},
		{
			name: "Empty",
			args: []int{},
			want: false,
		},
		{
			name: "Nil",
			args: nil,
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fn := func(v int) bool { return v > 2 }

			result := FromSlice(tt.args).Any(fn)

			if result != tt.want {
				t.Errorf("Any() = %v, want %v", result, tt.want)
			}
		})
	}
}
