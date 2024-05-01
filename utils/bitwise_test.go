package utils

import (
	"fmt"
	"math/rand"
	"testing"
)

func generateTestData(numDigits int) []int {
	testData := make([]int, numDigits)
	for i := 0; i < numDigits; i++ {
		testData[i] = int(rand.Intn(999999999))
	}

	return testData
}

var testData = generateTestData(5000)

func TestIsEven(t *testing.T) {
	value := 12318230
	if !IsEven(value) {
		t.Errorf("IsEven(%d) is not even; want even", value)
	}
}

func TestIsOdd(t *testing.T) {
	value := 14258231
	if !IsOdd(value) {
		t.Errorf("IsOdd(%d) is not even; want even", value)
	}
}

func BenchmarkIsEven(b *testing.B) {
	for _, val := range testData {
		b.Run(fmt.Sprintf("Bitwise-%d", val), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				IsEven(val)
			}
		})
		// b.Run(fmt.Sprintf("Modulo-%d", val), func(b *testing.B) {
		// 	for i := 0; i < b.N; i++ {
		// 		if val % 2 == 0 {
		// 			pass
		// 		}
		// 	}
		// })
	}
}
