package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UTC().UnixNano())
	// People say I look like both my mother and father.
	//Father was a teacher.
	//I was my father’s favourite.
	//I’m looking forward to the weekend.
	//My grandfather was French!

	fmt.Println("Hi I'm Eliza.")
	//question 1
	fmt.Println("People say I look like both my mother and father.")
	fmt.Println(elizaResponse("People say I look like both my mother and father."))

}

//The function should randomly return one of the following three strings.
var response = []string{
	"I’m not sure what you’re trying to say. Could you explain it to me?",
	"How does that make you feel?",
	"Why do you say that?",
}

func elizaResponse(takeInput string) string {

	return response[rand.Intn(len(response))]
}
