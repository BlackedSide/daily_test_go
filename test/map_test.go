package test

import (
	"fmt"
	"testing"
)

type Data struct {
	Name string
	Age  int
}

func TestMap(t *testing.T) {
	m := make(map[string]map[int]Data)
	m["fan"] = make(map[int]Data)
	m["fan"][30] = Data{
		Name: "Fan",
		Age:  30,
	}
	fmt.Printf("Yuan: %v\n", m["yuan"][20])
}
