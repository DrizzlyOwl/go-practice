package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Todos = []struct {
	uid  int    `json:"userId"`
	id   int    `json:"id"`
	task string `json:"title"`
	done bool   `json:"completed"`
}

func main() {
	data, err := getTodos(5)

	if err != nil {
		fmt.Println("Unexpected error:")
		fmt.Println(err.Error())
		return
	}

	fmt.Print(data)
}

func getTodos(count int) (string, error) {
	const baseUri string = "https://jsonplaceholder.typicode.com/todos/"
	var requestUri string = baseUri + fmt.Sprintf("%d", count)

	resp, err := http.Get(requestUri)
	fmt.Println("GET: " + requestUri)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var todos Todos
	err = json.Unmarshal(bytes, &todos)
	if err != nil {
		return "", err
	}

	for value := range todos {
		fmt.Printf("%+v", value)
	}

	return "", nil
}
