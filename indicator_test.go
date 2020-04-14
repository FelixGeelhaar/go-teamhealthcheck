package teamhealthcheck

import (
	"fmt"
	"testing"
)

func TestRed(t *testing.T) {
	indicator := Indicator{}

	actual := indicator.Red()
	expected := 0

	if actual != expected {
		t.Errorf("The actual value is not equal to the expected value: %#v != %#v", actual, expected)
	}
}

func TestSetRed(t *testing.T) {
	indicator := Indicator{}

	indicator.SetRed(1)
	actual := indicator.Red()
	expected := 1

	if actual != expected {
		t.Errorf("The actual value is not equal to the expected value: %#v != %#v", actual, expected)
	}
}

func TestYellow(t *testing.T) {
	indicator := Indicator{}

	actual := indicator.Yellow()
	expected := 0

	if actual != expected {
		t.Errorf("The actual value is not equal to the expected value: %#v != %#v", actual, expected)
	}
}

func TestSetYellow(t *testing.T) {
	indicator := Indicator{}

	indicator.SetYellow(1)
	actual := indicator.Yellow()
	expected := 1

	if actual != expected {
		t.Errorf("The actual value is not equal to the expected value: %#v != %#v", actual, expected)
	}
}

func TestGreen(t *testing.T) {
	indicator := Indicator{}

	actual := indicator.Green()
	expected := 0

	if actual != expected {
		t.Errorf("The actual value is not equal to the expected value: %#v != %#v", actual, expected)
	}
}

func TestSetGreen(t *testing.T) {
	indicator := Indicator{}

	indicator.SetGreen(1)
	actual := indicator.Green()
	expected := 1

	if actual != expected {
		t.Errorf("The actual value is not equal to the expected value: %#v != %#v", actual, expected)
	}
}

func TestIndicator(t *testing.T) {
	indicator := Indicator{}

	actual := fmt.Sprintf("%#v", indicator.Indicator())
	expected := fmt.Sprintf("%#v", Indicator{})

	if actual != expected {
		t.Errorf("The actual value is not equal to the expected value: %#v != %#v", actual, expected)
	}
}
