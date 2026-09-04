package linq2go

import (
	"maps"
	"strings"
	"testing"
)

func Test_Transform(t *testing.T) {
	tests := []struct {
		name string
		args map[string]int
		want map[int]string
	}{
		{
			name: "Swap key and value",
			args: map[string]int{"a": 1, "b": 2},
			want: map[int]string{1: "a", 2: "b"},
		},
		{
			name: "Empty",
			args: map[string]int{},
			want: map[int]string{},
		},
		{
			name: "Nil",
			args: nil,
			want: map[int]string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fn := func(k string, v int) (int, string) { return v, k }

			result := FromMap(tt.args).Transform(fn).ToMap()

			if result == nil {
				t.Fatal("Transform().ToMap() = nil, want an empty map, never nil")
			}
			if !maps.Equal(result, tt.want) {
				t.Errorf("Transform().ToMap() = %v, want %v", result, tt.want)
			}
		})
	}
}

func Test_Transform_KeepsOneEntryOnKeyCollision(t *testing.T) {
	fn := func(k string, v int) (int, int) { return 0, v }

	// which entry survives a collision is not predictable, only that exactly one does
	result := FromMap(map[string]int{"a": 1, "b": 2}).Transform(fn).ToMap()

	if len(result) != 1 {
		t.Fatalf("Transform().ToMap() has %d entries, want 1: %v", len(result), result)
	}
	if v := result[0]; v != 1 && v != 2 {
		t.Errorf("Transform().ToMap()[0] = %v, want 1 or 2", v)
	}
}

func Test_ChangeKey(t *testing.T) {
	tests := []struct {
		name string
		args map[string]int
		want map[string]int
	}{
		{
			name: "Upper-case keys",
			args: map[string]int{"a": 1, "b": 2},
			want: map[string]int{"A": 1, "B": 2},
		},
		{
			name: "Empty",
			args: map[string]int{},
			want: map[string]int{},
		},
		{
			name: "Nil",
			args: nil,
			want: map[string]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fn := func(k string, v int) string { return strings.ToUpper(k) }

			result := FromMap(tt.args).ChangeKey(fn).ToMap()

			if result == nil {
				t.Fatal("ChangeKey().ToMap() = nil, want an empty map, never nil")
			}
			if !maps.Equal(result, tt.want) {
				t.Errorf("ChangeKey().ToMap() = %v, want %v", result, tt.want)
			}
		})
	}
}
