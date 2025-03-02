package randomizer

import (
	"math/rand"
)

const SIMPLE_DIFFICULT_LEVEL = 3

var actionsMap = map[int][]string{1: {" + ", " - "}}
var sumLimitMap = map[int][]int{1: {0, 10}, 2: {0, 20}, 3: {0, 100}}
var randLimitMap = map[int][]int{1: {0, 6}, 2: {0, 11}, 3: {10, 51}}

func GetRandomValues(difficult int) (int, int, string) {
	var val1, val2 int
	actions, ok := actionsMap[difficult]
	if !ok {
		actions = actionsMap[1]
	}
	randLimit, ok := randLimitMap[difficult]
	if !ok {
		randLimit = randLimitMap[1]
	}
	sumLimit, ok := sumLimitMap[difficult]
	if !ok {
		sumLimit = sumLimitMap[1]
	}
	if difficult <= SIMPLE_DIFFICULT_LEVEL {
		val1, val2 = getSimpleData(randLimit, sumLimit)
	}
	return val1, val2, actions[rand.Intn(2)]
}

func getSimpleData(randLimit, sumLimit []int) (int, int) {
reRand:
	val1 := rand.Intn(randLimit[1]) + randLimit[0]
	val2 := rand.Intn(randLimit[1]) + randLimit[0]
	if val1 < val2 && val1+val2 <= sumLimit[1] {
		valTemp := val1
		val1 = val2
		val2 = valTemp
	} else if val1+val2 > sumLimit[1] {
		goto reRand
	}
	return val1, val2
}
