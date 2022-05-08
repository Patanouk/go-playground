package _struct

import (
	"math"
	"testing"
)

func TestPerimeter(t *testing.T) {
	type args struct {
		shape Shape
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"Rectangle Perimeter", args{Rectangle{10,10}}, 40.0},
		{"Rectangle Perimeter", args{Circle{10}}, 2 * math.Pi * 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.shape.Perimeter(); got != tt.want {
				t.Errorf("Perimeter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArea(t *testing.T) {
	type args struct {
		shape Shape
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"Rectangle Area", args{Rectangle{12,6}}, 72.0},
		{"Circle Area", args{Circle{10}}, math.Pi * 10 * 10},
		{"Triangle Area", args{Triangle{12, 6}}, 36.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.shape.Area(); got != tt.want {
				t.Errorf("Shape: %#v, Area() = %v, want %v",tt.args.shape, got, tt.want)
			}
		})
	}
}