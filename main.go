package main

import (
	"fmt"
	Internal "getYoutube/Internal"
)

func main() {
	var url string
	print("Insert URL video Youtube: ")
	fmt.Scan(&url)

	dataApi := Internal.GetApi(Internal.GetID(url))

	for {
		var choice int
		Internal.PrintChoice()

		fmt.Scan(&choice)

		switch {
		case choice == 1:
			println(dataApi.GetTitle())
		case choice == 2:
			println(dataApi.GetChannel())
		case choice == 3:
			println(dataApi.GetDuration())
		case choice == 4:
			println(dataApi.GetViews())
		case choice == 5:
			println(dataApi.GetLikes())
		case choice == 6:
			println(dataApi.GetFavorites())
		case choice == 7:
			println(dataApi.GetComments())
		case choice == 8:
			print("\nInsert URL video Youtube: ")
			fmt.Scan(&url)
			dataApi = Internal.GetApi(Internal.GetID(url))
		case choice >= 9:
			println("\nOps choice not found...")
		}
		if choice == 0 {
			println("\nThanks!!")
			println("GitHub: @cafesao")
			break
		}
	}
}
