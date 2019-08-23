package main

import "fmt"

func dec2roman(dec int) string {
	switch {
	case dec < 0:
		return "Error: Number cannot be less than 0"
	case dec == 0:
		return ""
	case dec == 1:
		return "I"
	case dec < 4:
		return dec2roman(dec-1) + dec2roman(1)
	case dec == 4:
		return dec2roman(1) + dec2roman(5)
	case dec == 5:
		return "V"
	case dec < 9:
		return dec2roman(5) + dec2roman(dec-5)
	case dec == 9:
		return dec2roman(1) + dec2roman(10)
	case dec == 10:
		return "X"
	case dec < 40:
		return dec2roman(10) + dec2roman(dec-10)
	case dec == 40:
		return dec2roman(10) + dec2roman(50)
	case dec < 50:
		return dec2roman(40) + dec2roman(dec-40)
	case dec == 50:
		return "L"
	case dec < 90:
		return dec2roman(50) + dec2roman(dec-50)
	case dec == 90:
		return dec2roman(10) + dec2roman(100)
	case dec < 100:
		return dec2roman(90) + dec2roman(dec-90)
	case dec == 100:
		return "C"
	default:
		return ("Error: Number cannot be greater than 100.")
	}
}

func main() {
	for i := -1; i < 102; i++ {
		fmt.Println(i, " = ", dec2roman(i))
	}
}
