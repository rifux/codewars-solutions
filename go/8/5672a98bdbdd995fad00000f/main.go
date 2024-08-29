/*
	Rock Paper Scissors
	Let's play! You have to return which player won! In case of a draw return Draw!.

	Examples(Input1, Input2 --> Output):

	"scissors", "paper" --> "Player 1 won!"
	"scissors", "rock" --> "Player 2 won!"
	"paper", "paper" --> "Draw!"

	(Source: https://www.codewars.com/kata/5672a98bdbdd995fad00000f/train/go )
*/

package main

import "fmt"

func Rps(p1, p2 string) string {
	m := make(map[string]int)
	r := make(map[int]string)
	m["scissors"], m["paper"], m["rock"] = 3, 2, 1
	r[-2], r[-1], r[0], r[1], r[2] = "Player 1 won!", "Player 2 won!", "Draw!", "Player 1 won!", "Player 2 won!"
	return r[m[p1]-m[p2]]
}

/*
	Scissors beats paper
	Paper beats scissors
	Rock beats scissors

	p1 has scissors, p2 has rock  -> p1 = 3, p2 = 1 -> r = 2
	p1 has scissors, p2 has paper -> p1 = 3, p2 = 2 -> r = 1
	p1 has paper, p2 has scissors -> p1 = 2, p2 = 3 -> r = -1
	p1 has paper, p2 has rock     -> p1 = 2, p2 = 1 -> r = 1
	p1 has rock, p2 has scissors  -> p1 = 1, p2 = 3 -> r = -2
	p1 has rock, p2 has paper     -> p1 = 1, p2 = 2 -> r = -1
*/

func main() {
	fmt.Println(Rps("rock", "scissors"))
	fmt.Println(Rps("scissors", "rock"))
	fmt.Println(Rps("rock", "rock"))
}
