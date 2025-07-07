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
				fmt.Println("use items")
				for i, items := range p.Inventory {
					fmt.Println(i, items)
				}
				var choiceItem int
				fmt.Print("enter you choice :")
				fmt.Scan(&choiceItem)
				if 0 <= choiceItem && choiceItem <= len(p.Inventory)-1 {
					fmt.Println("items used", p.Inventory[choiceItem])
					switch p.Inventory[choiceItem].ItemType {
					case 0:
						fmt.Println("heal")
					case 1:
						fmt.Println("attack")
					case 2:
						fmt.Println("defense")

					}
				} else {
					fmt.Println("choice dont exist")
				}

			default:
				fmt.Println("invalid choice")
			}
		}
	}
}
