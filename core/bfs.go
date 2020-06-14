package core
import "fmt"
import queue "../queue"
import . "../definitions"

const MAX_INT int = (1 << 32) - 1 

type Visited map[Point]bool
type Distance map[Point]int
type Parent map[Point]Point

var dx = [4]int{0 , 1, 0,-1}
var dy = [4]int{-1, 0, 1, 0}


//Current problems/edge case
/*
1- When the food target is unreachable - that means the snake body probably is around the target - 
since bfs cant find a path to the target it will return {-1 -1} 
solution thoughts : when a food is unreachable follow tail instead ???

2 - the snakes targets a food that is reachable but lead to a dead end 
solution thougths : run a dfs on the sub graph that surround the food - if the number of vertices is less than the body of the snake, 
it is highly probable that the snake cant escape. then target a new food. This solution might be expensive cant run dfs on every move move request

3 - 
*/
func GetBestMoveToFood(start Point, b Board, you BattleSnake, g *GameState) string {
	
	
	visited, parent := Init(b)

	ChooseClosestFoodTarget(b, you.Head, &g.FoodTarget)

	fmt.Println("food chosen ", g.FoodTarget)
	
	q := queue.New()


	visited[start] = true
	queue.Push(q, start)
	// fmt.Println("starting bfs ", start)

	for ; !queue.Empty(q) ; {
		var currentVertex, _ = queue.Pop(q)

		// fmt.Println("currently visiting ", currentVertex)

		if currentVertex == g.FoodTarget {
			//fmt.Println("food target reached ! ", currentVertex, g.FoodTarget)
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
			if neighborVertex.X >= g.Width || neighborVertex.Y >= g.Height { continue }
			//skip snakes body
			if _ , in := g.EnemySnakes[neighborVertex]; in { continue }

			// fmt.Println("neigh", neighborVertex)
			if !visited[neighborVertex] {
				visited[neighborVertex] = true
				parent[neighborVertex] = currentVertex
				queue.Push(q, neighborVertex)

				// fmt.Println("pushing ", neighborVertex, "parent ", parent[neighborVertex])
			}

		}
		

	}
	// fmt.Println("parent map is", parent)

	//Find the path from food to start
	move := findPath(parent, start, g.FoodTarget)
	
	//if the food is unreachable
	noPath := Point{-1, -1}
	if noPath == move { 
		//tell the snake to target its tail
		GetBestMoveToFood(g.Tail, b, you, g)
	}
	
	moveString := pointToStringDirection(start, move)
	//TODO fix the logic 
	return moveString
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



func findPath(path Parent, start Point, goal Point) Point {
	// fmt.Println("trying to find path from ", start)
	move := goal

	for ; start != path[move] ; {

		move = path[move]


	} 
	
	
	// fmt.Println("best move is ", move)
	return move
	
	
	
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