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
		{
			name: "TestContainDuplicates001",
			args: struct{ arr []int }{arr: []int{2, 3, 4, 5, 6, 7, 3}},
			want: true,
		},
		{
			name: "TestContainDuplicates002",
			args: struct{ arr []int }{arr: []int{2, 9, 4, 5, 6, 7, 3}},
			want: false,
		},
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
		{
			name: "TestFindLargest001",
			args: struct{ arr []int }{arr: []int{2, 4, 53, 13213, 31223, 1312, 2, 312, 312312, 3124124975, 56663}},
			want: 3124124975,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindLargest(tt.args.arr); got != tt.want {
				t.Errorf("FindLargest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDividingPoint(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "TestDividingPoint001",
			args: args{arr: []int{2, 3, 2, 5, 7, 8, 9, 3, 3, 32, 2, 2}},
			want: 5,
		},
		{
			name: "TestDividingPoint001",
			args: args{arr: []int{9, 3, 3, 32, 2, 2}},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DividingPoint(tt.args.arr); got != tt.want {
				t.Errorf("DividingPoint() = %v, want %v", got, tt.want)
			}
		})
	}
}
