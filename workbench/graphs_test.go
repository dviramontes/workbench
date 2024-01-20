package workbench

import "testing"

func TestConvertCurrency(t *testing.T) {
	type args struct {
		currencies   []CurrencyPair
		fromCurrency string
		toCurrency   string
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "example problem",
			args: args{
				currencies: []CurrencyPair{
					{"USD", "JPY", 100},
					{"JPY", "CHN", 20},
					{"CHN", "THAI", 200},
				},
				fromCurrency: "USD",
				toCurrency:   "CHN",
			},
			want: 2000.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertCurrency(tt.args.currencies, tt.args.fromCurrency, tt.args.toCurrency); got != tt.want {
				t.Errorf("ConvertCurrency(), got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRoundTo(t *testing.T) {
	num := 10.050001
	expected := 10.05
	result := RoundTo(num, 2)
	if result != expected {
		t.Errorf("Expected %2f, but got %2f", expected, result)
	}
}
