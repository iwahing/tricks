package utils

func IsEven(num int) bool { // num % 2 == 0 or !IsOdd(num)
	return num&1 == 0
}

func IsOdd(num int) bool { // num % 2 == 1 or !IsEven(num)
	return num&1 == 1
}
