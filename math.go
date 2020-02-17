package mathematic

// Sum ...
func Sum(a, b int) int {
	return sum(a, b)
}

func sum(a, b int) int {
	return (a + b)
}

// Subtract ...
func Subtract(a, b int) int {
	return (a - b)
}

// Multiply ...
func Multiply(a, b int) int {
	return (a * b)
}

// Division ...
func Division(a, b int) int {
	if b == 0 {
		return 0
	} else {
		return a / b
	}
}

i have added this line// ....
