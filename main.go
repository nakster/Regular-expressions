package main
//adapted from https://gist.github.com/ianmcloughlin/c4c2b8dc586d06943f54b75d9e2250fe
//https://shapeshed.com/golang-regexp/
//imports
import (
	"fmt"
	"math/rand"
	"regexp"
	"strings"
	"time"
)

func main() {

	rand.Seed(time.Now().UTC().UnixNano())

	//Have main print both the input to and output from the function to the terminal.
	fmt.Println("\nHi I'm Eliza.")
	//question 1
	fmt.Println("People say I look like both my mother and father.")
	fmt.Println(elizaResponse("People say I look like both my mother and father."))
	//question 2
	fmt.Println("\nFather was a teacher.")
	fmt.Println(elizaResponse("Father was a teacher."))
	//question 3
	fmt.Println("\nI was my father’s favourite.")
	fmt.Println(elizaResponse("I was my father’s favourite."))
	//question 4
	fmt.Println("\nI’m looking forward to the weekend.")
	fmt.Println(elizaResponse("I’m looking forward to the weekend."))
	//question 5
	fmt.Println("\nMy grandfather was French!")
	fmt.Println(elizaResponse("My grandfather was French!"))
	//question 6
	fmt.Println("\nI am happy.")
	fmt.Println(elizaResponse("I am happy."))
	//question 7
	fmt.Println("\nI'm not happy with your responses.")
	fmt.Println(elizaResponse("I'm not happy with your responses."))
	//question 8
	fmt.Println("\nIm not sure that you understand the effect that your questions are having on me.")
	fmt.Println(elizaResponse("Im not sure that you understand the effect that your questions are having on me."))
	//question 9
	fmt.Println("\nI AM supposed to just take what you’re saying at face value?")
	fmt.Println(elizaResponse("I AM supposed to just take what you’re saying at face value?"))

	fmt.Println("\nGood Morning")
	fmt.Println(elizaResponse("Good Morning"))

}

//The function should randomly return one of the following three strings.
var response = []string{
	"I’m not sure what you’re trying to say. Could you explain it to me?",
	"How does that make you feel?",
	"Why do you say that?",
}

// takes a single string as input and returns a single string as output
func elizaResponse(takeInput string) string {

	//check, using a regular expression, if the input contains the word “father”.
	if matched, _ := regexp.MatchString(`(?i).*\bfather\b.*`, takeInput); matched {
		return "Why don’t you tell me more about your father?"
	}

	//if the input does not contain the word “father”, check the input begins with “I am “
	//Adapt the function to respond in the same way as with “i am “, “I AM “, “I’m “, “Im “, “i’m “
	re := regexp.MustCompile(`(?i)i(?:'|\sa)?m (.*)`)
	deleteIAm := re.FindStringSubmatch(takeInput)

	// List the reflections.
	reflections := [][]string{
		{"your", "my"},
		{"you’re", "i am"},
		{"I", "you"},
		{"you", "I"},
		{"me", "you"},
	}

	if len(deleteIAm) > 0 {

		splitInput := strings.Split(deleteIAm[1], " ")

		// Loop through each word, reflecting it if there's a match.
		for i, check := range splitInput {
			for _, reflection := range reflections {
				if matched, _ := regexp.MatchString(reflection[0], check); matched {
					splitInput[i] = reflection[1]
					break
				}
			}
		}

		// Put the string back together.
		joinString := strings.Join(splitInput, " ")

		return fmt.Sprintf("How do you know you are %s?", joinString)

	}

	//part 6(1)
	//make it lower case for find answer better
	takeInput = strings.TrimRight(takeInput, "\n.!")
	takeInput = strings.ToLower(takeInput)

	if matched, _ := regexp.MatchString(`(?i).*\bgood morning\b.*`, takeInput); matched {
		return choice(responseToGoodM)
	}



	//return one of the following three strings from the string array.
	return choice(response)

}

//this stores response to good morning
var responseToGoodM = []string{
	"Good Morning to you to",
	"Good Morning, how are you today?",
	"Bonjour having a nice day?",
}

//this func returns a random string from the array
func choice(ls []string) string {
	randnum := rand.Intn(len(ls))
	return ls[randnum]
}