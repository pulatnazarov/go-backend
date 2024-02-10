package main

import (
	"testing"
)

func TestPerimeterRectangle(t *testing.T) {
	testCases := []struct {
		name string
		arg  Rectangle
		want interface{}
	}{
		{
			name: "normal case",
			arg: Rectangle{
				X: 20,
				Y: 30,
			},
			want: 100,
		},
		{
			name: "x = 0 case",
			arg: Rectangle{
				X: 0,
				Y: 30,
			},
			want: "X can not be 0",
		},
		{
			name: "y = 0 case",
			arg: Rectangle{
				X: 10,
				Y: 0,
			},
			want: "Y can not be 0",
		},
		{
			name: "X or Y negative",
			arg: Rectangle{
				X: -2,
				Y: 30,
			},
			want: "X or Y can not be negative",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			perimeter := PerimeterRectangle(tt.arg)
			if perimeter != tt.want {
				t.Errorf("expected %v, but got %v", tt.want, perimeter)
			}
		})
	}

}

func TestPerimeterSquare(t *testing.T) {
	testCases := []struct {
		name string
		arg  Square
		want interface{}
	}{
		{
			name: "normal case",
			arg: Square{
				X: 10,
			},
			want: 40,
		},
		{
			name: "when x < 0",
			arg: Square{
				X: -10,
			},
			want: "X can not negative",
		},
		{
			name: "when x = 0",
			arg: Square{
				X: 0,
			},
			want: "X can not be 0",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			perimeter := PerimeterSquare(tt.arg)
			if perimeter != tt.want {
				t.Errorf("expeceted %v, but got %v", tt.want, perimeter)
			}
		})
	}
}
