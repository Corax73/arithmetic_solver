package randomizer

import (
	"math/rand"
)

const SIMPLE_DIFFICULT_LEVEL = 2
const MEDIUM_DIFFICULT_LEVEL = 3

var simpleActionsMap = map[int][]string{1: {" + ", " - "}}
var fullActionsMap = map[int][]string{4: {" + ", " - ", " * ", " / "}}
var sumLimitMap = map[int][]int{1: {0, 10}, 2: {0, 20}, 3: {0, 100}, 4: {0, 1000}}
var randLimitMap = map[int][]int{1: {0, 6}, 2: {0, 11}, 3: {10, 51}, 4: {10, 501}}

func GetRandomValues(difficult int) (int, int, int, string, string) {
	var val1, val2, val3 int
	var action1, action2 string
	actions, ok := simpleActionsMap[difficult]
	if !ok {
		actions = simpleActionsMap[1]
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
		action1 = actions[rand.Intn(2)]
	} else if difficult == MEDIUM_DIFFICULT_LEVEL {
		val1, val2, val3 = getMediumData(randLimit, sumLimit)
		actions = simpleActionsMap[1]
		action1 = actions[rand.Intn(2)]
		action2 = actions[rand.Intn(2)]
	} else {
		val1, val2, val3 = getMediumData(randLimit, sumLimit)
		actions, ok := fullActionsMap[difficult]
		if !ok {
			actions = fullActionsMap[4]
		}
		action1 = actions[rand.Intn(4)]
		action2 = actions[rand.Intn(4)]
	}
	return val1, val2, val3, action1, action2
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

func getMediumData(randLimit, sumLimit []int) (int, int, int) {
reRand:
	val1 := rand.Intn(randLimit[1]) + randLimit[0]
	val2 := rand.Intn(randLimit[1]) + randLimit[0]
	val3 := rand.Intn(randLimit[1]) + randLimit[0]
	if val1 < val2 && val1+val2+val3 <= sumLimit[1] {
		valTemp := val1
		val1 = val2
		val2 = valTemp
	} else if val1+val2+val3 > sumLimit[1] {
		goto reRand
	}
	return val1, val2, val3
}
