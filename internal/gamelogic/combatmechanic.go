package gamelogic

import (
	"fmt"
)

func calculateDamage(attack, defense int) int {
	damage := attack - (defense / 2)
	return damage
}
func Combat(p *Player, b *Boss, dialogue string) {
	fmt.Println("SUPERPOWER", p.SuperPower)
	fmt.Println("attack", b.Attack)
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
			if 0 < choice && choice <= 4 {
				playerTurn(p, b, choice)
				bossTurn(p, b)
			} else {
				fmt.Println("choice dont exist")
				continue
			}
		}

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

	case 4:
		p.Health += p.SuperPower.Health
		b.Health -= calculateDamage(p.SuperPower.Attack, b.Defense)
		p.Defense += p.SuperPower.Defense
		fmt.Println("CHEATCODE ACTIVATED")
	}
}

func bossTurn(p *Player, b *Boss) {
	fmt.Printf("%s attacks you!\n", b.Name)
	damage := calculateDamage(b.Attack, p.Defense)
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

	item := p.Inventory[choiceItem]

	p.Health += item.Health
	fmt.Println("You used a healing item.")
	b.Health -= calculateDamage(item.Attack, b.Defense)
	fmt.Println("You used an attack item.")
	p.Defense += item.Defense
	fmt.Println("You used a defense item.")

	// Remove item
	p.Inventory = append(p.Inventory[:choiceItem], p.Inventory[choiceItem+1:]...)
}
