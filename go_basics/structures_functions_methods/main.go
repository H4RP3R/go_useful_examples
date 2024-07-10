package main

import (
	"encoding/json"
	"fmt"
	"time"
)

func main() {
	PrintPersonInfo(2525, "Anna", "Smith", time.Date(1999, 11, 24, 0, 0, 0, 0, time.UTC))
	PrintPersonAge(1234, time.Date(1987, 7, 1, 0, 0, 0, 0, time.UTC))

	x1, x2, err := SolveQuadraticEquation(4, -5, -12)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("x1: %v, x2 %v\n", x1, x2)

	someJsonContent := `{
		"field_a": 15,
		"field_b": "string will be written",
		"field_c": "string won't be written"
	}`

	ts := Tagged{}
	if err := json.Unmarshal([]byte(someJsonContent), &ts); err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", ts)
}
