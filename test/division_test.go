package main

import (
	"testing"
)

func TestDivision(t *testing.T) {
	type args struct {
		x float64
		y float64
	}

	type testCase struct {
		name string
		args args
		want interface{}
	}

	testCases := []testCase{
		{
			name: "positive",
			args: args{
				x: 10.0,
				y: 2.0,
			},
			want: 5.0,
		},
		{
			name: "negative",
			args: args{
				x: 5.0,
				y: -2.0,
			},
			want: 5.0 / (-2.0),
		},
		{
			name: "divide by 0",
			args: args{
				x: 323.0,
				y: 0.0,
			},
			want: "can not divide by 0",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := Division(tt.args.x, tt.args.y)
			if got != tt.want {
				t.Errorf("expected: %v, but got: %v", tt.want, got)
			}
		})
	}
}
