package perimeters

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	result := Perimeter(rectangle)
	expect := 40.0

	if result != expect {
		t.Errorf("result %.2f expect %.2f", result, expect)
	}
}

func TestArea(t *testing.T) {
	testsArea := []struct {
		name   string
		form   Form
		expect float64
	}{
		{form: Rectangle{Width: 12, Height: 6}, expect: 72.0},
		{form: Circle{Radius: 10}, expect: 314.1592653589793},
		{form: Triangle{Base: 12, Height: 6}, expect: 36.0},
	}

	for _, tt := range testsArea {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.form.Area()

			if result != tt.expect {
				t.Errorf("%#v result %.2f expect %.2f", tt.form, result, tt.expect)
			}
		})
	}
}
