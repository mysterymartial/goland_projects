package main

func Add(number1, number2 int) int {
	return number1 + number2

}
func subtract(number1, number2 int) int {
	return number1 - number2

}
func multiply(number1, number2 int) int {
	return number1 * number2

}
func evenNumber(numbers []int) []int {
	var even []int
	for _, number := range numbers {
		if number%2 == 0 {
			even = append(even, number)

		}
	}
	return even
}
func divide(number1, number2 int) int {
	return number1 / number2

}
func oddNumber(numbers []int) []int {
	var odd []int
	for _, number := range numbers {
		if number%2 != 0 {
			odd = append(odd, number)
		}
	}
	return odd
}
func reverseArray(numbers []int) []int {
	for count, counter := 0, len(numbers)-1; count < counter; count, counter = count+1, counter-1 {
		numbers[count], numbers[counter] = numbers[counter], numbers[count]
	}
	return numbers

}
func sortArray(numbers []int) []int {

	for count := 0; count < len(numbers)-1; count++ {
		if numbers[count] > numbers[count+1] {
			temp := numbers[count]
			numbers[count] = numbers[count+1]
			numbers[count+1] = temp
		}
	}
	return numbers

}
func isPalindrome(word string) bool {
	reverse := ""
	for count := len(word) - 1; count >= 0; count-- {
		reverse = reverse + string(word[count])
		if word == reverse {
			return true
		}

	}
	return false
}
func isLeapYear(year int) bool {
	if year%4 == 0 && year%100 != 0 {
		return true
	}
	return false

}
func factorial(number int) int {
	factorial := 1
	for count := 1; count <= number; count++ {
		factorial *= count
	}
	return factorial
}
func firstDuplicate(numbers []int) int {
	for count := 0; count < len(numbers); count++ {
		for counter := count + 1; counter < len(numbers); counter++ {
			if numbers[count] == numbers[counter] {
				return numbers[count]
			}
		}
	}
	return -1

}
