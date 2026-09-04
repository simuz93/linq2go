package linq2go

import "testing"

func Test_All(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want bool
	}{
		{
			name: "All true",
			args: []int{5, 6, 7, 8},
			want: true,
		},
		{
			name: "All false",
			args: []int{1, 2, 3, 4},
			want: false,
		},
		{
			name: "Empty",
			args: []int{},
			want: true,
		},
		{
			name: "Nil",
			args: nil,
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fn := func(v int) bool { return v > 2 }

			result := FromSlice(tt.args).All(fn)

			if result != tt.want {
				t.Errorf("All() = %v, want %v", result, tt.want)
			}
		})
	}
}
