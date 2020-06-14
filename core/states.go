package core
import . "../definitions"
import "fmt"
//some essential data structures representing the game state

type GameState struct {
	EnemySnakes map[Point]struct{} //set
	AllSnakes map[Point]struct{}
	FoodTarget Point
	Tail Point
	Width int
	Height int
}
var empty struct{}


//this function will be called at the /start POST request
func InitGameState(m MainRequest, g *GameState) {


	g.Width = m.Board.Width
	g.Height = m.Board.Height

	ChooseClosestFoodTarget(m.Board, m.You.Head, &g.FoodTarget)
	fmt.Println("food chosen init ", g.FoodTarget)

}
//This function will be called for every move request
func UpdateGameState(m MainRequest, g *GameState) {
	UpdateEnemySnakes(m.Board.Snakes, g)
	UpdateSnakeTail(m.You, g)
}

func UpdateEnemySnakes(snakes []BattleSnake, g *GameState) {
	g.EnemySnakes = make(map[Point]struct{})
	for _, snake :=  range snakes {
		for _, body := range snake.Body {
			g.EnemySnakes[body] = empty
		}
	}
}

func UpdateSnakeTail(you BattleSnake, g *GameState) {
	len := len(you.Body)
	g.Tail = you.Body[len - 1]
}