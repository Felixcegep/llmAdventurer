package gamelogic

import (
	"awesomeProject1/internal/llm"
	"fmt"
)

func Combat(p *llm.Player, b *llm.Boss, dialogue string) {

	fmt.Printf("combat %v vs %v \n", p.HeroName, b.Name)
	fmt.Println("Dialogue : ", dialogue)
	for {
		if b.Health <= 0 {
			break
		} else {
			var choice int
			fmt.Println(b.Health)
			fmt.Println("enter your choices 1. attack 2. use items")
			fmt.Scan(&choice)
			switch choice {
			case 1:
				fmt.Println("attack")
				b.Health -= 50

			case 2:
				fmt.Println("heal")
			case 3:
				fmt.Println("defense")
			default:
				fmt.Println("invalid choice")
			}
		}
	}
}
