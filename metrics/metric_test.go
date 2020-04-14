package metrics

import (
	"fmt"
	"testing"
)

func TestNewMetric(t *testing.T) {
	metric := New("Test")

	actual := fmt.Sprintf("%#v", metric)
	expected := fmt.Sprintf("%#v", HealthMetric{title: "Test"})

	if actual != expected {
		t.Errorf("The actual value is not equal to the expected value: %#v != %#v", actual, expected)
	}
}

func TestSetTitle(t *testing.T) {
	metric := New("Test")

	metric.SetTitle("New Test")
	actual := metric.Title()
	expected := "New Test"

	if actual != expected {
		t.Errorf("The actual value is not equal to the expected value: %#v != %#v", actual, expected)
	}
}

func TestTitle(t *testing.T) {
	metric := New("Test")

	actual := metric.Title()
	expected := "Test"

	if actual != expected {
		t.Errorf("The actual value is not equal to the expected value: %#v != %#v", actual, expected)
	}
}

func TestTendency(t *testing.T) {
	metric := New("Test")

	actual := metric.Tendency()
	expected := 0

	if actual != expected {
		t.Errorf("The actual value is not equal to the expected value: %#v != %#v", actual, expected)
	}
}

func TestSetTendency(t *testing.T) {
	metric := New("Test")

	metric.SetTendency(5)
	actual := metric.Tendency()
	expected := 5

	if actual != expected {
		t.Errorf("The actual value is not equal to the expected value: %#v != %#v", actual, expected)
	}
}
