package main

import (
	"reflect"
	"testing"
)

func Test_findAnograms(t *testing.T) {
	type args struct {
		arr *[]string
	}
	tests := []struct {
		name string
		args args
		want *map[string][]string
	}{
		{
			name: "findAnograms",
			args: args{arr: &[]string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик"}},
			want: &map[string][]string{"листок": {"слиток", "столик"}, "пятак": {"пятка", "тяпка"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findAnograms(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findAnograms() = %v, want %v", got, tt.want)
			}
		})
	}
}
