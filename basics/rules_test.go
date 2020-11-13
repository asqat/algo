package basics

import "testing"

func TestContainDuplicates(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainDuplicates(tt.args.arr); got != tt.want {
				t.Errorf("ContainDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindLargest(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindLargest(tt.args.arr); got != tt.want {
				t.Errorf("FindLargest() = %v, want %v", got, tt.want)
			}
		})
	}
}
