package main

import "fmt"

type cipher interface {
	Encrypt() []int
}
type InpString string
type IntArr []int
type InpMap map[string]int

// encription method for string
func (s InpString) Encrypt() []int {
	result := make([]int, len(s))
	for i, char := range s {
		if int(char) > 95 {
			result[i] = int(char) - 96
		} else {
			result[i] = int(char) - 64
		}
	}
	return result
}

// encription method for int array

func (arr IntArr) Encrypt() []int {
	return arr
}

// encription method for map
func (m InpMap) Encrypt() []int {
	result := make([]int, len(m))
	i := 0
	for char, val := range m {
		result[i] = int(val) + int(char[0])
		i++
	}
	return result
}

func main() {
	var userInput cipher
	fmt.Print("Enter a string, an int array, or a map: ")

	var input interface{}
	fmt.Scan(&input)

	switch v := input.(type) {
	case string:
		userInput = InpString(v)
	case []int:
		userInput = IntArr(v)
	case map[string]int:
		userInput = InpMap(v)
	default:
		fmt.Println("Invalid input type")
		return
	}

	result := userInput.Encrypt()

	fmt.Println("Encryption Result:", result)
}
