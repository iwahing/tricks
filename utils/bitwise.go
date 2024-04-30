package utils

// TODO : Testfile to benchmark with big
// 		  number if bitwise or modulo is faster

func IsEven(num int) bool {
	// num % 2 == 0
	// !IsOdd(num)
	return num&1 == 0
}

func IsOdd(num int) bool {
	// num % 2 == 1
	// !IsEven(num)
	return num&1 == 1
}
