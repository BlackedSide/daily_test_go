package test

import (
	"daily_test_go/service"
	"testing"
)

func TestShape(t *testing.T) {
	checkArea := func(t *testing.T, shape service.Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %.2f, want %.2f", got, want)
		}
	}

	t.Run("rectangles", func(t *testing.T) {
		rectangle := service.Rectangle{Width: 12, Height: 6}
		checkArea(t, rectangle, 72)
	})

	t.Run("circles", func(t *testing.T) {
		circle := service.Circle{Radius: 10}
		checkArea(t, circle, 314.1592653589793)
	})
}

func TestShapeByTable(t *testing.T) {
	areaTests := []struct {
		name  string
		shape service.Shape
		want  float64
	}{
		{"rectangles", service.Rectangle{Width: 12, Height: 6}, 72},
		{"circles", service.Circle{Radius: 10}, 314.1592653589793},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.want {
				t.Errorf("%#v got %.2f, want %.2f", tt.shape, got, tt.want)
			}
		})
	}
}
