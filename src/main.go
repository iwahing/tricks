package main

import (
	"fmt"

	golang "github.com/iwahing/tricks/golang"
	utils "github.com/iwahing/tricks/utils"
)

func main() {
	// Datetime manipulation
	yesterday, _ := utils.GetDateTimestamp("2024-04-29T12:00:00Z")
	fmt.Println(utils.GetDaysSinceUnixEpoch(yesterday))
	nineDaysInSeconds := utils.GetDaysInSeconds(9)
	fmt.Println(nineDaysInSeconds)
	newDate := utils.GetDateBefore(yesterday, nineDaysInSeconds)
	newDateTimestamp, _ := utils.GetDateTimestamp(newDate)
	fmt.Println(utils.UnixTimestampToUTCDateTimeString(newDateTimestamp))

	// Bitwise
	fmt.Println(utils.IsEven(241239))
	fmt.Println(utils.IsOdd(241234567))

	// Specific Golang Tricks
	list := []int{5, 77, 453, 123, 567}
	mySliceList := golang.MySlice{5, 77, 453, 123, 567}
	fmt.Println(golang.RemoveFromSliceWithOrder(list, 2))
	fmt.Println(mySliceList.Remove(2))
}
