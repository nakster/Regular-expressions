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
	fmt.Println("Hi I'm Eliza.")
	//question 1
	fmt.Println("People say I look like both my mother and father.")
	fmt.Println(elizaResponse("People say I look like both my mother and father."))
	//question 2
	fmt.Println("Father was a teacher.")
	fmt.Println(elizaResponse("Father was a teacher."))
	//question 3
	fmt.Println("I was my father’s favourite.")
	fmt.Println(elizaResponse("I was my father’s favourite."))
	//question 4
	fmt.Println("I’m looking forward to the weekend.")
	fmt.Println(elizaResponse("I’m looking forward to the weekend."))
	//question 5
	fmt.Println("My grandfather was French!")
	fmt.Println(elizaResponse("My grandfather was French!"))

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

	//return one of the following three strings.
	return response[rand.Intn(len(response))]

}
