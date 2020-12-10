//Api pratice, giving an string array/list and print all of string
 package main

import (
	"fmt"
    "log"
	"net/http"
	//"encoding/json"
)

type Sudoku struct {
	name string `json:"name"`
	input []int `json:"input"`
}

//type Sudoku sudokuInput

func ShowEveryInput(w http.ResponseWriter, r *http.Request){
	inputStruc := Sudoku{name: "Darren", input: []int{1,2,3,4,5,6,7,8,9}}
		
	//}
	fmt.Println(inputStruc.name)
	for i,v := range inputStruc.input {
		fmt.Println(i,v)
	}
	//json.NewEncoder(w).Encode(inputStruc);
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	
    fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/Sudoku", ShowEveryInput)
    log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
    handleRequests()
}
