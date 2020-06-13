package main

import (
	"fmt"
	"net/http"
	"log"
	"encoding/json"
	
)
import . "./utils"
import . "./core"

var mySnake SnakeConfig = SnakeConfig { 
	Apiversion : "1",
	Author : "wail",
	Color :"#E80978",
	Head : "dead",
	Tail : "bolt",
}

var currentGame MainRequest

func entry(w http.ResponseWriter, r *http.Request){
	snakeJSON, _ := json.Marshal(&mySnake)
	w.Header().Set("Content-Type", "application/json")
	//Write calls WriteHeader with http.StatusOk (200) before send data
	w.Write(snakeJSON)
	fmt.Printf("%s %s %s\n", r.Proto, r.Method, r.URL.Path)

}

func start(w http.ResponseWriter, r *http.Request){
	
	err := json.NewDecoder(r.Body).Decode(&currentGame)
	
	if err != nil { 
		fmt.Println(err) 
		return 
	}

	fmt.Printf("%s %s %s\n", r.Proto, r.Method, r.URL.Path)
	

}

func move(w http.ResponseWriter, r *http.Request) {

	err := json.NewDecoder(r.Body).Decode(&currentGame)
	
	if err != nil { 
		fmt.Println(err) 
		return 
	}


	nextMove := GetBestMoveToFood(currentGame.You, currentGame.Board)

	var move Move

	move.Move = nextMove
	moveJSON, _ := json.Marshal(&move)

	fmt.Println("sending move ", move)
	w.Header().Set("Content-Type", "application/json")
	w.Write(moveJSON)
	fmt.Printf("%s %s %s\n", r.Proto, r.Method, r.URL.Path)

}

func end(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%s %s %s\n", r.Proto, r.Method, r.URL.Path)
}


func main(){
	
	http.HandleFunc("/", entry)
	http.HandleFunc("/start", start)
	http.HandleFunc("/move", move)
	http.HandleFunc("/end", end)
	fmt.Println("started listening on port 3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
	// log.Print("sf")

	// board := [][]int{
	// 	{0,0,0,0,0},
	// 	{0,0,0,0,0},
	// 	{0,0,0,1,0},
	// 	{0,0,0,0,0},
	// }



	// dx := [4]int{0, 1, 0 ,-1}
	// dy := [4]int{-1, 0, 1 ,0}

	// for i := range dx {
	// 	fmt.Println(dx[i])
	// 	fmt.Println(dy[i])
	// }



}