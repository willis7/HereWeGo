package main

import "fmt"

func main() {

	listOfNums := []float64{1,2,3,4,5}

	fmt.Println("Sum :", addThemUp(listOfNums))

	// Get 2 values from a function

	num1, num2 := next2Values(5)

	fmt.Println(num1, num2)

	// Send an undefined number of values to a function (Variadic Function)

	fmt.Println(subtractThem(1,2,3,4,5))


	// You can create a function in a function. It has access to the
	// local variables of the containing function
	// A function like this with no local variables is a closure

	num3 := 3

	doubleNum := func() int {

		num3 *= 2
		return num3

	}

	fmt.Println(doubleNum());
	fmt.Println(doubleNum());

	// Calling a recursive function

	fmt.Println(factorial(3))

	// Defer executes a function after the inclosing function finishes
	// Defer can be used to keep functions together in a logical way
	// but at the same time execute one last as a clean up operation
	// Ex. Defer closing a file after we open it and perform operations

	defer printTwo()
	printOne()

	// Use recover() to catch a division by 0 error

	fmt.Println(safeDiv(3, 0))
	fmt.Println(safeDiv(3, 2))

	// We can catch our own errors and recover with panic & recover

	demPanic()

}

// Functions allow us to reuse code and provide structure
// func funcName(parametersPassed) returnType
// Functions don't have access to any variables aside from those
// passed into it

func addThemUp(numbers []float64) float64 {

	sum := 0.0

	for _, val := range numbers {

		// Shorthand for sum = sum + val
		sum += val

	}

	return sum

}

// Go functions can return multiple values

func next2Values(number int) (int, int){

	return number+1, number+2

}

// You can receive an undefined number of values with args ...int

func subtractThem(args ...int) int{

	finalValue := 0

	for _, value := range args {
		finalValue -= value
	}

	return finalValue

}

// Example of recursion : Function calls itself
// factorial(3)
// 3 * factorial(2) == 3 * 2
// 2 * factorial(1) == 2 * 1
// factorial(0) == 1

func factorial(num int) int {
	if num == 0 {
		return 1
	}
	return num * factorial(num - 1)
}

// Used to demonstrate defer

func printOne(){ fmt.Println(1)}

func printTwo(){ fmt.Println(2)}

// If an error occurs we can catch the error with recover and allow
// code to continue to execute

func safeDiv(num1, num2 int) int {
    defer func() {
        fmt.Println(recover())
    }()
    solution := num1 / num2
    return solution
}

// Demonstrate how to call panic and handle it with recover

func demPanic(){

	defer func() {

		// If I didn't print the message nothing would show

		fmt.Println(recover())

	}()
	panic("PANIC")

}
