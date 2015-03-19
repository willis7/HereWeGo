package main

import ("fmt"
"strings"
"sort"
"os"
"log"
"io/ioutil"
"strconv")

func main() {

	// STRING FUNCTIONS

	sampString := "Hello World"

	// Returns true if phrase exists in string
	fmt.Println(strings.Contains(sampString, "lo"))

	// Returns the index for the match
	fmt.Println(strings.Index(sampString, "lo"))

	// Returns the number of matches for the string
	fmt.Println(strings.Count(sampString, "l"))

	// Replaces the first letter with the second as many times as you define
	fmt.Println(strings.Replace(sampString, "l", "x", 3))

	// Return a list separating with the defined separator
	csvString := "1,2,3,4,5,6"
	fmt.Println(strings.Split(csvString, ","))

	listOfLetters := []string{"c", "a", "b"}
	sort.Strings(listOfLetters)
	fmt.Println("Letters:", listOfLetters)

	// Returns a string using the values passed in separated with separator
	listOfNums := strings.Join([]string{"3", "2", "1"}, ", ");

	fmt.Println(listOfNums);

	// FILE I/O

	// Create a file
	file, err := os.Create("samp.txt")

	// Output any errors
	if err != nil {
		log.Fatal(err)
	}

	// Write a string to the file
	file.WriteString("This is some random text")

	// Close the file
	file.Close()

	// Try to open the file
	stream, err := ioutil.ReadFile("samp.txt")

	if err != nil {
		log.Fatal(err)
	}

	// Convert into a string
	readString := string(stream)

	fmt.Println(readString)

	// EXCEPTING INPUT

	fmt.Println("What is your name? ")

	var name string

	fmt.Scan(&name)

	fmt.Println("Hello", name)

	// CASTING

	randInt := 5
	randFloat := 10.5
	randString := "100"
	randString2 := "250.5"

	// Convert numbers types
	fmt.Println(float64(randInt))
	fmt.Println(int(randFloat))

	// Convert a string into an int
	newInt, _ := strconv.ParseInt(randString, 0, 64)
    fmt.Println(newInt)

    // Convert a string into a float
    newFloat, _ := strconv.ParseFloat(randString2, 64)
    fmt.Println(newFloat)

}
