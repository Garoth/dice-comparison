package main

import "fmt"

type Die struct {
	Label   string
	Numbers []int
}

func NewDie(label string, numbers []int) *Die {
	me := &Die{}
	me.Label = label
	me.Numbers = numbers
	return me
}

func CompareDice(die1, die2 *Die) {
	die1win := 0
	die2win := 0
	ties := 0

	for _, die1num := range die1.Numbers {
		for _, die2num := range die2.Numbers {
			if die1num > die2num {
				die1win++
			} else if die2num > die1num {
				die2win++
			} else {
				ties++
			}
		}
	}

	fmt.Println(die1.Label, "die vs", die2.Label, "die results:")
	fmt.Println(die1.Label, "wins", die1win)
	fmt.Println(die2.Label, "wins", die2win)
	sum := die1win + die2win + ties
	fmt.Println("So", die1.Label, "wins about", 100.0*die1win/sum, "percent",
		"of the time")
}

func main() {
	dice := []*Die{
		NewDie("Red", []int{3, 3, 3, 3, 3, 6}),
		NewDie("Blue", []int{2, 2, 2, 5, 5, 5}),
		NewDie("Olive", []int{1, 4, 4, 4, 4, 4}),
	}

	CompareDice(dice[0], dice[1])
	fmt.Println("")
	CompareDice(dice[1], dice[2])
	fmt.Println("")
	CompareDice(dice[2], dice[0])
}
