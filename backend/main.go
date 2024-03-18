package main

import (
	"encoding/json"
	"log"
	"os"
)

// The struct to parse the JSON input with the values we will pass to the joke API
type JokeInput struct {
	Categories []string `json:"categories" validate:"required,dive,oneof=programming misc dark pun spooky christmas"`
	Language string `json:"language" validate:"required,oneof=cs de en es fr pt"`
	Flags []string `json:"flags" validate:"omitempty,dive,oneof=nsfw religious political racist sexist explicit"`
	Amount int `json:"amount" validate:"omitempty,min=1,max=10"`

}


func main(){
	// String to be parsed
	args := os.Args
	if len(args) < 2 {
		log.Fatal("Please provide a string argument")
	}
	jsonstring := args[1]

	// Parse the JSON string into a JokeInput struct
	var joke JokeInput
	err := json.Unmarshal([]byte(jsonstring), &joke)
	if err != nil {
		log.Fatalf("Error parsing input: %v", err)
	}

	// Validate the parsed input
	err = Validate(joke)
	if err != nil {
		log.Fatalf("Error validating parsed input: %v", err)
	}

	// Log the parsed input
	LogParsedInput(joke)

	// Make the HTTP request
	HTTPRequest(joke)
}