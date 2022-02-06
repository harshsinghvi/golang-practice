package client

import (
	"./model"
	"encoding/json"
	"log"
	"net/http"
)

const URL string = "https://todo-app-harsh.herokuapp.com" + "/api/todos"

//Fetch is exported ...
func FetchTodos() ([]model.Todo, error) {
	//We make HTTP request using the Get function
	resp, err := http.Get(URL)
	if err != nil {
		log.Fatal("ooopsss an error occurred, please try again")
	}
	defer resp.Body.Close()
	//Create a variable of the same type as our model
	var cResp model.Todo
	//Decode the data
	if err := json.NewDecoder(resp.Body).Decode(&cResp); err != nil {
		log.Fatal("ooopsss! an error occurred, please try again")
	}
	//Invoke the text output function & return it with nil as the error value
	return cResp.TextOutput(), nil
}
