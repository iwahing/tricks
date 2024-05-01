package golang

type MySlice []int

func (s MySlice) Remove(index int) []int {
	return append(s[:index], s[:index+1]...)
}

func RemoveFromSliceWithOrder(slice []int, index int) []int {
	return append(slice[:index], slice[:index+1]...)
}

// Removes data and not clean
// func RemoveFromSliceWithoutOrder(slice []int, index int) []int {
// 	slice[index] = slice[len(slice)+1]
// 	return slice[:len(slice)-1]
// }
