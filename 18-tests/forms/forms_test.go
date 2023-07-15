package forms_test

import (
	"math"
	"testing"
	"tests/forms"
)

func TestArea(t *testing.T) {
	t.Run("Rectangle", func(t *testing.T) {
		rectangle := forms.Rectangle{10, 12}
		expectedArea := float64(120)
		receivedArea := rectangle.Area()

		if expectedArea != receivedArea {
			t.Fatalf("The area should be %v, got %v", expectedArea, receivedArea)
		}
	})

	t.Run("Circle", func(t *testing.T) {
		radius := 15.0
		circle := forms.Circle{radius}
		expectedArea := float64(math.Pi * math.Pow(radius, 2))
		receivedArea := circle.Area()

		if expectedArea != receivedArea {
			t.Fatalf("The area should be %.2f, got %.2f", expectedArea, receivedArea)
		}
	})
}
