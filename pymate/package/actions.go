package pymate

type movement interface {
	randomDirection(currX int, currY int, currHeading int, distance int) int
	randomDistance(currX int, currY int, currHeading int, newHeading int) int
	fullRandom(currX int, currY int, currHeading int) (int, int)
}

type changeEnergeticState interface {
	eat(currCalories int, caloriesEaten int) int // increase energy by amount eaten
	expend(currCalories int, caloriesExpended int) int // decrease energy by amount expended
}