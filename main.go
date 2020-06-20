package main

import (
	"fmt"
	"net/http"
	"log"
	"encoding/json"
	
)
import . "./core"
import . "./definitions"

var mySnake SnakeConfig = SnakeConfig { 
	Apiversion : "1",
	Author : "bigboi",
	Color :"#000000",
	Head : "dead",
	Tail : "bolt",
}

var currentGame MainRequest
var gameState GameState

func entry(w http.ResponseWriter, r *http.Request){
	snakeJSON, _ := json.Marshal(&mySnake)
	w.Header().Set("Content-Type", "application/json")
	//Write calls WriteHeader with http.StatusOk (200) before send data
	w.Write(snakeJSON)
	fmt.Printf("%s %s %s\n", r.Proto, r.Method, r.URL.Path)

}

func start(w http.ResponseWriter, r *http.Request){
	fmt.Printf("%s %s %s\n", r.Proto, r.Method, r.URL.Path)
	
	err := json.NewDecoder(r.Body).Decode(&currentGame)
	
	if err != nil { 
		fmt.Println(err) 
		return 
	}

	InitGameState(currentGame, &gameState)

	// fmt.Println("after init food chosen ", gameState.FoodTarget)
	

}

func move(w http.ResponseWriter, r *http.Request) {

	fmt.Printf("%s %s %s\n", r.Proto, r.Method, r.URL.Path)

	err := json.NewDecoder(r.Body).Decode(&currentGame)
	
	if err != nil { 
		fmt.Println(err) 
		return 
	}

	UpdateGameState(currentGame, &gameState)
	fmt.Println("my tail", gameState.Tail)
	nextMove := ""
	nextMove = GetBestMoveToFood(currentGame.You.Head, currentGame.Board, currentGame.You, &gameState, false)

	if nextMove == "nopath" {
		gameState.FoodTarget = gameState.Tail 
		nextMove = GetBestMoveToFood(currentGame.You.Head, currentGame.Board, currentGame.You, &gameState, true)
	}
	var move Move

	move.Move = nextMove
	moveJSON, _ := json.Marshal(&move)

	fmt.Println("sending move ", move)
	w.Header().Set("Content-Type", "application/json")
	w.Write(moveJSON)

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