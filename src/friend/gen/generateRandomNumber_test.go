package gen

import (
	"testing"
	"fmt"
)

func BenchmarkRandomTestBase(b *testing.B) {
	number :=GenerateRandomNumber(0, 10000, 10000)
	fmt.Println(number)
}