package main

import (
	"io/ioutil"
	"log"
	"net/http"

	"bytes"
	"encoding/json"

	"fmt"
	"harshsinghvi/todo-app-cli/testModule"
	"os"
)

const URL string = "https://todo-app-harsh.herokuapp.com" + "/api/todos"

func main() {
	// for index, arg := range os.Args {
	// 	fmt.Printf("%d - %s\n", index, arg)
	// }
	
	// fmt.Println(os.Args)

	if len(os.Args) < 2{
		printHelp()
		return
	}

	// myProgramName := os.Args[0]

	switch os.Args[1] {
		case "list": listTodos()
		case "add": createTodo(os.Args[2])
		// case "delete":
		// case "complete":
		case "help": printHelp()
		case "greet": fmt.Println(testModule.Hello(os.Args[2]))
		default: printError(os.Args[1])
	}
}

func printError(actionCommand string){
	fmt.Println(actionCommand, "Command not found \n use help arg to get for list of supported args.")
}
func printHelp(){
	fmt.Println("List of Supported Args")
	fmt.Println("list             -- list todos")
	fmt.Println("add <name>       -- add a todo")
	fmt.Println("delete <oid>     -- delete a todo")
	fmt.Println("complete <oid>   -- mark complete to a todo")
	fmt.Println("help             -- to get for list of supported args")
	fmt.Println("greet <Username> -- to greet a user")
}

func listTodos() {
	resp, err := http.Get(URL)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	jsonData := []byte(string(body))

	if err != nil {
		log.Fatalln(err)
	}

	// fmt.Println(body)
	// fmt.Println(jsonData)

	defer resp.Body.Close()

	// listOfTodos := [25]Todo{}

	var listOfTodos []Todo

	err = json.Unmarshal(jsonData, &listOfTodos)

	if err != nil {

		// if error is not nil
		// print error
		fmt.Println(err)
	}
	completed := 0 
	for i := range listOfTodos {
		fmt.Println(listOfTodos[i].TextOutput())
		// fmt.Printf("%d\n",listOfTodos[i].Name)
		if listOfTodos[i].completed == true{
			completed++
		}
	}

	fmt.Println("Number of Todos:", len(listOfTodos))
	fmt.Println("Completed", completed)
	// if err := json.NewDecoder(resp.Body).Decode(&listOfTodos); err != nil {
	// 	log.Fatal("ooopsss! an error occurred, please try again")
	// }

	// fmt.Println(listOfTodos.TextOutput())

	// for _, t := range listOfTodos{
	// 	// fmt.Println(TextOutput(t))
	// 	t.TextOutput()
	// }
	// sb := string(body)
	// log.Printf(sb)
}

func createTodo(name string) {
	//Encode the data
	postBody, _ := json.Marshal(map[string]string{
		"name": name,
	})
	responseBody := bytes.NewBuffer(postBody)
	//Leverage Go's HTTP Post function to make request
	resp, err := http.Post(URL, "application/json", responseBody)
	//Handle Error
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer resp.Body.Close()
	fmt.Println(resp.Status)
	//Read the response body
	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// sb := string(body)
	// log.Printf(sb)
}

// Cryptoresponse is exported, it models the data we receive.
type Todo struct {
	Oid       string `json:"_id"`
	Name      string //`json:"name"`
	Date      string //`json:"date"`
	completed bool   //`json:"completed"`
}

//TextOutput is exported,it formats the data to plain text.
func (c Todo) TextOutput() string {
	// p := fmt.Sprintf(
	//   "Name: %s\nCompleted : %t\n_id: %s\ndate: %s\n",
	//   c.Name, c.completed, c.Oid, c.Date)

	p := fmt.Sprintf(
		"%s %t {%s %s}",
		c.Name, c.completed, c.Oid, c.Date)
	return p
}
