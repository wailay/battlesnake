package core
import . "../definitions"
import "math"
//This class implements some logic to choose a food target



//Main function choose a food to target if none targeted in the moment 
//and keep targetting the same food until gone.
//TODO Strategies : target food with least enemy around
func ChooseClosestFoodTarget(b Board, head Point, target *Point) {

	foods := b.Food
	distance := 10000000.0
	for _, food := range foods {
		tempDistance := math.Abs(float64(food.X - head.X)) + math.Abs(float64(food.Y - head.Y))	
		
		if (tempDistance < distance ) {
			distance = tempDistance
			*target = food
		}
	}
	
	return 
	
	
}

//strat 1 : check if the food currently targeted have no dead ends
/*
□ □ □ □        □ ▤ ▤ ▤
▤ ▤ □ ▪   or   □ ▤ ▪ ▤ 
▪ ▤ □ □        □ □ ▤ ▤
               ▤ ▤ ▤ □
the food in the bottom left corner will surely lead to death
*/

func checkTargetDeadEnd(food Point, b Board) bool {

	return true
}

