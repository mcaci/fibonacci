package fibonacci

import "testing"

func TestNegativeInput(t *testing.T) {
	numbers := Serie(-1)
	if len(numbers) > 0 {
		t.Fatalf("Expecting len to be 0 but was %d", len(numbers))
	}
}

func TestFibonacci0Len(t *testing.T) {
	numbers := Serie(0)
	if len(numbers) > 0 {
		t.Fatalf("Expecting len to be 0 but was %d", len(numbers))
	}
}

func TestFibonacci10Len(t *testing.T) {
	numbers := Serie(10)
	if len(numbers) > 10 {
		t.Fatalf("Expecting len to be 10 but was %d", len(numbers))
	}
}

func TestFibonacci1(t *testing.T) {
	numbers := Serie(1)
	expected := uint64(1)
	actual := numbers[0]
	if expected != actual {
		t.Fatalf("Expecting last number of series to be %d but was %d", expected, actual)
	}
}

func TestFibonacci5(t *testing.T) {
	numbers := Serie(5)
	expected := uint64(5)
	actual := numbers[4]
	if expected != actual {
		t.Fatalf("Expecting last number of series to be %d but was %d", expected, actual)
	}
}

func TestFibonacci10(t *testing.T) {
	numbers := Serie(10)
	expected := uint64(55)
	actual := numbers[9]
	if expected != actual {
		t.Fatalf("Expecting last number of series to be %d but was %d", expected, actual)
	}
}
