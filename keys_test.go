package linq2go

import (
	"slices"
	"testing"
)

func Test_Keys(t *testing.T) {
	tests := []struct {
		name string
		args map[string]int
		want []string
	}{
		{
			name: "Keys",
			args: map[string]int{"a": 1, "b": 2},
			want: []string{"a", "b"},
		},
		{
			name: "Empty",
			args: map[string]int{},
			want: []string{},
		},
		{
			name: "Nil",
			args: nil,
			want: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FromMap(tt.args).Keys()

			if result == nil {
				t.Fatal("Keys() = nil, want an empty slice, never nil")
			}

			// the order is unspecified by contract, so compare sorted
			slices.Sort(result)
			if !slices.Equal(result, tt.want) {
				t.Errorf("Keys() = %v, want %v", result, tt.want)
			}
		})
	}
}
