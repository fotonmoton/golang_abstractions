package examples

import (
	"fmt"
	"slices"
	"strings"
)

var numbers = []int{7, 5, 5, 2}

var strgs = []string{"c", "b", "v", "b"}

func SortExample() {

	slices.SortFunc(numbers, func(a, b int) int { return b - a })

	slices.SortFunc(strgs, func(a, b string) int { return strings.Compare(a, b) })

	fmt.Println(numbers)
	fmt.Println(strgs)
}

func PrintExample() {
	fmt.Println(numbers)
}
