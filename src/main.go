//The package for this project
//A package is a collection of source files in the same directory that are compiled together (a project)
package main

//A list of all imports to be used in this file
import (
	"fmt"
	"os"
	"bufio"
)

/*
Main method of the program
This is automatically ran when the program starts (just like Java's main method)
*/
func main() {
	//Instantiate three new variables with the three values returned from getWordListFromFile
	words, uniqueWords, totalWords := getWordListFromFile()

	//Print a message overviewing the results of the test
	fmt.Printf("\n%d out of the %d words entered were unique.\nThose unique words are as follows:\n", uniqueWords, totalWords)

	//Loop through all of the elements in the words array and print them out
	for _, word := range words {
		//Go does not support nil strings, so every "empty" element in the words array is just a string of length 0
		//Check length before printing out to avoid printing the empty elements
		if len(word) > 0 {
			fmt.Println(word)
		}
	}
}

/*
Polls for input from the user and returns as a string.
Intended to be used to ask the user for the name of a text file
*/
func getInputFileName() string {
	//Create but do not instantiate a new string variable, intended to be the name of the input text file
	var inputFile string

	//Ask user for name of input text file and then read in input
	fmt.Print("Enter the name of the input text file: ")
	fmt.Scanln(&inputFile)

	//Return input
	return inputFile
}

/*
Reads and parses the content of a selected file
Repeatedly polls the user for a valid file until one is found
Once found, each unique line in the file is inserted into a string array
A maximum of 50 words will be read from the file
Returns three values: array of unique lines in the file, the number of unique words found, the number of words read in
*/
func getWordListFromFile() ([50]string, int, int) {
	//Poll the user for a file name
	var fileName = getInputFileName()

	//Attempt to read in that file
	file, err := os.Open(fileName)

	//err will be nil if the file was read in successfully
	if (err != nil) {
		//Failed to read in file, report error and try again (using recursion)
		fmt.Println("Unable to read file, try again...")
		return getWordListFromFile()
	}
	//Set the file to be closed
	//"Deferring" this means that this line will be ran once this method is done executing
	//This is the same concept as a "finally" block in Java
	defer file.Close()

	//Create a new scanner for the found file
	scanner := bufio.NewScanner(file)

	//Create variables to be used to parse the input of the file
	var result [50]string
	index, wordCount := 0, 0

	//A loop which will run once for each line of the file
	for scanner.Scan() {
		//If we read in 50 words already, end execution of this loop
		if (wordCount == 50) {
			break;
		}

		//Read in input for next new line and increment the word counter
		var input string = scanner.Text()
		wordCount++;

		//Check if we already read in this word before, we're only storing unique words
		if (!stringArrayContainsString(result, input)) {
			//This new word confirmed to be unique, store in array
			result[index] = input
			index++;
		}
	}

	//Done parsing file
	//Return results
	//"index" is - coincidentally - the total number of words read in
	return result, index, wordCount
}
/*
Returns whether or not the given array contains the given value
*/
func stringArrayContainsString(arr [50]string, val string) bool {
	//Loops through the given array
	for _, a := range arr {
		//If current element in array has same value as given value, return true
		if a == val {
			return true
		}
	}

	//Matching element not found, return false
	return false
}
