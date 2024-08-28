/*
	Create a combat function that takes the player's current health and the amount of damage received, and returns the player's new health. Health can't be less than 0.

	(Source: https://www.codewars.com/kata/586c1cf4b98de0399300001d/train/go )
*/

package main

func combat(health, damage float64) float64 {
	if damage > health {
		return 0
	}
	return health - damage
}
