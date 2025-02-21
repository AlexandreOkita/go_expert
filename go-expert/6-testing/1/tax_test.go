package tax

import "testing"

func TestCalculateTax(t *testing.T) {
	amount := 500
	expected := 5.0
	result := CalculateTax(float64(amount))
	if result != expected {
		t.Errorf("Expected %f, got %f", expected, result)
	}
}

func TestCalculateTaxBatch(t *testing.T) {
	type calcTax struct {
		amount, expect float64
	}

	table := []calcTax{
		{500, 5.0},
		{1000, 10.0},
		{1500, 10.0},
		{0, 0},
	}

	for _, v := range table {
		result := CalculateTax(v.amount)
		if result != v.expect {
			t.Errorf("Expected %f, got %f", v.expect, result)
		}
	}
}

func BenchmarkCalculateTax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax2(500)
	}
}

func FuzzCalculateTax(f *testing.F) {
	seed := []float64{-1, -2, -2.5, 500, 1000, 1501.0}
	for _, amount := range seed {
		f.Add(amount)
	}
	f.Fuzz(func(t *testing.T, amount float64) {
		result := CalculateTax(amount)
		if amount <= 0 && result != 0 {
			t.Errorf("Received %f but expected 0", result)
		}
		if amount >= 20000 && result != 20 {
			t.Errorf("Received %f but expected 20", result)
		}
	})
}
