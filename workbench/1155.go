package workbench

// import "fmt"

type cacheKey struct {
	d, target int
}

// NumberOfDiceRollsToTarget
//
//	https://leetcode.com/problems/number-of-dice-rolls-with-target-sum/
func NumberOfDiceRollsToTarget(d int, f int, target int) int {
	cache := make(map[cacheKey]int)

	var diceRolls func(int, int) int
	diceRolls = func(dice, remaining int) int {
		if dice == 0 { // base case
			if remaining == 0 {
				return 1
			}
			return 0
		}

		// get computed value from cache
		if val, ok := cache[cacheKey{dice, remaining}]; ok {
			return val
		}

		count := 0

		// iterate through dice faces
		for i := 1; i <= f; i++ {
			if i <= remaining {
				count += diceRolls(dice-1, remaining-i)
			}
		}

		// store in hashmap
		cache[cacheKey{dice, remaining}] = count

		return count
	}

	return diceRolls(d, target)
}