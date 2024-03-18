package main

import (
	"fmt"
	"net/http"
	"strings"
)

func HTTPRequest(input JokeInput) error {
	// If the categories are not set, we want to get jokes from any category
	if(input.Categories == nil) {
		input.Categories = []string{"Any"}
	}
	
	url := "https://v2.jokeapi.dev/joke/"+strings.Join(input.Categories, ",")
	
	// Because the API requires the language to be set, we will always set it
	url += "?lang="+input.Language

	// If the flags are set, we want to blacklist those flags
	if(input.Flags != nil) {
		url += "&blacklistFlags="+strings.Join(input.Flags, ",")
	}

	// If the amount is set, we want to get that amount of jokes
	if(input.Amount > 0) {
		url += "&amount="+fmt.Sprint(input.Amount)
	}

	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if(response.StatusCode != http.StatusOK) {
		return fmt.Errorf("received failed response from API: %s", response.Status)
	}

	// Maybe we want to do something with the response body?
	// bodyBytes, err := io.ReadAll(response.Body)
	// json := string(bodyBytes)
	// fmt.Println("API response: ", json)

	return nil
}