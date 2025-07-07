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
			b.Health -= 50
			fmt.Println("enemis ", b.Health)
		}
	}
}
