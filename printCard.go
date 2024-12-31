package main

import "fmt"

func printCard(cards []string) {
	// fmt.Println("5 de Diamantes")
	for _, card := range cards {
		fmt.Println(card)
	}
}
