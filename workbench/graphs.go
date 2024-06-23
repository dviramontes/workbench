package workbench

import (
	// "fmt"
	"math"
)

// CurrencyPair example problem:
//
// convert_currency([
//
//	["USD", "JPY", 110],
//	["USD", "AUD", 1.45],
//	["JPY", "GBP", 0.0070]],
//	"GBP",
//	"AUD"
//
// ) = 1.89

type CurrencyPair struct {
	FromCurrency string
	ToCurrency   string
	ExchangeRate float64
}

type Graph map[string]map[string]float64

type set map[string]bool

func ConvertCurrency(currencies []CurrencyPair, fromCurrency string, toCurrency string) float64 {
	// build the graph from the currency pairs
	var graph = make(Graph)
	for _, c := range currencies {
		if _, ok := graph[c.FromCurrency]; !ok {
			graph[c.FromCurrency] = make(map[string]float64)
		}
		if _, ok := graph[c.ToCurrency]; !ok {
			graph[c.ToCurrency] = make(map[string]float64)
		}
		// round to 2 decimal places
		graph[c.FromCurrency][c.ToCurrency] = RoundTo(c.ExchangeRate, 2)
		graph[c.ToCurrency][c.FromCurrency] = RoundTo(1/c.ExchangeRate, 2)
	}

	return DFS(graph, fromCurrency, toCurrency, 1.0, make(set))
}

// round a decimal to 2 decimal places
func RoundTo(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

func DFS(graph Graph, curr string, target string, pathValue float64, visited set) float64 {
	if curr == target {
		return pathValue
	}

	visited[curr] = true // mark as visited

	// iterate over the neighbors
	for neighbor, _ := range graph[curr] {
		if !visited[neighbor] {
			// recurse
			edgeValue := graph[curr][neighbor]
			if result := DFS(graph, neighbor, target, pathValue*edgeValue, visited); result != -1.0 {
				return result
			}
		}
	}

	// not found
	return -1.0
}
