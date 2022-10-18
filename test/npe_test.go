package test

import (
	"daily_test_go/service"
	"fmt"
	"testing"
)

func TestNpe(t *testing.T) {
	fmt.Println(service.GetWords("B"))
	for _, c := range service.GetWords("B") {
		fmt.Println(c)
	}
}
