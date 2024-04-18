package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

// package main
// import { libs }
// type 'name' struct { PropName PropType `json:"ParamNameInJson"`}
// var variableName type -> var can be optional if variable is immutable.

type ProductResult struct {
	Title             string  `json:"title"`
	Price             float64 `json:"price"`
	AvailableQuantity int     `json:"available_quantity"`
}

type Response struct {
	Results []ProductResult `json:"results"`
}

func main() {
	input := readInput()
	fmt.Println("You entered: ", input)

	results := search(input)
	for i, result := range results {
		fmt.Println(fmt.Sprintf(" [%d] %s", i, result.Title))
	}
}

func search(query string) []ProductResult {
	url := fmt.Sprintf("http://api.mercadolibre.com/sites/MLA/search?q=%s", query)
	response, err := http.Get(url)

	if err != nil {
		fmt.Println("Error querying MLA")
		return nil
	}

	bytes, err := io.ReadAll(response.Body)

	if err != nil {
		fmt.Println("Error reading response")
		return nil
	}

	var searchResponse Response
	err = json.Unmarshal(bytes, &searchResponse)

	if err != nil {
		fmt.Println("Error decoding json")
		return nil
	}

	return searchResponse.Results
}

func readInput() string {
	var input string
	fmt.Print("Busca un producto: ")
	_, err := fmt.Scan(&input)

	if err != nil {
		fmt.Println("Error reading input")
	}

	return input
}

func writeFile(path string, content []byte) error {
	err := os.WriteFile(path, content, os.ModeAppend)
	return err
}

func readFile(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer func() {
		err = file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}
