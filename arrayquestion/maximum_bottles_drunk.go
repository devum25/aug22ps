package arrayquestion

func MaxBottlesDrunk(numBottles int, numExchange int) int {
	drunk := numBottles
	bottles := 0
	emptyBottles := drunk

	for emptyBottles >= numExchange {
		bottles++
		emptyBottles = emptyBottles - numExchange
		numExchange++
		if emptyBottles < numExchange {
			drunk += bottles
			emptyBottles = emptyBottles + bottles
			bottles = 0
		}
	}

	return drunk
}
