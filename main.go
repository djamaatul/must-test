package main

import (
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("\n======= SOAL 1 =======")
	var letters = []string{
		"italem irad irigayaj",
		"iadab itsap ulalreb",
		"nalub kusutret gnalali",
	}
	for _, i := range letters {
		fmt.Println(reverse(i))
	}

	fmt.Println("\n======= SOAL 2 =======")
	for _, result := range generateNumber(100) {
		fmt.Println(result)
	}

	fmt.Println("\n======= SOAL 3 =======")
	fmt.Println(fibonnaci(20))

	fmt.Println("\n======= SOAL 4 =======")
	test4 := [][]int{
		{7, 8, 3, 10, 8},
		{5, 12, 11, 12, 10},
		{7, 18, 27, 10, 29},
		{20, 17, 15, 14, 10},
	}

	for _, numbers := range test4 {
		fmt.Println("Max Profit of :", hightProfit(numbers))
	}

	fmt.Println("\n======= SOAL 5 =======")
	tests5 := [][]interface{}{
		{"b", 7, "h", 6, "h", "k", "i", 5, "g", 7, 8},
		{7, "b", 8, 5, 6, 9, "n", "f", "y", 6, 9},
		{"u", "h", "b", "n", 7, 6, 5, 1, "g", 7, 9},
		{"12", "asda", "OK", nil, false, true, 1, "AAA"},
	}

	for _, strings := range tests5 {
		fmt.Println("total number :", countNumber(strings))
	}

}

// SOAL 1
func reverse(letter string) (result string) {
	splited := strings.Split(letter, " ")
	var reversed = []string{}
	for _, e := range splited {
		s := strings.Split(e, "")
		slices.Reverse(s)
		reversed = append(reversed, strings.Join(s, ""))
	}
	result = strings.Join(reversed, " ")
	return
}

// SOAL 2
func generateNumber(n int) (results []string) {
	for i := range n {
		element := i + 1
		result := strconv.Itoa(element) + " "
		if element%3 == 0 {
			result += "Fizz"
		}
		if element%5 == 0 {
			result += "Buzz"
		}
		results = append(results, result)
	}
	return
}

// SOAL 3
func fibonnaci(n int) (result []int) {
	var a, b int
	for range n {
		if len(result) == 0 {
			result = append(result, a)
			b = 1
			continue
		}
		if len(result) > 1 {
			b = result[len(result)-1]
			a = result[len(result)-2]
		}
		result = append(result, a+b)
	}
	return
}

// SOAL 4
func hightProfit(numbers []int) int {
	var highest int
	var highestFrom int
	for index, number := range numbers {
		if index < len(numbers)-1 {
			profit := numbers[index+1] - number
			if index == 0 {
				highest = profit
			}
			if highest < profit {
				highest = profit
				highestFrom = number
			}
		}
	}
	return highestFrom
}

// SOAL 5
func countNumber(numbers []interface{}) (result int) {
	regex, _ := regexp.Compile(`\d`)
	for _, variable := range numbers {
		var str string
		switch variable.(type) {
		case string:
			str = variable.(string)
		case int:
			str = strconv.Itoa(variable.(int))
		default:
			str = ""
		}
		if regex.MatchString(str) {
			result += 1
		}
	}
	return
}
