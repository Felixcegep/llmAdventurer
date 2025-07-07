package gamelogic

import (
	"fmt"
)

func calculateDamage(attack, defense int) int {
	damage := attack - (defense / 2)
	return damage
}
func Combat(p *Player, b *Boss, dialogue string) {

	fmt.Printf("combat %v vs %v \n", p.HeroName, b.Name)
	fmt.Println("Dialogue : ", dialogue)
	for {
		fmt.Printf("%s %d  vs  %s %d \n", p.HeroName, p.Health, b.Name, b.Health)

		if b.Health <= 0 || p.Health <= 0 {

			break
		} else {
			var choice int
			fmt.Println("1. default attack\n 2. healing\n 3. select item")
			fmt.Print("Enter choice: ")
			fmt.Scan(&choice)
			if 0 <= choice && choice <= 3 {

				playerTurn(p, b, choice)
			} else {
				fmt.Println("choice dont exist")
			}
		}
		// boss combat here
	}
}
func playerTurn(p *Player, b *Boss, choices int) {
	switch choices {
	case 1:
		fmt.Println("attack")
		b.Health -= calculateDamage(p.Attack, b.Defense)

	case 2:
		p.healingpotion()
	case 3:
		PlayerUseItem(p, b)
	}

}

func bossTurn(p *Player, b *Boss) {
	fmt.Printf("%s attacks you!\n", b.Name)
	damage := calculateDamage(b.Health, p.Defense)
	p.Health -= damage
	fmt.Printf("You took %d damage! Your health is now %d\n", damage, p.Health)
}
func PlayerUseItem(p *Player, b *Boss) {
	for i, item := range p.Inventory {
		fmt.Println(i, item)
	}

	var choiceItem int
	fmt.Print("Choose item: ")
	fmt.Scan(&choiceItem)

	if choiceItem >= 0 && choiceItem < len(p.Inventory) {
		item := p.Inventory[choiceItem]
		switch item.ItemType {
		case 0: // Heal
			p.Health += item.Health
			fmt.Println("You used a healing item.")
		case 1: // Attack
			b.Health -= calculateDamage(item.Attack, b.Defense)
			fmt.Println("You used an attack item.")
		case 2: // Defense
			p.Defense += item.Defense
			fmt.Println("You used a defense item.")
		}
		// Remove item
		p.Inventory = append(p.Inventory[:choiceItem], p.Inventory[choiceItem+1:]...)
	} else {
		fmt.Println("Invalid item choice.")
	}
}
