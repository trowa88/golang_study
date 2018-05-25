package main

import "fmt"

func main() {
	i := 99
	for i >= 1 {
		bot := "bottles"
		if i == 1 {
			bot = "bottle"
		}
		fmt.Printf("%d %s of beer on the wail, %d %s of beer.\n", i, bot, i, bot)
		i--
		if i == 0 {
			fmt.Println("No more bottles of beer on the wall, No more bottles of beer.")
			i = 99
			fmt.Printf("Go to the store and buy some more, %d bottles of beer on the wall.", i)
			break
		} else {
			if i == 1 {
				bot = "bottle"
			}
			fmt.Printf("Take one down, pass it around, %d %s of beer on the wail.\n", i, bot)
		}
	}
}
