package fizzbuzz

import (
	"fmt"
	"testing"
)

func TestGameGenerate(t *testing.T) {
	var game Game
	for i := 0; i < 15; i++ {
		output := game.Generate()
		if output != examples[i].Output {
			t.Fatalf("Expected %v got %v instead\n", examples[i].Output, output)
		}
	}
}

func TestVerify(t *testing.T) {
	for name, test := range map[string]struct {
		Tests []example
		Valid bool
	}{
		"valid": {Tests: examples, Valid: true},
		"invalid": {
			Tests: []example{
				{1, "fizzbuzz"},
				{2, "buzz"},
				{3, "3"},
			},
		},
	} {
		t.Run(name, func(t *testing.T) {
			var game Game
			game.count = 0
			for _, input := range test.Tests {
				output := game.Verify(input.Output)
				if (output == "") != test.Valid {
					t.Errorf("Expected %v, got %v", test.Valid, !test.Valid)
				}
			}
		})
	}
}

func validStr(b bool) string {
	if b {
		return "valid"
	}
	return "invalid"
}

func TestGenerate(t *testing.T) {
	for _, test := range examples {
		output := generate(test.Input)
		if output != test.Output {
			t.Fatalf(
				"For (%v) expected \"%v\" but got \"%v\"\n",
				test.Input,
				test.Output,
				output,
			)
		}
	}
}

type example struct {
	Input  int
	Output string
}

func genNum(i int) example { return example{i, fmt.Sprintf("%d", i)} }

var examples = []example{
	genNum(1),
	genNum(2),
	{3, "fizz"},
	genNum(4),
	{5, "buzz"},
	{6, "fizz"},
	genNum(7),
	genNum(8),
	{9, "fizz"},
	{10, "buzz"},
	genNum(11),
	{12, "fizz"},
	genNum(13),
	genNum(14),
	{15, "fizzbuzz"},
}
