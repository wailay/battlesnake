package core
import "fmt"
import queue "../queue"
import . "../utils"

const MAX_INT int = (1 << 32) - 1 

type Visited map[Point]bool
type Distance map[Point]int
type Parent map[Point]Point

var foodTarget Point

type EnemySnakes map[Point]struct{} //this implements a set , keep the struct value empty
var empty struct {}

var dx = [4]int{0 , 1, 0,-1}
var dy = [4]int{-1, 0, 1, 0}


func GetBestMoveToFood(you BattleSnake, b Board) string {
	
	width, height := b.Width, b.Height
	visited, parent := Init(b)
	enemySnakes := updateEnemySnakes(b.Snakes)

	chooseTargetFood(b.Food, &foodTarget)
	
	fmt.Println("food chosen ", foodTarget)
	q := queue.New()

	start := you.Head
	visited[start] = true
	queue.Push(q, start)
	fmt.Println("starting bfs ", start)

	for ; !queue.Empty(q) ; {
		var currentVertex, _ = queue.Pop(q)

		// fmt.Println("currently visiting ", currentVertex)

		if currentVertex == foodTarget {
			fmt.Println("food target reached ! ", currentVertex, foodTarget)
			break
		}
		//visit neighbors
		for i := range dx {
			neighborVertex := Point {
				X : currentVertex.X + dx[i],
				Y : currentVertex.Y + dy[i],
			}

			
			//skip vertex that are out of bounds
			if neighborVertex.X < 0 || neighborVertex.Y < 0 { continue }
			if neighborVertex.X >= width || neighborVertex.Y >= height { continue }
			//skip snakes body
			if _ , in := enemySnakes[neighborVertex]; in { continue }

			// fmt.Println("neigh", neighborVertex)
			if !visited[neighborVertex] {
				visited[neighborVertex] = true
				parent[neighborVertex] = currentVertex
				queue.Push(q, neighborVertex)

				// fmt.Println("pushing ", neighborVertex, "parent ", parent[neighborVertex])
			}

		}
		

	}
	fmt.Println("parent map is", parent)

	//Find the path from food to start
	move := findPath(parent, start, foodTarget)
	
	
	
	//TODO fix the logic 
	return move
}

//Initiliase some data structures needed for bfs algorithm
func Init(b Board) (Visited, Parent) {

	visited := make(Visited)
	distance := make(Distance)
	parent := make(Parent)

	width, height := b.Width, b.Height
	for x := 0 ; x < width; x++ {
		for y:=0; y < height; y++ {
			p := Point{X : x, Y : y}
			visited[p] = false
			distance[p] = MAX_INT
			parent[p] = Point{-1, -1}
		} 
	}

	return visited, parent
	
}

func updateEnemySnakes(snakes []BattleSnake) EnemySnakes {
	enemySnakes := make(EnemySnakes)
	for _, snake :=  range snakes {
		for _, body := range snake.Body {
			enemySnakes[body] = empty
		}
	}

	return enemySnakes
}

func findPath(path Parent, start Point, goal Point) string {
	fmt.Println("trying to find path from ", start)
	move := goal

	for ; start != path[move] ; {

		move = path[move]


	} 
	
	
	fmt.Println("best move is ", move)
	return pointToStringDirection(start, move)
	
	
	
}

func pointToStringDirection(parent Point, move Point) string {
	x := move.X - parent.X
	y := move.Y - parent.Y
	fmt.Println("diff", x, y)
	if x > 0 { return "right" }
	if x < 0 { return "left"}
	if y < 0 { return "down"}
	if y > 0 { return "up"}

	return "up" //should not reach here / TODO - refactor ? bad code ?
}

//this function choose a food to target if none targeted in the moment 
//and keep targetting the same food until gone.
//TODO Strategies : target food with least enemy around
func chooseTargetFood(foods []Point, target *Point) {
	if target == nil {
		*target = foods[0]
		return 
	}

	for _, food := range foods {

		if food == *target { break }
		*target = food
	}
	return 
	
	
}