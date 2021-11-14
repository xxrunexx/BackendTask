// # library/consumeApi.go

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// A Response struct to map the Entire Response
type StarWarsPeople struct {
	Name      string   `json:"name"`
	Height    string   `json:"height"`
	Mass      string   `json:"mass"`
	HairColor string   `json:"hair_color"`
	SkinColor string   `json:"skin_color"`
	EyeColor  string   `json:"eye_color"`
	BirthYear string   `json:"birth_year"`
	Gender    string   `json:"gender"`
	Films     []string `json:"films"`
}

func main() {
	// Request data to API
	response, _ := http.Get("https://swapi.dev/api/people/1")

	// Read data from variable response and
	// get response.body using library ioutil.Readall
	responseData, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()

	// Parsing data to Data structure named LukeSkyWalker
	LukeSkyWalker := StarWarsPeople{}
	json.Unmarshal(responseData, &LukeSkyWalker)

	fmt.Println("---- Print Field ----")
	fmt.Println(LukeSkyWalker.Name)
	fmt.Println(LukeSkyWalker.Height)
	fmt.Println(LukeSkyWalker.Mass)
	fmt.Println(LukeSkyWalker.HairColor)
	fmt.Println(LukeSkyWalker.Films[:])
}
