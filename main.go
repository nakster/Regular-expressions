package main

import (
	"fmt"
	"math/rand"
	"regexp"
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
	fmt.Println("\nI am not happy with your responses.")
	fmt.Println(elizaResponse("I am not happy with your responses."))
	//question 8
	fmt.Println("\nI am not sure that you understand the effect that your questions are having on me.")
	fmt.Println(elizaResponse("I am not sure that you understand the effect that your questions are having on me."))
	//question 9
	fmt.Println("\nI am supposed to just take what you’re saying at face value?")
	fmt.Println(elizaResponse("I am supposed to just take what you’re saying at face value?"))

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
	re := regexp.MustCompile(`(?i)I am ([^.?!]*)[.?!]?`)

	if matchedIAM := re.MatchString(takeInput); matchedIAM {

		return re.ReplaceAllString(takeInput, "How do you know you are $1?")

	}

	//return one of the following three strings from the string array.
	return response[rand.Intn(len(response))]

}
