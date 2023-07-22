package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Country struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Please provide atleast 3 arguments to the program: FirstName, MiddleName(optional), LastName, Country Code")
		return
	}

	firstName := os.Args[1]
	lastName := os.Args[len(os.Args)-2]
	countryCode := os.Args[len(os.Args)-1]

	hasMiddleName := len(os.Args) > 4
	var middleName string
	if hasMiddleName {
		middleName = strings.Join(os.Args[2:len(os.Args)-2], " ")
	}

	file, err := os.Open("countries.json")
	if err != nil {
		fmt.Println("Failed to open file", err)
		return
	}

	defer file.Close()

	data, err := ioutil.ReadAll(file)

	if err != nil {
		fmt.Println("Failed to read file", err)
		return
	}

	// Parse JSON data
	var countries []Country
	err = json.Unmarshal(data, &countries)
	if err != nil {
		fmt.Println("Failed to parse JSON:", err)
		return
	}

	isValid := false
	for _, country := range countries {
		if country.Code == countryCode {
			isValid = true
			break
		}
	}

	if !isValid {
		fmt.Println("wrong country code")
		return
	}

	easternNameOrder := []string{"CN", "JP", "KR", "VN", "KP", "HU", "MN"}
	isEastern := contains(easternNameOrder, countryCode)

	if isEastern {
		printName(lastName, middleName, firstName)
	} else {
		printName(firstName, middleName, lastName)
	}
}

func contains(slice []string, value string) bool {
	for _, item := range slice {
		if item == value {
			return true
		}
	}
	return false
}

func printName(names ...string) {
	var formattedNames []string
	for _, name := range names {
		if name != "" {
			formattedNames = append(formattedNames, name)
		}
	}
	fmt.Println(strings.Join(formattedNames, " "))
}
