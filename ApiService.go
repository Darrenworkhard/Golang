//Api pratice, giving an string array/list and print all of string
package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	//"encoding/json"
)

type Sudoku struct {
	name  string `json:"name"`
	input []int  `json:"input"`
}

//Get func
func GetFunc(w http.ResponseWriter, r *http.Request) {

	name, ok := r.URL.Query()["name"]
	input, ok := r.URL.Query()["input"]

	if !ok || len(name[0]) < 1 {
		log.Println("Missing name")
		return
	}

	if !ok || len(input[0]) < 1 {
		log.Println("Missing name")
		return
	}
	inputString := []int{}
	for index, v := range input {
		i, _ := strconv.Atoi(v)
		inputString[index] = i

	}

	inputStruc := Sudoku{name: name[0], input: inputString}
	PrintGet(inputStruc, w)

}

func PrintGet(inputStruc Sudoku, w http.ResponseWriter) {
	for _, v := range inputStruc.input {
		fmt.Fprintf(w, strconv.Itoa(v))
	}
}

//type Sudoku sudokuInput

func ShowEveryInput(w http.ResponseWriter, r *http.Request) {
	inputStruc := Sudoku{name: "Darren", input: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}}
	//}
	fmt.Println(inputStruc.name)
	for i, v := range inputStruc.input {
		fmt.Println(i, v)
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
	http.HandleFunc("/GetFunc", GetFunc)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	handleRequests()
}
